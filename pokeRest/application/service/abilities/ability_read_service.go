package abilities

import (
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type AbilityReadService struct {
	repository repository.IAbilityRepository
}

// ポケモンの持つ特性取得
func (s AbilityReadService) GetAbilitiesOfPokemon(pokemonId int) interface{} {
	return nil
}

// 特性一覧取得
func (s AbilityReadService) GetAllAbility(cmd command.GetAllEntityCommand) ([]interface{}, error) {
	return nil, nil
}
