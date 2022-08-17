package battles

type PowerCorrections struct {
	side        BattleSideType
	corrections *BattleCorrectionValues
}

func NewPowerCorrections(values *BattleCorrectionValues) *PowerCorrections {
	targets := GetPowerCorrectionTargets()
	return &PowerCorrections{
		side:        BattleAttackSide,
		corrections: values.get(targets...),
	}
}

func (c PowerCorrections) Apply(value uint16, dataset IPokemonBattleDataSet) uint16 {
	if c.corrections.IsEmpty() {
		return value
	}
	target := CorrectionNone
	if dataset.MoveSpecies().IsPhysical() {
		target = CorrectionPhysicalPower
	}
	if dataset.MoveSpecies().IsSpecial() {
		target = CorrectionSpecialPower
	}
	return c.corrections.Apply(value, target, dataset, c.side)
}
