package pokemons

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type PokemonReadService struct {
	pokemonRepository repository.IPokemonRepository
}

func NewPokemonReadService(repository repository.IPokemonRepository) PokemonReadService {
	return PokemonReadService{pokemonRepository: repository}
}

// UC: 単体詳細取得
func (s PokemonReadService) GetPokemon(id uint) (pokemons.Pokemon, error) {
	pokemon, err := s.pokemonRepository.FindById(id)
	return *pokemon, err
}

// UC: 一覧取得 StatsOfPokemonFinder

// UC: 検索
func (s PokemonReadService) SearchPokemons() {
	return
}
