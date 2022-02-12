package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service/abilities"
	d_abilities "github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type AbilityController struct {
	service          abilities.AbilityReadService
	responseResolver response.ResponseResolver[d_abilities.Abilities, server.Abilities]
}

func NewAbilityController(service abilities.AbilityReadService) *AbilityController {
	return &AbilityController{
		service:          service,
		responseResolver: response.NewResponseResolver(response.ConvertAbilitiesToResponse),
	}
}

func (c AbilityController) GetAbilities(ctx echo.Context, next int, pageSize int) error {
	cmd := command.NewPaginationCommand(next, pageSize)
	domain, err := c.service.FindAll(cmd)
	return c.responseResolver.Resolve(ctx, domain, err)
}
