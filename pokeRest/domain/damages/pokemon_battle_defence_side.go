package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type DefenceSidePokemon struct {
	defencePokemonActualValues value.PokemonActualValues
	defencePokemonNature       value.PokemonNature
	defencePokemonTypes        value.PokemonTypeSet
}

func NewDefenceSidePokemon(
	deffencePokemonActualValues value.PokemonActualValues, defencePokemonNature value.PokemonNature,
	defencePokemonTypes value.PokemonTypeSet) DefenceSidePokemon {
	return DefenceSidePokemon{
		defencePokemonActualValues: deffencePokemonActualValues,
		defencePokemonNature:       defencePokemonNature,
		defencePokemonTypes:        defencePokemonTypes,
	}
}

type DefenceSideBattleEffects struct {
	side value.BattleSideType
	value.StatusCorrections
	value.DamageCorrections
	value.BattleOverrideValues
}

func NewDefenceSideBattleEffects(effects value.BattleEffects) DefenceSideBattleEffects {
	return DefenceSideBattleEffects{
		side:                 value.BattleDefenceSide,
		StatusCorrections:    value.NewStatusCorrections(effects.Corrections(), value.BattleDefenceSide),
		DamageCorrections:    value.NewDamageCorrections(effects.Corrections(), value.BattleDefenceSide),
		BattleOverrideValues: effects.Overrides(),
	}
}
