package httpadapter

import (
	"net/http"
	"strconv"

	"tempUnitConverter/app"
	"tempUnitConverter/domain"
	"tempUnitConverter/framework/httpadapter/api"

	eh "error_handler/errorhandler"

	"github.com/labstack/echo/v4"
)

type Adapter struct {
	app app.AppPort
}

func (a *Adapter) handleGetReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		return eh.NewHttpError(err.Error(), http.StatusBadRequest)
	}
	r, err := a.app.GetReport(id)
	if err != nil {
		return eh.NewHttpError(err.Error(), http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, api.ReponseGetReport{Data: r})
}

func (a *Adapter) handleGetAllReports(c echo.Context) error {
	r, err := a.app.GetReports()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, api.ResponseGetAllReports{Data: r})
}

func (a *Adapter) handleGetConvertTemp(c echo.Context) error {
	var res domain.TemperatureUnit
	var err error
	req := api.RequestConverTemp{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	switch req.Unit {
	case "C":
		res, err = a.app.CtoF(domain.C(req.Temp))
	case "F":
		res, err = a.app.CtoF(domain.C(req.Temp))
	default:
		return eh.NewHttpError("unknown temp unit", http.StatusBadRequest)
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, api.ReponseCovertTemp{Unit: res.Unit(), Temp: res.Float64()})
}

func (a *Adapter) handlePostSaveReport(c echo.Context) error {
	return nil
}
