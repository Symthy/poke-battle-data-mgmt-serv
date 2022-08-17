package battles

import (
	"github.com/Symthy/PokeRest/internal/domain/value"
)

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
	actualValues := c.resolveActualValues(dataset)
	return c.applyCorrection(actualValues.A(), CorrectionStatusA, dataset)
}

func (c StatusCorrections) ApplyB(dataset IPokemonBattleDataSet) uint16 {
	actualValues := c.resolveActualValues(dataset)
	return c.applyCorrection(actualValues.B(), CorrectionStatusB, dataset)
}

func (c StatusCorrections) ApplyC(dataset IPokemonBattleDataSet) uint16 {
	actualValues := c.resolveActualValues(dataset)
	return c.applyCorrection(actualValues.C(), CorrectionStatusC, dataset)
}

func (c StatusCorrections) ApplyD(dataset IPokemonBattleDataSet) uint16 {
	actualValues := c.resolveActualValues(dataset)
	return c.applyCorrection(actualValues.D(), CorrectionStatusD, dataset)
}

func (c StatusCorrections) ApplyS(dataset IPokemonBattleDataSet) uint16 {
	actualValues := c.resolveActualValues(dataset)
	return c.applyCorrection(actualValues.S(), CorrectionStatusS, dataset)
}

func (c StatusCorrections) applyCorrection(value uint16, correctionTarget CorrectionTarget, dataset TriggerConditionParams) uint16 {
	return c.correctionValues.Apply(value, correctionTarget, dataset, c.side)
}

func (c StatusCorrections) resolveActualValues(dataset IPokemonBattleDataSet) *value.PokemonActualValues {
	if c.side == BattleAttackSide {
		return dataset.AttackPokemonActualValues()
	}
	return dataset.DefensePokemonActualValues()
}
