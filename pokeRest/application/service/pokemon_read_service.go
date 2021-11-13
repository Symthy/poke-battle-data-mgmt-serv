package service

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

func (s PokemonReadService) GetPokemon(id uint) (model.Pokemon, error) {
	pokemon, err := s.pokemonRepository.FindById(id)
	return pokemon, err
}
