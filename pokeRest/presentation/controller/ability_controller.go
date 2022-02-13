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

func (c AbilityController) GetAbilityById(ctx echo.Context, id int) error {
	domain, err := c.service.FindAbility(uint(id))
	return c.singleDataResolver.Resolve(ctx, domain, err)
}

func (c AbilityController) GetAbilities(ctx echo.Context, next int, pageSize int) error {
	cmd := command.NewPaginationCommand(next, pageSize)
	domains, err := c.service.FindAll(cmd)
	return c.multiDataResolver.Resolve(ctx, domains, err)
}

func (c AbilityController) GetAbilityByPokemonId(ctx echo.Context, pokemonId int) error {
	domains, err := c.service.FindOfPokemon(uint(pokemonId))
	return c.multiDataResolver.Resolve(ctx, domains, err)
}
