package pokemon

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type PokemonReadService struct {
	pokemonRepository repository.IPokemonRepository
}

func (s PokemonReadService) GetPokemon(id int) model.Pokemon {
	s.pokemonRepository.FindById(id)
}
