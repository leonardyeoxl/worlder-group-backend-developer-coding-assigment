package main

import (
  "time"
  "log"
  "strconv"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "golang.org/x/net/context"
	"google.golang.org/grpc"
  pb "github.com/leonardyeoxl/worlder-group-backend-developer-coding-assigment/src/data"
)

type Data struct {
  SensorValue    int64  `json:"SensorValue"`
  ID1            int64  `json:"ID1"`
  ID2            string `json:"ID2"`
  Timestamp      string  `json:"Timestamp"`
}

type Server struct {
}

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/data", getData)

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getData(c echo.Context) error {
  ID1 := c.QueryParam("ID1")
  ID2 := c.QueryParam("ID2")
  start_timestamp := c.QueryParam("start_timestamp")
  end_timestamp := c.QueryParam("end_timestamp")
  _ID1, _ := strconv.ParseInt(ID1, 10, 64)
  _start_timestamp, _ := strconv.ParseInt(start_timestamp, 10, 64)
  _end_timestamp, _ := strconv.ParseInt(end_timestamp, 10, 64)

  conn, err := grpc.Dial("database-client:50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewRetrieveDataServiceClient(conn)
	in := &pb.Request{
    ID1: _ID1,
    ID2: ID2,
    StartTimestamp: _start_timestamp,
    EndTimestamp: _end_timestamp,
  }
	response, err := client.RetrieveData(context.Background(), in)
	if err != nil {
		log.Fatalf("retrieve data error %v", err)
	}

  retrieve_data := response.GetData()

  var collection []*Data

  if len(retrieve_data) > 0 {
    for _, rd := range retrieve_data {
      datetime := time.Unix(rd.GetTimestamp(), 0)
      d := &Data{
        ID1: rd.GetID1(),
        ID2: rd.GetID2(),
        SensorValue: rd.GetSensorValue(),
        Timestamp: datetime.Format(time.RFC1123Z),
      }
      collection = append(collection, d)
    }
  }

  return c.JSON(http.StatusOK, collection)
}