package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

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
	side battles.BattleSideType
	*battles.StatusCorrections
	*battles.DamageCorrections
	*battles.BattleOverrideValues
}

func NewDefenseSideBattleEffects(effects battles.BattleEffects) DefenseSideBattleEffects {
	return DefenseSideBattleEffects{
		side:                 battles.BattleDefenseSide,
		StatusCorrections:    battles.NewStatusCorrections(effects.Corrections(), battles.BattleDefenseSide),
		DamageCorrections:    battles.NewDamageCorrections(effects.Corrections(), battles.BattleDefenseSide),
		BattleOverrideValues: effects.Overrides(),
	}
}
