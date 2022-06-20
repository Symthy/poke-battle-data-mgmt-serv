package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

type DamageCalcElements struct {
	dataSet        *PokemonBattleDataSet
	attackEffects  *AttackSideBattleEffects
	defenseEffects *DefenseSideBattleEffects
}

func NewDamageCalcElements(dataSet *PokemonBattleDataSet, attackEffects *AttackSideBattleEffects,
	defenseEffects *DefenseSideBattleEffects) *DamageCalcElements {
	return &DamageCalcElements{
		dataSet:        dataSet,
		attackEffects:  attackEffects,
		defenseEffects: defenseEffects,
	}
}

// for testing
func (e *DamageCalcElements) GetPokemonBattleDataSet() *PokemonBattleDataSet {
	return e.dataSet
}

func (e DamageCalcElements) ResolvePowerValue() uint16 {
	attackPower := e.resolveAttackPowerValue()
	movePower := e.resolveMovePowerValue()

	return fmath.RoundUpIfDecimalGreaterFive[uint16](float64(movePower*attackPower) / 4096)
}

func (e DamageCalcElements) resolveMovePowerValue() uint16 {
	applier := e.attackEffects.SupplyMovePowerCorrectionApplier(e.dataSet.species, e.dataSet)
	return applier(e.dataSet.AttackMove.power)
}

func (e DamageCalcElements) resolveAttackPowerValue() uint16 {
	applier := e.attackEffects.SupplyPowerCorrectionApplier(e.dataSet.species, e.dataSet)
	return applier(4096)
}

func (e DamageCalcElements) ResolveAttackValue() uint16 {
	if e.dataSet.species.IsPhysical() {
		// Todo: refactor method
		valueA := e.dataSet.AttackSidePokemon.attackPokemonActualValues.A()
		return e.attackEffects.StatusCorrections.Apply(valueA, battles.CorrectionStatusA, e.dataSet, battles.BattleAttackSide)
	}
	if e.dataSet.species.IsSpecial() {
		// Todo: refactor method
		valueA := e.dataSet.AttackSidePokemon.attackPokemonActualValues.C()
		return e.attackEffects.StatusCorrections.Apply(valueA, battles.CorrectionStatusB, e.dataSet, battles.BattleAttackSide)
	}
	return 0
}

func (e DamageCalcElements) ResolveDefenseValue() uint16 {
	if e.dataSet.species.IsPhysical() {
		// Todo: refactor method
		valueB := e.dataSet.AttackSidePokemon.attackPokemonActualValues.B()
		return e.attackEffects.StatusCorrections.Apply(valueB, battles.CorrectionStatusB, e.dataSet, battles.BattleDefenseSide)
	}
	if e.dataSet.species.IsSpecial() {
		// Todo: refactor method
		valueD := e.dataSet.AttackSidePokemon.attackPokemonActualValues.D()
		return e.attackEffects.StatusCorrections.Apply(valueD, battles.CorrectionStatusD, e.dataSet, battles.BattleDefenseSide)
	}
	return 0
}
