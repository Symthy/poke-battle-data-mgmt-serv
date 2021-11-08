package controller

import "github.com/Symthy/PokeRest/pokeRest/application/service/pokemon"

type PokemonController struct {
	service pokemon.PokemonReadService
}

// Todo: use wire
func NewPokemonController(service pokemon.PokemonReadService) *PokemonController {
	return &PokemonController{service: service}
}

//func GetPokemon()
