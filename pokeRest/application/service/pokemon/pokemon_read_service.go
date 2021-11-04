package pokemon

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/database"
)

type PokemonReadService struct {
	pokemonRepository repository.IPokemonRepository
}

func DefaultPokemonReadService() PokemonReadService {
	return PokemonReadService{pokemonRepository: database.NewPokemonRepository()}
}

func NewPokemonReadService(repository repository.IPokemonRepository) PokemonReadService {
	return PokemonReadService{pokemonRepository: repository}
}

func (s PokemonReadService) GetPokemon(id int) model.Pokemon {
	return s.pokemonRepository.FindById(id)
}
