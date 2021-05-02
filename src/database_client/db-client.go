package main

import (
	"net"
	"encoding/json"
	"log"
	"github.com/streadway/amqp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/leonardyeoxl/worlder-group-backend-developer-coding-assigment/src/data"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

type StreamTask struct {
	SensorValue int64
	ID1 int64
	ID2 string
	Timestamp int64
}

type server struct{}

func (s server) RetrieveData(ctx context.Context, in *pb.Request) (*pb.Collection, error) {
	db, err := sql.Open("mysql", "username:password@tcp(mysql-database:3306)/")
	name := "iotdb"
	_,err = db.Exec("USE "+name)
	if err != nil {
		panic(err)
	}
	
	var selDB *sql.Rows
	if (in.ID1 == 0 || in.ID2 == "") {
		selDB, err = db.Query("SELECT * FROM Stream WHERE Timestamp >= ? and Timestamp <= ?", in.StartTimestamp, in.EndTimestamp)
	} else if (in.StartTimestamp == 0 && in.EndTimestamp == 0) {
		selDB, err = db.Query("SELECT * FROM Stream WHERE ID1=? and ID2=?", in.ID1, in.ID2)
	} else {
		selDB, err = db.Query("SELECT * FROM Stream WHERE (ID1=? and ID2=?) and (Timestamp >= ? and Timestamp <= ?)", in.ID1, in.ID2, in.StartTimestamp, in.EndTimestamp)
	}
	
	
    if err != nil {
        panic(err.Error())
    }
	var collection []*pb.Data
    for selDB.Next() {
        var SensorValue, ID1, Timestamp int64
        var ID2 string
        err = selDB.Scan(&SensorValue, &ID1, &ID2, &Timestamp)
        if err != nil {
            panic(err.Error())
        }
		new_data := &pb.Data{
            SensorValue: SensorValue,
            ID1: ID1,
            ID2: ID2,
			Timestamp: Timestamp,
        }

		collection = append(collection, new_data)
    }
	return &pb.Collection{Data: collection}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func createDatabase() (db *sql.DB) {
	db, err := sql.Open("mysql", "username:password@tcp(mysql-database:3306)/")
	name := "iotdb"
	_,err = db.Exec("USE "+name)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE IF NOT EXISTS Stream ( SensorValue integer, ID1 integer, ID2 varchar(32), Timestamp integer)")
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	stream := StreamTask{}

	go func() {
		// create listiner
		lis, err := net.Listen("tcp", ":50005")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		
		// create grpc server
		s := grpc.NewServer()
		pb.RegisterRetrieveDataServiceServer(s, server{})
		
		log.Println("start server")
		// and start...
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	db := createDatabase()

	conn, err := amqp.Dial("amqp://guest:guest@mq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"stream-from-listener", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			json.Unmarshal([]byte(d.Body), &stream)
			insForm, err := db.Prepare("INSERT INTO Stream(SensorValue,ID1,ID2,Timestamp) VALUES(?,?,?,?)")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(stream.SensorValue, stream.ID1, stream.ID2, stream.Timestamp)
		}
	}()

	<-forever
}