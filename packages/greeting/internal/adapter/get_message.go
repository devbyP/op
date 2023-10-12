package adapter

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) HandleGetMessage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	res, err := a.greetingService.Greeting(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (a *Adapter) HandleGetAllMessage(c echo.Context) error {
	res, err := a.greetingService.GreetingAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
