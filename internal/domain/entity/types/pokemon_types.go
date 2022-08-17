package types

import "github.com/Symthy/PokeRest/internal/domain/value"

type PokemonTypes struct {
	items []value.PokemonType
}

func NewPokemonTypes(types ...value.PokemonType) PokemonTypes {
	items := types
	if len(types) == 0 {
		items = value.PokemonTypeAll()
	}
	return PokemonTypes{
		items: items,
	}
}

func (t PokemonTypes) Items() []value.PokemonType {
	return t.items
}

func (t PokemonTypes) Size() int {
	return len(t.items)
}

func (t PokemonTypes) GenerateTypeNames(lang string) []string {
	typeOrder := make([]string, t.Size())
	for i, t := range t.Items() {
		typeOrder[i] = t.ResolveTypeName(lang)
	}
	return typeOrder
}
