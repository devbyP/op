package httpadapter

import (
	"tempUnitConverter/app"

	"github.com/labstack/echo/v4"
)

type Adapter struct {
	app app.AppPort
}

func (a *Adapter) handleGetReport(c echo.Context) error {
	return nil
}

func (a *Adapter) handleGetAllReports(c echo.Context) error {
	return nil
}

func (a *Adapter) handleConvertCtoF(c echo.Context) error {
	return nil
}

func (a *Adapter) handleConvertFtoC(c echo.Context) error {
	return nil
}
