package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service/types"
	"github.com/Symthy/PokeRest/pokeRest/presentation"
)

type TypeController struct {
	service types.TypeReadService
}

func NewTypeController(service types.TypeReadService) *TypeController {
	return &TypeController{service: service}
}

func (c TypeController) GetTypeCompatibility(lang string) server.TypeCompatibility {
	return presentation.ConvertTypesToResponse(c.service.GetTypeCompatibility(), lang)
}

func (c TypeController) GetTypes(lang string) []string {
	return c.service.GetTypes().GenerateTypeNames(lang)
}
