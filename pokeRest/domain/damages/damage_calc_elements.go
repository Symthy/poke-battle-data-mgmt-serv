package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

var _ IDamageClacElements = (*DamageCalcElements)(nil)

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

func (e DamageCalcElements) IsNoDamage() bool {
	return e.dataset.IsNoDamage()
}

func (e DamageCalcElements) FinalAttackValue() uint16 {
	attackPower := e.resolvePowerCorrectedValue()
	attackPower = fmath.Round[uint16](float64(attackPower*e.FieldCorrectedValue()) / 4096.0)
	movePower := e.resolveMovePowerValue()
	return fmath.RoundUpIfDecimalGreaterFive[uint16](float64(movePower*attackPower) / 4096.0)
}

func (e DamageCalcElements) resolveMovePowerValue() uint16 {
	return e.dataset.MovePowerValue()
}

func (e DamageCalcElements) resolvePowerCorrectedValue() uint16 {
	return e.dataset.PowerCorrectedValue()
}

func (e DamageCalcElements) AttackActualValue() uint16 {
	return e.dataset.AttackCorrectedActualValue()
}

func (e DamageCalcElements) DefenseActualValue() uint16 {
	return e.dataset.DefenseCorrectedActualValue()
}

func (e DamageCalcElements) DamageCorrectedValue() uint16 {
	return e.dataset.DamageCorrectedValue()
}

func (e DamageCalcElements) FieldCorrectedValue() uint16 {
	return e.dataset.FieldCorrectedValue()
}

func (e DamageCalcElements) WeatherCorrectedValue() uint16 {
	return e.dataset.WeatherCorrectedValue()
}

func (e DamageCalcElements) IsTypeMatch() bool {
	return e.dataset.IsTypeMatch()
}

func (e DamageCalcElements) TypeCompatibilityDamageRate() float64 {
	return float64(e.dataset.TypeCompatibilityDamageRate())
}
