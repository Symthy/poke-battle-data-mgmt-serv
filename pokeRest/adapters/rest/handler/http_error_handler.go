package handler

import (
	"github.com/Symthy/PokeRest/pokeRest/errs"
	"github.com/labstack/echo/v4"
)

func PokeRestHttpErrorHandler(err error, c echo.Context) {
	appErr := errs.BuildAppError(err)
	appErr.WriteServerError(c.Echo().Logger)
	// delegate to echo default handler
	c.Echo().DefaultHTTPErrorHandler(appErr.ApiErrorResponse(), c)
}
