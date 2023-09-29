package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

const DefaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		t := time.Now().Unix()
		return c.String(http.StatusOK, fmt.Sprintf("Hello, World!, %d", t))
	})
	e.Logger.Fatal(e.Start(":" + port))
}
