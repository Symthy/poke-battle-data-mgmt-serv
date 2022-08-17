package controller

import (
	"github.com/Symthy/PokeRest/internal/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/internal/application/command"
	"github.com/Symthy/PokeRest/internal/application/service/abilities"
	d_abilities "github.com/Symthy/PokeRest/internal/domain/entity/abilities"
	"github.com/Symthy/PokeRest/internal/presentation/response"
	"github.com/labstack/echo/v4"
)

type AbilityController struct {
	service            abilities.AbilityReadService
	singleDataResolver response.ResponseResolver[d_abilities.Ability, server.Ability]
	multiDataResolver  response.ResponseResolver[d_abilities.Abilities, server.Abilities]
}

func NewAbilityController(service abilities.AbilityReadService) *AbilityController {
	return &AbilityController{
		service:            service,
		singleDataResolver: response.NewResponseResolver(response.ConvertAbilityToResponse),
		multiDataResolver:  response.NewResponseResolver(response.ConvertAbilitiesToResponse),
	}
}

func (c AbilityController) GetAbilityById(ctx echo.Context, id uint64) error {
	domain, err := c.service.FindAbility(id)
	return c.singleDataResolver.Resolve(ctx, domain, err)
}

func (c AbilityController) GetAbilities(ctx echo.Context, next uint64, pageSize uint64) error {
	cmd := command.NewPaginationCommand(next, pageSize)
	domains, err := c.service.FindAll(cmd)
	return c.multiDataResolver.Resolve(ctx, domains, err)
}

func (c AbilityController) GetAbilityByPokemonId(ctx echo.Context, pokemonId uint64) error {
	domains, err := c.service.FindOfPokemon(pokemonId)
	return c.multiDataResolver.Resolve(ctx, domains, err)
}
