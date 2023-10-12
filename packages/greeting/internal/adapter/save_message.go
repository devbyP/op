package adapter

import (
	"net/http"

	"op/greetingapi"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) HandleSaveMessage(c echo.Context) error {
	req := &greetingapi.SaveNewGreetingRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	res, err := a.greetingService.SaveMessage(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
