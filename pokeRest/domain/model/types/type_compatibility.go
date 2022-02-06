package types

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type TypeCompatibility struct {
	compatibilityTable [][]float32
	types              []value.PokemonType
}
