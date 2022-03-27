package service

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/damages"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type DefenceSideData struct {
	actualValues      value.PokemonActualValues
	nature            value.PokemonNature
	pokemonTypes      value.PokemonTypeSet
	additionalEffects value.BattleEffects
}

func NewDefenceSideData(
	actualValues value.PokemonActualValues, pokemonTypes value.PokemonTypeSet,
	nature value.PokemonNature, additionalEffects value.BattleEffects) DefenceSideData {
	return DefenceSideData{
		actualValues:      actualValues,
		pokemonTypes:      pokemonTypes,
		nature:            nature,
		additionalEffects: additionalEffects,
	}
}

func (b DefenceSideData) PokemonTypes() value.PokemonTypeSet {
	return b.pokemonTypes
}

func (b DefenceSideData) toDefenceSidePokemon() damages.DefenceSidePokemon {
	return damages.NewDefenceSidePokemon(b.actualValues, b.nature, b.pokemonTypes)
}

func (p DefenceSideData) toDefenceSideBattleEffects() damages.DefenceSideBattleEffects {
	return damages.NewDefenceSideBattleEffects(p.additionalEffects)
}
