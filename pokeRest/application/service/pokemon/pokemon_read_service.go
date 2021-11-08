package pokemon

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type PokemonReadService struct {
	pokemonRepository repository.IPokemonRepository
}

func NewPokemonReadService(repository repository.IPokemonRepository) PokemonReadService {
	return PokemonReadService{pokemonRepository: repository}
}

func (s PokemonReadService) GetPokemon(id int) model.Pokemon {
	return s.pokemonRepository.FindById(id)
}
