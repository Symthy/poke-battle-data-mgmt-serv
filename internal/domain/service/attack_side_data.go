package service

import (
	"github.com/Symthy/PokeRest/internal/domain/damages"
	"github.com/Symthy/PokeRest/internal/domain/entity/moves"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
)

var _ moves.IMoveNotificationForDamageCalc = (*AttackSide)(nil)

type AttackSide struct {
	actualValues      *value.PokemonActualValues
	pokemonTypeSet    *value.PokemonTypeSet
	nature            value.PokemonNature
	moveSpecies       value.MoveSpecies
	moveType          value.PokemonType
	movePower         uint16 // 威力
	isMoveContained   bool
	canMoveGuard      bool
	hasItem           bool
	additionalEffects *battles.BattleEffects
}

func NewAttackSide(
	actualValues *value.PokemonActualValues, pokemonType *value.PokemonTypeSet, nature value.PokemonNature,
	additionalEffects *battles.BattleEffects, move *moves.Move, hasItem bool,
) AttackSide {
	attackPokemon := &AttackSide{
		actualValues:      actualValues,
		pokemonTypeSet:    pokemonType,
		nature:            nature,
		hasItem:           hasItem,
		additionalEffects: additionalEffects,
	}
	move.NotifyForDamageCalc(attackPokemon)
	return *attackPokemon
}

func (b *AttackSide) SetMoveSpecies(moveSpecies value.MoveSpecies) {
	b.moveSpecies = moveSpecies
}
func (a *AttackSide) SetMoveType(moveType value.PokemonType) {
	a.moveType = moveType
}
func (b *AttackSide) SetMovePower(movePower uint16) {
	b.movePower = movePower
}
func (b *AttackSide) SetIsMoveContained(isMoveContained bool) {
	b.isMoveContained = isMoveContained
}
func (b *AttackSide) SetCanMoveGuard(canMoveGuard bool) {
	b.canMoveGuard = canMoveGuard
}

func (p AttackSide) toAttackSidePokemon() *damages.AttackSidePokemon {
	return damages.NewAttackSidePokemon(p.actualValues, p.pokemonTypeSet, p.nature, p.hasItem)
}

func (p AttackSide) toAttackMove() *damages.AttackMove {
	return damages.NewAttackMove(
		p.moveSpecies,
		p.moveType,
		p.movePower,
		p.isMoveContained,
		p.canMoveGuard,
	)
}

func (p AttackSide) toAttackSideBattleEffects() *damages.AttackSideBattleEffects {
	return damages.NewAttackSideBattleEffects(p.additionalEffects)
}
