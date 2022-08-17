package service

import (
	"github.com/Symthy/PokeRest/internal/domain/damages"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
)

type DefenseSide struct {
	actualValues      *value.PokemonActualValues
	nature            value.PokemonNature
	pokemonTypeSet    *value.PokemonTypeSet
	hasItem           bool
	additionalEffects *battles.BattleEffects
}

func NewDefenseSide(
	actualValues *value.PokemonActualValues, pokemonTypes *value.PokemonTypeSet,
	nature value.PokemonNature, additionalEffects *battles.BattleEffects, hasItem bool) DefenseSide {
	return DefenseSide{
		actualValues:      actualValues,
		pokemonTypeSet:    pokemonTypes,
		nature:            nature,
		hasItem:           hasItem,
		additionalEffects: additionalEffects,
	}
}

func (b DefenseSide) PokemonTypes() *value.PokemonTypeSet {
	return b.pokemonTypeSet
}

func (b DefenseSide) toDefenseSidePokemon() *damages.DefenseSidePokemon {
	return damages.NewDefenseSidePokemon(b.actualValues, b.nature, b.pokemonTypeSet, b.hasItem)
}

func (p DefenseSide) toDefenseSideBattleEffects() *damages.DefenseSideBattleEffects {
	return damages.NewDefenseSideBattleEffects(p.additionalEffects)
}
