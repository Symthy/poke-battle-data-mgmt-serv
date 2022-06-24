package battles

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

func (c MovePowerCorrections) Apply(value uint16, dataset IPokemonBattleDataSet) uint16 {
	if c.corrections.IsEmpty() {
		return value
	}
	target := CorrectionNone
	if dataset.MoveSpecies().IsPhysical() {
		target = CorrectionPhysicalMove
	}
	if dataset.MoveSpecies().IsSpecial() {
		target = CorrectionSpecialMove
	}
	return c.corrections.Apply(value, target, dataset, c.side)
}
