package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

var _ IDamageClacElements = (*DamageCalcElements)(nil)

type DamageCalcElements struct {
	dataset battles.IPokemonBattleDataSet
}

func NewDamageCalcElements(dataset *PokemonBattleDataSet) *DamageCalcElements {
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

func (e DamageCalcElements) IsTypeMatchAttackSide() bool {
	return e.dataset.IsTypeMatchAttackSide()
}

func (e DamageCalcElements) IsBurnAttackSide() bool {
	return e.dataset.IsBurnAttackSide()
}

func (e DamageCalcElements) AbnormalStateAttackSideCorectedValue() uint16 {
	return e.dataset.AbnormalStateAttackSideCorectedValue()
}

func (e DamageCalcElements) TypeCompatibilityDamageRate() float64 {
	return float64(e.dataset.TypeCompatibilityDamageRate())
}
