package battles

type DamageCorrections struct {
	targets []CorrectionTarget
	side    BattleSideType
	BattleCorrectionValues
}

func NewDamageCorrections(values *BattleCorrectionValues, side BattleSideType) *DamageCorrections {
	return &DamageCorrections{
		BattleCorrectionValues: *values.get(GetDamageCorrectionTargets()...),
	}
}

func (c DamageCorrections) SupplyDamageCorrectionApplier(
	data IPokemonBattleDataSet) correctionApplier {
	return func(value uint16) uint16 {
		return c.Apply(value, CorrectionDamage, data, c.side)
	}
}
