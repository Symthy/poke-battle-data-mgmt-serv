package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type DefenseSidePokemon struct {
	defensePokemonActualValues value.PokemonActualValues
	defensePokemonNature       value.PokemonNature
	defensePokemonTypes        value.PokemonTypeSet
}

func NewDefenseSidePokemon(
	deffencePokemonActualValues value.PokemonActualValues, defensePokemonNature value.PokemonNature,
	defensePokemonTypes value.PokemonTypeSet) DefenseSidePokemon {
	return DefenseSidePokemon{
		defensePokemonActualValues: deffencePokemonActualValues,
		defensePokemonNature:       defensePokemonNature,
		defensePokemonTypes:        defensePokemonTypes,
	}
}

type DefenseSideBattleEffects struct {
	side value.BattleSideType
	*value.StatusCorrections
	*value.DamageCorrections
	*value.BattleOverrideValues
}

func NewDefenseSideBattleEffects(effects value.BattleEffects) DefenseSideBattleEffects {
	return DefenseSideBattleEffects{
		side:                 value.BattleDefenseSide,
		StatusCorrections:    value.NewStatusCorrections(effects.Corrections(), value.BattleDefenseSide),
		DamageCorrections:    value.NewDamageCorrections(effects.Corrections(), value.BattleDefenseSide),
		BattleOverrideValues: effects.Overrides(),
	}
}
