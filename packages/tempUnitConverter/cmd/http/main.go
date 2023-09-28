package main

import (
	"net/http"
	"os"

	//
	// adapter "tempUnitConverter/framework/httpadapter"

	"error_handler/errorhandler"
	"github.com/labstack/echo/v4"
)

const DefaultPort = "8001"

func main() {
	e := echo.New()
	e.HTTPErrorHandler = errorhandler.EchoErrorHandler
	e.GET("/ctof", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, ctof")
	})
	e.GET("/ftoc", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, ctof")
	})
	e.Start(":" + getPort())
}

func getPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = DefaultPort
	}
	return p
}
