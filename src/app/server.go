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

	e.Logger.Fatal(e.Start(":8080"))
}

func sleep(c echo.Context) error {
	time.Sleep(time.Second * 10)
	return c.String(http.StatusOK, "Hello, World!")
}
