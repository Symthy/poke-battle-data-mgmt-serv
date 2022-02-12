package types

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type TypeCompatibility struct {
	targetType       value.PokemonType
	typeToDamageRate map[value.PokemonType]float32
	typeOrder        PokemonTypes
}

func NewTypeCompatibility(
	targetType value.PokemonType,
	types map[value.PokemonType]float32,
	typeOrder PokemonTypes) TypeCompatibility {
	return TypeCompatibility{
		targetType:       targetType,
		typeToDamageRate: types,
		typeOrder:        typeOrder,
	}
}

func (t TypeCompatibility) ResolveTargetType(lang string) string {
	return t.targetType.ResolveTypeName(lang)
}

func (t TypeCompatibility) GenerateTypeDamageRates() []float32 {
	typeRecord := make([]float32, t.typeOrder.Size())
	for i, pType := range t.typeOrder.Items() {
		typeRecord[i] = t.typeToDamageRate[pType]
	}
	return typeRecord
}

func (t TypeCompatibility) GenerateTypeNames(lang string) []string {
	return t.typeOrder.GenerateTypeNames(lang)
}
