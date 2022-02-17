package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", sleep)
	e.GET("/non", non_sleep)

	e.Logger.Fatal(e.Start(":8080"))
}

type JsonMessage struct {
	Message string `json:"message"`
}

func sleep(c echo.Context) error {
	time.Sleep(time.Second * 5)
	data := JsonMessage{
		Message: "sleep",
	}
	return c.JSON(http.StatusOK, data)
}

func non_sleep(c echo.Context) error {
	data := JsonMessage{
		Message: "non_sleep",
	}
	return c.JSON(http.StatusOK, data)
}
