package identifier

type PokemonId struct {
	ValueId
}

func NewPokemonId(value uint) PokemonId {
	return PokemonId{ValueId{value}}
}

func NewEmptyPokemonId() PokemonId {
	return PokemonId{ValueId{0}}
}
