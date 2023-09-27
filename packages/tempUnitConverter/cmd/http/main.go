package main

import (
	"net/http"
	"os"

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
	e.GET("/test-bind", func(c echo.Context) error {
		f := Foo("")
		if err := c.Bind(&f); err != nil {
			return err
		}
		return c.String(http.StatusOK, string(f))
	})
	e.Start(":" + getPort())
}

type Foo string

func getPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = DefaultPort
	}
	return p
}
