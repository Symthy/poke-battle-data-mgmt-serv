package service

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/damages"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

var _ moves.IMoveNotificationForDamageCalc = (*AttackSide)(nil)

type AttackSide struct {
	actualValues      value.PokemonActualValues
	pokemonType       value.PokemonTypeSet
	nature            value.PokemonNature
	moveSpecies       value.MoveSpecies
	moveType          value.PokemonType
	movePower         uint16 // 威力
	isMoveContained   bool
	canMoveGuard      bool
	additionalEffects value.BattleEffects
}

func NewAttackSide(
	actualValues value.PokemonActualValues, pokemonType value.PokemonTypeSet,
	nature value.PokemonNature, additionalEffects value.BattleEffects, move moves.Move,
) AttackSide {
	attackPokemon := &AttackSide{
		actualValues:      actualValues,
		pokemonType:       pokemonType,
		nature:            nature,
		additionalEffects: additionalEffects,
	}
	move.NotifyForDamageCalc(attackPokemon)
	return *attackPokemon
}

func (a AttackSide) SetMoveType(moveType value.PokemonType) {
	a.moveType = moveType
}
func (b AttackSide) SetMoveSpecies(moveSpecies value.MoveSpecies) {
	b.moveSpecies = moveSpecies
}
func (b AttackSide) SetMovePower(movePower uint16) {
	b.movePower = movePower
}
func (b AttackSide) SetIsMoveContained(isMoveContained bool) {
	b.isMoveContained = isMoveContained
}
func (b AttackSide) SetCanMoveGuard(canMoveGuard bool) {
	b.canMoveGuard = canMoveGuard
}

func (p AttackSide) MoveType() value.PokemonType {
	return p.moveType
}

func (p AttackSide) toAttackSidePokemon() damages.AttackSidePokemon {
	return damages.NewAttackSidePokemon(p.actualValues, p.pokemonType, p.nature)
}

func (p AttackSide) toAttackMove() damages.AttackMove {
	return damages.NewAttackMove(
		p.moveSpecies,
		p.moveType,
		p.movePower,
		p.isMoveContained,
		p.canMoveGuard,
	)
}

func (p AttackSide) toAttackSideBattleEffects() damages.AttackSideBattleEffects {
	return damages.NewAttackSideBattleEffects(p.additionalEffects)
}
