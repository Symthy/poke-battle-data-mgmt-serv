package battles

import "github.com/Symthy/PokeRest/internal/pkg/fmath"

type MovePowerCorrections struct {
	side        BattleSideType
	corrections *BattleCorrectionValues
}

func NewMovePowerCorrections(values *BattleCorrectionValues) *MovePowerCorrections {
	return &MovePowerCorrections{
		side:        BattleAttackSide,
		corrections: values.get(GetMovePowerCorrectionTargets()...),
	}
}

func (c MovePowerCorrections) Apply(movePower uint16, dataset IPokemonBattleDataSet) uint16 {
	if c.corrections.IsEmpty() || movePower == 0 {
		return movePower
	}
	target := CorrectionNone
	if dataset.MoveSpecies().IsPhysical() {
		target = CorrectionPhysicalMove
	}
	if dataset.MoveSpecies().IsSpecial() {
		target = CorrectionSpecialMove
	}
	correctionValue := c.corrections.Apply(4096, target, dataset, c.side)
	correctedMovePower := fmath.RoundUpIfDecimalGreaterFive[uint16](float64(movePower*correctionValue) / 4096.0)
	if correctedMovePower < 1 {
		return 1
	}
	return correctedMovePower
}
