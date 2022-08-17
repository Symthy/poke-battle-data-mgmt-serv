package damages

import (
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
)

type DefenseSidePokemon struct {
	defensePokemonActualValues *value.PokemonActualValues
	defensePokemonNature       value.PokemonNature
	defensePokemonTypes        *value.PokemonTypeSet
	defensePokemonHasItems     bool
}

func NewDefenseSidePokemon(
	deffencePokemonActualValues *value.PokemonActualValues, defensePokemonNature value.PokemonNature,
	defensePokemonTypes *value.PokemonTypeSet, hasItem bool) *DefenseSidePokemon {
	return &DefenseSidePokemon{
		defensePokemonActualValues: deffencePokemonActualValues,
		defensePokemonNature:       defensePokemonNature,
		defensePokemonTypes:        defensePokemonTypes,
		defensePokemonHasItems:     hasItem,
	}
}

type DefenseSideBattleEffects struct {
	side              battles.BattleSideType
	statusCorrections *battles.StatusCorrections
	damageCorrections *battles.DamageCorrections
	overrides         *battles.BattleOverrideValues
}

func NewDefenseSideBattleEffects(effects *battles.BattleEffects) *DefenseSideBattleEffects {
	return &DefenseSideBattleEffects{
		side:              battles.BattleDefenseSide,
		statusCorrections: battles.NewStatusCorrections(effects.Corrections(), battles.BattleDefenseSide),
		damageCorrections: battles.NewDamageCorrections(effects.Corrections(), battles.BattleDefenseSide),
		overrides:         effects.Overrides(),
	}
}
