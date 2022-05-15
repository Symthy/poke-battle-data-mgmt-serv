package service

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/damages"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

var _ moves.IMoveNotificationForDamageCalc = (*AttackSideData)(nil)

type AttackSideData struct {
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

func NewAttackSideData(
	actualValues value.PokemonActualValues, pokemonType value.PokemonTypeSet,
	nature value.PokemonNature, additionalEffects value.BattleEffects, move moves.Move,
) AttackSideData {
	attackPokemon := &AttackSideData{
		actualValues:      actualValues,
		pokemonType:       pokemonType,
		nature:            nature,
		additionalEffects: additionalEffects,
	}
	move.NotifyForDamageCalc(attackPokemon)
	return *attackPokemon
}

func (a AttackSideData) SetMoveType(moveType value.PokemonType) {
	a.moveType = moveType
}
func (b AttackSideData) SetMoveSpecies(moveSpecies value.MoveSpecies) {
	b.moveSpecies = moveSpecies
}
func (b AttackSideData) SetMovePower(movePower uint16) {
	b.movePower = movePower
}
func (b AttackSideData) SetIsMoveContained(isMoveContained bool) {
	b.isMoveContained = isMoveContained
}
func (b AttackSideData) SetCanMoveGuard(canMoveGuard bool) {
	b.canMoveGuard = canMoveGuard
}

func (p AttackSideData) MoveType() value.PokemonType {
	return p.moveType
}

func (p AttackSideData) toAttackSidePokemon() damages.AttackSidePokemon {
	return damages.NewAttackSidePokemon(p.actualValues, p.pokemonType, p.nature)
}

func (p AttackSideData) toAttackMove() damages.AttackMove {
	return damages.NewAttackMove(
		p.moveSpecies,
		p.moveType,
		p.movePower,
		p.isMoveContained,
		p.canMoveGuard,
	)
}

func (p AttackSideData) toAttackSideBattleEffects() damages.AttackSideBattleEffects {
	return damages.NewAttackSideBattleEffects(p.additionalEffects)
}
