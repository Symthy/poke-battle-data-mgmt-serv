package types

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/types"
)

type TypeReadService struct {
}

func NewTypeReadService() TypeReadService {
	return TypeReadService{}
}

// UC: タイプ相性表取得
func (s TypeReadService) GetTypeCompatibility() types.TypeCompatibility {
	t := types.NewTypeCompatibility()
	return t
}

// UC: タイプ一覧取得
func (s TypeReadService) GetTypes() types.PokemonTypes {
	return types.NewPokemonTypes()
}
