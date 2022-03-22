package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type BattleDefencePokemon struct {
	value.PokemonActualValues
	nature            value.PokemonNature
	pokemonTypes      value.PokemonTypeSet
	additionalEffects value.BattleEffects
}

func NewBattleDefencePokemon(
	actualValues value.PokemonActualValues, pokemonTypes value.PokemonTypeSet,
	nature value.PokemonNature, additionalEffects value.BattleEffects) BattleDefencePokemon {
	return BattleDefencePokemon{
		PokemonActualValues: actualValues,
		pokemonTypes:        pokemonTypes,
		nature:              nature,
		additionalEffects:   additionalEffects,
	}
}

func (b BattleDefencePokemon) PokemonTypes() value.PokemonTypeSet {
	return b.pokemonTypes
}
