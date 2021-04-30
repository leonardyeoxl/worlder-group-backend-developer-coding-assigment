package main

import (
  "fmt"
  "log"
  "strconv"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

// User
type Data struct {
  ID1  int64 `json:"ID1"`
  ID2 int64 `json:"ID2"`
  StartTimestamp int64 `json:"StartTimestamp"`
  EndTimestamp int64 `json:"EndTimestamp"`
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
  _ID2, _ := strconv.ParseInt(ID2, 10, 64)
  _start_timestamp, _ := strconv.ParseInt(start_timestamp, 10, 64)
  _end_timestamp, _ := strconv.ParseInt(end_timestamp, 10, 64)

  d := &Data{
    ID1: _ID1,
    ID2: _ID2,
    StartTimestamp: _start_timestamp,
    EndTimestamp: _end_timestamp,
  }
  return c.JSON(http.StatusOK, d)
}