package types

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type TypeReadService struct {
}

func NewTypeReadService() TypeReadService {
	return TypeReadService{}
}

// UC: タイプ相性表取得
func (s TypeReadService) GetTypeCompatibility() types.TypeCompatibilityTable {
	t := types.NewTypeCompatibilityTable()
	return t
}

// UC: タイプ一覧取得
func (s TypeReadService) GetTypes() types.PokemonTypes {
	return types.NewPokemonTypes()
}

// UC: 攻撃側タイプ相性取得
func (s TypeReadService) GetAttackTypeCompatibility(attackType string) types.TypeCompatibility {
	pType := value.NewPokemonType(attackType)
	return types.NewTypeCompatibilityTable().ResolveAttackTypeCompatibility(pType)
}

// UC: 防御側タイプ相性取得
func (s TypeReadService) GetDeffenceTypeCompatibility(deffenceType string) types.TypeCompatibility {
	pType := value.NewPokemonType(deffenceType)
	return types.NewTypeCompatibilityTable().ResolveDeffenceTypeCompatibility(pType)
}
