package types

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model/types"
)

type TypeReadService struct {
}

func NewTypeReadService() TypeReadService {
	return TypeReadService{}
}

// UC: タイプ相性表取得
func (s TypeReadService) GetTypeCompatibility() types.TypeCompatibility {
	t := types.NewTypeCompatibility()
	print(t.GenerateTypeCompatibilityTable())
	return t
}

// UC: タイプ一覧取得
func (s TypeReadService) GetTypes() types.PokemonTypes {
	return types.NewPokemonTypes()
}