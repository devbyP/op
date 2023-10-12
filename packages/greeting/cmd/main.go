package main

import (
	"os"

	"op/internal/adapter"
	"op/internal/handler"

	"github.com/labstack/echo/v4"
)

const DefaultPort = "8000"

var (
	service *handler.Service
	adt     *adapter.Adapter
)

func init() {
	service = handler.New("./message.csv")
	adt = adapter.New(service)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}
	e := echo.New()
	e.GET("/:id", adt.HandleGetMessage)
	e.GET("/", adt.HandleGetAllMessage)
	e.POST("/", adt.HandleSaveMessage)
	e.Logger.Fatal(e.Start(":" + port))
}
