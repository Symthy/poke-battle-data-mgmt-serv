package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

var _ moves.IMoveNotificationForDamageCalc = (*BattleAttackPokemon)(nil)

type BattleAttackPokemon struct {
	value.PokemonActualValues
	pokemonType       value.PokemonTypeSet
	nature            value.PokemonNature
	moveSpecies       value.MoveSpecies
	moveType          value.PokemonType
	movePower         int // 威力
	isMoveContained   bool
	canMoveGuard      bool
	additionalEffects value.BattleEffects
}

func NewBattleAttackPokemon(
	actualValues value.PokemonActualValues, pokemonType value.PokemonTypeSet,
	nature value.PokemonNature, additionalEffects value.BattleEffects, move moves.Move,
) BattleAttackPokemon {
	attackPokemon := &BattleAttackPokemon{
		PokemonActualValues: actualValues,
		pokemonType:         pokemonType,
		nature:              nature,
		additionalEffects:   additionalEffects,
	}
	move.NotifyForDamageCalc(attackPokemon)
	return *attackPokemon
}

func (b BattleAttackPokemon) SetMoveType(moveType value.PokemonType) {
	b.moveType = moveType
}
func (b BattleAttackPokemon) SetMoveSpecies(moveSpecies value.MoveSpecies) {
	b.moveSpecies = moveSpecies
}
func (b BattleAttackPokemon) SetMovePower(movePower int) {
	b.movePower = movePower
}
func (b BattleAttackPokemon) SetIsMoveContained(isMoveContained bool) {
	b.isMoveContained = isMoveContained
}
func (b BattleAttackPokemon) SetCanMoveGuard(canMoveGuard bool) {
	b.canMoveGuard = canMoveGuard
}

func (p BattleAttackPokemon) MoveType() value.PokemonType {
	return p.moveType
}

func (p BattleAttackPokemon) ResolvePowerValue() int {
	correctionTarget := value.PhysicalMoveCorrection
	if p.moveSpecies.IsSpecial() {
		correctionTarget = value.SpecialMoveCorrection
	}
	powerValue := p.additionalEffects.ApplyCorrection(correctionTarget, float32(p.movePower))
	return int(powerValue)
}

func (p BattleAttackPokemon) ResolveAttackValue
