package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

type Configuration struct {
	AMQPConnectionURL string
}

type StreamTask struct {
	SensorValue int
	ID1 int
	ID2 string
	Timestamp int64
}

var Config = Configuration{
	AMQPConnectionURL: "amqp://guest:guest@mq:5672/",
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

func main() {
	conn, err := amqp.Dial(Config.AMQPConnectionURL)
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare(
		"stream-from-listener", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	handleError(err, "Could not declare `stream` queue")

	ID2Map := [2]string{"A", "B"}
	ID1Map := [3]int{1,2,3}
	max := 100
	min := 1
	rand.Seed(time.Now().UnixNano())

	for {
		time.Sleep(time.Second * 1)
		
		randomIndexID2 := rand.Intn(len(ID2Map))
		ID2Value := ID2Map[randomIndexID2]
		randomIndexID1 := rand.Intn(len(ID1Map))
		ID1Value := ID1Map[randomIndexID1]

		streamTask := StreamTask{
			SensorValue: rand.Intn(max - min) + min,
			ID1: ID1Value,
			ID2: ID2Value,
			Timestamp: time.Now().Unix(),
		}

		body, err := json.Marshal(streamTask)
		if err != nil {
			handleError(err, "Error encoding JSON")
		}

		err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		})

		if err != nil {
			log.Fatalf("Error publishing message: %s", err)
		}

		log.Printf("StreamTask: %d %d %s %d", streamTask.SensorValue, streamTask.ID1, streamTask.ID2, streamTask.Timestamp)
	}

}