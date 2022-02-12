package controller

import (
	"net/http"

	"github.com/Symthy/PokeRest/pokeRest/application/service/types"
	"github.com/Symthy/PokeRest/pokeRest/presentation/lang"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type TypeController struct {
	service types.TypeReadService
}

func NewTypeController(service types.TypeReadService) *TypeController {
	return &TypeController{service: service}
}

func (c TypeController) GetTypeCompatibility(ctx echo.Context) error {
	lang := lang.NewRequestLanguage(*ctx.Request())
	types := c.service.GetTypeCompatibility()
	res := response.ConvertTypesToResponse(types, lang.Lang())
	ctx.JSON(http.StatusOK, res)
	return nil
}

func (c TypeController) GetTypes(ctx echo.Context) error {
	lang := lang.NewRequestLanguage(*ctx.Request())
	var types []string = c.service.GetTypes().GenerateTypeNames(lang.Lang())
	ctx.JSON(http.StatusOK, types)
	return nil
}
