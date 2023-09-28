package handlers

import (
	"net/http"
	"strings"

	"tempUnitConverter/domain"
	"tempUnitConverter/framework/httpadapter/api"

	eh "error_handler/errorhandler"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) HandlePostSaveReport(c echo.Context) error {
	var newId int
	var err error
	req := api.RequestSaveReport{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	// default C to F
	if req.Unit == "" {
		req.Unit = "C"
	}
	switch strings.ToUpper(req.Unit) {
	case "C":
		newId, err = a.app.SaveC(domain.C(req.Temp))
	case "F":
		newId, err = a.app.SaveFtoC(domain.F(req.Temp))
	default:
		return eh.NewHttpError("unknown temp unit", http.StatusBadRequest)
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, api.ResponseSaveReport{ID: newId})
}
