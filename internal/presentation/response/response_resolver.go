package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseResolver[TD any, TR any] struct {
	converter func(model TD) TR
}

func NewResponseResolver[TD any, TR any](converter func(model TD) TR) ResponseResolver[TD, TR] {
	return ResponseResolver[TD, TR]{
		converter: converter,
	}
}

func (r ResponseResolver[TD, TR]) Resolve(ctx echo.Context, domain *TD, err error) error {
	if err != nil {
		return err
	}
	res := r.converter(*domain)
	err = ctx.JSON(http.StatusOK, res)
	if err != nil {
		return err
	}
	return nil
}
