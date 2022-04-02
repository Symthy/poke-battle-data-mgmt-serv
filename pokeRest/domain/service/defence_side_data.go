package service

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/damages"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type DefenseSideData struct {
	actualValues      value.PokemonActualValues
	nature            value.PokemonNature
	pokemonTypes      value.PokemonTypeSet
	additionalEffects value.BattleEffects
}

func NewDefenseSideData(
	actualValues value.PokemonActualValues, pokemonTypes value.PokemonTypeSet,
	nature value.PokemonNature, additionalEffects value.BattleEffects) DefenseSideData {
	return DefenseSideData{
		actualValues:      actualValues,
		pokemonTypes:      pokemonTypes,
		nature:            nature,
		additionalEffects: additionalEffects,
	}
}

func (b DefenseSideData) PokemonTypes() value.PokemonTypeSet {
	return b.pokemonTypes
}

func (b DefenseSideData) toDefenseSidePokemon() damages.DefenseSidePokemon {
	return damages.NewDefenseSidePokemon(b.actualValues, b.nature, b.pokemonTypes)
}

func (p DefenseSideData) toDefenseSideBattleEffects() damages.DefenseSideBattleEffects {
	return damages.NewDefenseSideBattleEffects(p.additionalEffects)
}
