package types

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type PokemonTypes struct {
	items []value.PokemonType
}

func NewPokemonTypes(types ...value.PokemonType) PokemonTypes {
	items := types
	if len(types) == 0 {
		items = value.GetPokemonTypes()
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
		if lang == "ja-JP" {
			typeOrder[i] = t.NameJP()
			continue
		}
		typeOrder[i] = t.NameEN()
	}
	return typeOrder
}
