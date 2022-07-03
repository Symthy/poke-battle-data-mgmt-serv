package battles

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type StatusCorrections struct {
	targets          []CorrectionTarget
	side             BattleSideType
	correctionValues *BattleCorrectionValues
	applier          correctionsApplier[*value.PokemonActualValues]
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

func (c StatusCorrections) applyCorrection(actualValue uint16, correctionTarget CorrectionTarget, dataset IPokemonBattleDataSet) uint16 {
	correctionValue := c.correctionValues.Apply(4096, correctionTarget, dataset, c.side)
	return fmath.RoundUpIfDecimalGreaterFive[uint16](float64(actualValue*correctionValue) / 4096.0)
}
