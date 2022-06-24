package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

type DamageCalcElements struct {
	dataset battles.IPokemonBattleDataSet
}

func NewDamageCalcElements(dataset *PokemonBattleDataSet, attackEffects *AttackSideBattleEffects,
	defenseEffects *DefenseSideBattleEffects) *DamageCalcElements {
	return &DamageCalcElements{
		dataset: dataset,
	}
}

// for testing
func (e *DamageCalcElements) GetPokemonBattleDataSet() battles.IPokemonBattleDataSet {
	return e.dataset
}

func (e DamageCalcElements) ResolvePowerValue() uint16 {
	attackPower := e.resolveAttackPowerValue()
	movePower := e.resolveMovePowerValue()
	return fmath.RoundUpIfDecimalGreaterFive[uint16](float64(movePower*attackPower) / 4096)
}

func (e DamageCalcElements) resolveMovePowerValue() uint16 {
	return e.dataset.MovePowerValue()
}

func (e DamageCalcElements) resolveAttackPowerValue() uint16 {
	return e.dataset.AttackPowerValue()
}

func (e DamageCalcElements) ResolveAttackActualValue() uint16 {
	return e.dataset.AttackCorrectedActualValue()
}

func (e DamageCalcElements) ResolveDefenseActualValue() uint16 {
	return e.dataset.DefenseCorrectedActualValue()
}
