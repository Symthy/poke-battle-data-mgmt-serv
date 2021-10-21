package controller

import "github.com/Symthy/PokeRest/pokeRest/application/service/pokemon"

type PokemonController struct {
	service pokemon.PokemonReadService
}
