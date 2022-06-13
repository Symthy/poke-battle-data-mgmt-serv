package service

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/damages"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type DefenseSide struct {
	actualValues      value.PokemonActualValues
	nature            value.PokemonNature
	pokemonTypes      value.PokemonTypeSet
	additionalEffects value.BattleEffects
}

func NewDefenseSide(
	actualValues value.PokemonActualValues, pokemonTypes value.PokemonTypeSet,
	nature value.PokemonNature, additionalEffects value.BattleEffects) DefenseSide {
	return DefenseSide{
		actualValues:      actualValues,
		pokemonTypes:      pokemonTypes,
		nature:            nature,
		additionalEffects: additionalEffects,
	}
}

func (b DefenseSide) PokemonTypes() value.PokemonTypeSet {
	return b.pokemonTypes
}

func (b DefenseSide) toDefenseSidePokemon() damages.DefenseSidePokemon {
	return damages.NewDefenseSidePokemon(b.actualValues, b.nature, b.pokemonTypes)
}

func (p DefenseSide) toDefenseSideBattleEffects() damages.DefenseSideBattleEffects {
	return damages.NewDefenseSideBattleEffects(p.additionalEffects)
}
