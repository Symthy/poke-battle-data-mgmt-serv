package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service/pokemons"
	d_pokemons "github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type PokemonController struct {
	service          pokemons.PokemonSummaryReadService
	responseResolver response.ResponseResolver[d_pokemons.Pokemon, server.Pokemon]
}

func NewPokemonController(service pokemons.PokemonSummaryReadService) *PokemonController {
	return &PokemonController{
		service:          service,
		responseResolver: response.NewResponseResolver(response.ConvertPokemonToResponse),
	}
}

func (c PokemonController) GetPokemon(ctx echo.Context, id int) error {
	domain, error := c.service.FindPokemon(uint(id))
	return c.responseResolver.Resolve(ctx, domain, error)
}
