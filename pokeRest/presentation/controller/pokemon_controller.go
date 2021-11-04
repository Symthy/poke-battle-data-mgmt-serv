package controller

import "github.com/Symthy/PokeRest/pokeRest/application/service/pokemon"

type PokemonController struct {
	service pokemon.PokemonReadService
}

func DefaultPokemonController() PokemonController {
	return PokemonController{service: pokemon.DefaultPokemonReadService()}
}

// Todo: use wire
func NewPokemonController(service pokemon.PokemonReadService) *PokemonController {
	return &PokemonController{service: service}
}

func GetPokemon()
