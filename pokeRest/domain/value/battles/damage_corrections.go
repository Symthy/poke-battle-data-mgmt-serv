package battles

type DamageCorrections struct {
	side        BattleSideType
	corrections *BattleCorrectionValues
}

func NewDamageCorrections(values *BattleCorrectionValues, side BattleSideType) *DamageCorrections {
	return &DamageCorrections{
		side:        side,
		corrections: values.get(GetDamageCorrectionTargets()...),
	}
}

func (c DamageCorrections) Apply(value uint16, dataset IPokemonBattleDataSet) uint16 {
	return c.corrections.Apply(value, CorrectionDamage, dataset, c.side)
}
