package errorhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func EchoErrorHandler(err error, c echo.Context) {
	herr, ok := err.(HttpError)
	if !ok {
		herr = HttpError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	c.JSON(herr.Code, herr.Message)
}
