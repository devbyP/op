package main

import (
	"os"

	"tempUnitConverter/app"
	"tempUnitConverter/framework/httpadapter/handlers"
	"tempUnitConverter/framework/sampledb"

	"error_handler/errorhandler"

	"github.com/labstack/echo/v4"
)

const DefaultPort = "8001"

func main() {
	db := sampledb.New("./temp.csv")
	tempApp := app.NewTemperature(db)
	h := handlers.NewAdapter(tempApp)

	e := echo.New()
	e.HTTPErrorHandler = errorhandler.EchoErrorHandler

	e.GET("/convert-temp", h.HandleGetConvertTemp)
	e.GET("/report/:id", h.HandleGetConvertTemp)
	e.GET("/report", h.HandleGetConvertTemp)
	e.POST("/report", h.HandlePostSaveReport)

	e.Start(":" + getPort())
}

func getPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = DefaultPort
	}
	return p
}
