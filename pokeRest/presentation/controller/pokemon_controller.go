package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/presentation"
)

type PokemonController struct {
	service service.PokemonReadService
}

func NewPokemonController(service service.PokemonReadService) *PokemonController {
	return &PokemonController{service: service}
}

func (c PokemonController) GetPokemon(id float32) (server.Pokemon, error) {
	ret := server.Pokemon{}
	pokemon, err := c.service.GetPokemon(uint(id))
	if err == nil {
		ret = presentation.ConvertPokemonToResponse(pokemon)
	}
	return ret, err
}
