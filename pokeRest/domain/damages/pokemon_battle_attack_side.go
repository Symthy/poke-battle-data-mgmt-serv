package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

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
	power       int // 威力
	isContained bool
	canGuard    bool
}

func NewAttackMove(species value.MoveSpecies, pokemonType value.PokemonType, power int,
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
	side value.BattleSideType
	*value.StatusCorrections
	*value.PowerCorrections
	*value.MovePowerCorrections
	*value.DamageCorrections
	*value.BattleOverrideValues
}

func NewAttackSideBattleEffects(effects value.BattleEffects) AttackSideBattleEffects {
	return AttackSideBattleEffects{
		side:                 value.BattleAttackSide,
		StatusCorrections:    value.NewStatusCorrections(effects.Corrections(), value.BattleAttackSide),
		PowerCorrections:     value.NewPowerCorrections(effects.Corrections()),
		MovePowerCorrections: value.NewMovePowerCorrections(effects.Corrections()),
		DamageCorrections:    value.NewDamageCorrections(effects.Corrections(), value.BattleAttackSide),
		BattleOverrideValues: effects.Overrides(),
	}
}
