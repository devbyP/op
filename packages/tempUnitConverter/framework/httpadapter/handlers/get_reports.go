package handlers

import (
	"net/http"
	"strconv"

	"tempUnitConverter/framework/httpadapter/api"

	eh "error_handler/errorhandler"

	"github.com/labstack/echo/v4"
)

func (a *Adapter) HandleGetReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		return eh.NewHttpError(err.Error(), http.StatusBadRequest)
	}
	r, err := a.app.GetReportByID(id)
	if err != nil {
		return eh.NewHttpError(err.Error(), http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, api.ReponseGetReport{Data: r})
}

func (a *Adapter) HandleGetAllReports(c echo.Context) error {
	r, err := a.app.GetAllReports()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, api.ResponseGetAllReports{Data: r})
}
