package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"tempUnitConverter/domain"
	"tempUnitConverter/framework/httpadapter/api"

	eh "error_handler/errorhandler"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) HandleGetConvertTemp(c echo.Context) error {
	var res domain.TemperatureUnit
	var err error
	req := api.RequestConverTemp{}
	qp := c.QueryParams()
	if req.Temp, err = strconv.ParseFloat(qp.Get("temp"), 64); err != nil {
		return eh.NewHttpError(err.Error(), http.StatusBadRequest)
	}
	req.Unit = qp.Get("unit")
	qp.Get("")
	if req.Unit == "" {
		req.Unit = "C"
	}
	switch strings.ToUpper(req.Unit) {
	case "C":
		res = a.app.CToF(domain.C(req.Temp))
	case "F":
		res = a.app.FToC(domain.F(req.Temp))
	default:
		return eh.NewHttpError("unknown temp unit", http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, api.ReponseCovertTemp{Unit: res.Unit(), Temp: res.Float64()})
}
