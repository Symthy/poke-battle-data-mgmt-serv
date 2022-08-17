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

func (c DamageCorrections) Apply(dataset IPokemonBattleDataSet) uint16 {
	return c.corrections.Apply(4096, CorrectionDamage, dataset, c.side)
}
