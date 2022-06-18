package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

type DefenseSidePokemon struct {
	defensePokemonActualValues *value.PokemonActualValues
	defensePokemonNature       value.PokemonNature
	defensePokemonTypes        value.PokemonTypeSet
	defensePokemonHasItems     bool
}

func NewDefenseSidePokemon(
	deffencePokemonActualValues *value.PokemonActualValues, defensePokemonNature value.PokemonNature,
	defensePokemonTypes value.PokemonTypeSet, hasItem bool) *DefenseSidePokemon {
	return &DefenseSidePokemon{
		defensePokemonActualValues: deffencePokemonActualValues,
		defensePokemonNature:       defensePokemonNature,
		defensePokemonTypes:        defensePokemonTypes,
		defensePokemonHasItems:     hasItem,
	}
}

type DefenseSideBattleEffects struct {
	side battles.BattleSideType
	*battles.StatusCorrections
	*battles.DamageCorrections
	*battles.BattleOverrideValues
}

func NewDefenseSideBattleEffects(effects *battles.BattleEffects) *DefenseSideBattleEffects {
	return &DefenseSideBattleEffects{
		side:                 battles.BattleDefenseSide,
		StatusCorrections:    battles.NewStatusCorrections(effects.Corrections(), battles.BattleDefenseSide),
		DamageCorrections:    battles.NewDamageCorrections(effects.Corrections(), battles.BattleDefenseSide),
		BattleOverrideValues: effects.Overrides(),
	}
}
