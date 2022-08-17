package damages

import (
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
)

type AttackSidePokemon struct {
	attackPokemonActualValues *value.PokemonActualValues
	attackPokemonType         *value.PokemonTypeSet
	attackPokemonNature       value.PokemonNature
	attackPokemonHasItem      bool
}

func NewAttackSidePokemon(
	actualValues *value.PokemonActualValues, pokemonTypeSet *value.PokemonTypeSet,
	nature value.PokemonNature, hasItem bool) *AttackSidePokemon {
	return &AttackSidePokemon{
		attackPokemonActualValues: actualValues,
		attackPokemonType:         pokemonTypeSet,
		attackPokemonNature:       nature,
		attackPokemonHasItem:      hasItem,
	}
}

type AttackMove struct {
	species     value.MoveSpecies
	pokemonType value.PokemonType
	power       uint16 // 威力
	isContained bool
	canGuard    bool
}

func NewAttackMove(species value.MoveSpecies, pokemonType value.PokemonType, power uint16,
	isContained bool, canGuard bool) *AttackMove {
	return &AttackMove{
		species:     species,
		pokemonType: pokemonType,
		power:       power,
		isContained: isContained,
		canGuard:    canGuard,
	}
}

type AttackSideBattleEffects struct {
	side              battles.BattleSideType
	statusCorrections *battles.StatusCorrections
	powerCorrections  *battles.PowerCorrections
	moveCorrections   *battles.MovePowerCorrections
	damageCorrections *battles.DamageCorrections
	overrides         *battles.BattleOverrideValues
}

func NewAttackSideBattleEffects(effects *battles.BattleEffects) *AttackSideBattleEffects {
	return &AttackSideBattleEffects{
		side:              battles.BattleAttackSide,
		statusCorrections: battles.NewStatusCorrections(effects.Corrections(), battles.BattleAttackSide),
		powerCorrections:  battles.NewPowerCorrections(effects.Corrections()),
		moveCorrections:   battles.NewMovePowerCorrections(effects.Corrections()),
		damageCorrections: battles.NewDamageCorrections(effects.Corrections(), battles.BattleAttackSide),
		overrides:         effects.Overrides(),
	}
}
