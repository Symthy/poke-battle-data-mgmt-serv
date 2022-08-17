package battles

type StatusCorrections struct {
	targets          []CorrectionTarget
	side             BattleSideType
	correctionValues *BattleCorrectionValues
}

func NewStatusCorrections(values *BattleCorrectionValues, side BattleSideType) *StatusCorrections {
	targets := GetStatusCorrectionTargets()
	correctionValuse := values.get(targets...)
	return &StatusCorrections{
		targets:          targets,
		side:             side,
		correctionValues: correctionValuse,
	}
}

func (c StatusCorrections) ApplyA(dataset IPokemonBattleDataSet) uint16 {
	return c.applyCorrection(dataset.AttackPokemonActualValues().A(), CorrectionStatusA, dataset)
}

func (c StatusCorrections) ApplyB(dataset IPokemonBattleDataSet) uint16 {
	return c.applyCorrection(dataset.AttackPokemonActualValues().B(), CorrectionStatusB, dataset)
}

func (c StatusCorrections) ApplyC(dataset IPokemonBattleDataSet) uint16 {
	return c.applyCorrection(dataset.AttackPokemonActualValues().C(), CorrectionStatusC, dataset)
}

func (c StatusCorrections) ApplyD(dataset IPokemonBattleDataSet) uint16 {
	return c.applyCorrection(dataset.AttackPokemonActualValues().D(), CorrectionStatusD, dataset)
}

func (c StatusCorrections) ApplyS(dataset IPokemonBattleDataSet) uint16 {
	return c.applyCorrection(dataset.AttackPokemonActualValues().S(), CorrectionStatusS, dataset)
}

func (c StatusCorrections) applyCorrection(value uint16, correctionTarget CorrectionTarget, dataset TriggerConditionParams) uint16 {
	return c.correctionValues.Apply(value, correctionTarget, dataset, c.side)
}
