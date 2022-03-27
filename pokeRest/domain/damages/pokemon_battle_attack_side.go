package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type AttackSidePokemon struct {
	attackPokemon       value.PokemonActualValues
	attackPokemonType   value.PokemonTypeSet
	attackPokemonNature value.PokemonNature
}

func NewAttackSidePokemon(
	attackPokemon value.PokemonActualValues, attackPokemonType value.PokemonTypeSet,
	attackPokemonNature value.PokemonNature) AttackSidePokemon {
	return AttackSidePokemon{
		attackPokemon:       attackPokemon,
		attackPokemonType:   attackPokemonType,
		attackPokemonNature: attackPokemonNature,
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

func (m AttackMove) getSpecies() value.MoveSpecies {
	return m.species
}

func (m AttackMove) getPokemonType() value.PokemonType {
	return m.pokemonType
}
func (m AttackMove) getPower() int {
	return m.power
}

type AttackSideBattleEffects struct {
	side value.BattleSideType
	value.StatusCorrections
	value.PowerCorrections
	value.MovePowerCorrections
	value.DamageCorrections
	value.BattleOverrideValues
}

func NewAttackSideBattleEffects(effects value.BattleEffects) AttackSideBattleEffects {
	return AttackSideBattleEffects{
		side:                 value.BattleAttackSide,
		StatusCorrections:    value.NewStatusCorrections(effects.Corrections()),
		PowerCorrections:     value.NewPowerCorrections(effects.Corrections()),
		MovePowerCorrections: value.NewMovePowerCorrections(effects.Corrections()),
		DamageCorrections:    value.NewDamageCorrections(effects.Corrections()),
		BattleOverrideValues: effects.Overrides(),
	}
}
