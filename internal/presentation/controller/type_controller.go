package controller

import (
	"net/http"

	"github.com/Symthy/PokeRest/internal/application/service/types"
	"github.com/Symthy/PokeRest/internal/presentation/lang"
	"github.com/Symthy/PokeRest/internal/presentation/response"
	"github.com/labstack/echo/v4"
)

type TypeController struct {
	service types.TypeReadService
}

func NewTypeController(service types.TypeReadService) *TypeController {
	return &TypeController{service: service}
}

func (c TypeController) GetTypes(ctx echo.Context) error {
	lang := lang.NewRequestLanguage(*ctx.Request())
	var types []string = c.service.GetTypes().GenerateTypeNames(lang.Lang())
	return ctx.JSON(http.StatusOK, types)
}

func (c TypeController) GetTypeCompatibility(ctx echo.Context) error {
	lang := lang.NewRequestLanguage(*ctx.Request())
	typesTable := c.service.GetTypeCompatibility()
	res := response.ConvertTypeTableToResponse(typesTable, lang.Lang())
	return ctx.JSON(http.StatusOK, res)
}

func (c TypeController) GetAttackTypeCompatibility(ctx echo.Context, attackType string) error {
	lang := lang.NewRequestLanguage(*ctx.Request())
	types := c.service.GetAttackTypeCompatibility(attackType)
	res := response.ConvertTypesToResponse(types, lang.Lang())
	return ctx.JSON(http.StatusOK, res)
}

func (c TypeController) GetDeffenseTypeCompatibility(ctx echo.Context, deffenceType string) error {
	lang := lang.NewRequestLanguage(*ctx.Request())
	types := c.service.GetDeffenceTypeCompatibility(deffenceType)
	res := response.ConvertTypesToResponse(types, lang.Lang())
	return ctx.JSON(http.StatusOK, res)
}
