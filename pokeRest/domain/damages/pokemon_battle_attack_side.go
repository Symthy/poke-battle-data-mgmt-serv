package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

type AttackSidePokemon struct {
	attackPokemonActualValues value.PokemonActualValues
	attackPokemonType         value.PokemonTypeSet
	attackPokemonNature       value.PokemonNature
}

func NewAttackSidePokemon(
	attackPokemon value.PokemonActualValues, attackPokemonType value.PokemonTypeSet,
	attackPokemonNature value.PokemonNature) AttackSidePokemon {
	return AttackSidePokemon{
		attackPokemonActualValues: attackPokemon,
		attackPokemonType:         attackPokemonType,
		attackPokemonNature:       attackPokemonNature,
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
	isContained bool, canGuard bool) AttackMove {
	return AttackMove{
		species:     species,
		pokemonType: pokemonType,
		power:       power,
		isContained: isContained,
		canGuard:    canGuard,
	}
}

type AttackSideBattleEffects struct {
	side battles.BattleSideType
	*battles.StatusCorrections
	*battles.PowerCorrections
	*battles.MovePowerCorrections
	*battles.DamageCorrections
	*battles.BattleOverrideValues
}

func NewAttackSideBattleEffects(effects battles.BattleEffects) AttackSideBattleEffects {
	return AttackSideBattleEffects{
		side:                 battles.BattleAttackSide,
		StatusCorrections:    battles.NewStatusCorrections(effects.Corrections(), battles.BattleAttackSide),
		PowerCorrections:     battles.NewPowerCorrections(effects.Corrections()),
		MovePowerCorrections: battles.NewMovePowerCorrections(effects.Corrections()),
		DamageCorrections:    battles.NewDamageCorrections(effects.Corrections(), battles.BattleAttackSide),
		BattleOverrideValues: effects.Overrides(),
	}
}
