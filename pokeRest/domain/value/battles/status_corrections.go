package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

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
	return c.correctionValues.Apply(dataset.AttackPokemonActualValues().A(), CorrectionStatusA, dataset, c.side)
}

func (c StatusCorrections) ApplyB(dataset IPokemonBattleDataSet) uint16 {
	return c.correctionValues.Apply(dataset.AttackPokemonActualValues().B(), CorrectionStatusB, dataset, c.side)
}

func (c StatusCorrections) ApplyC(dataset IPokemonBattleDataSet) uint16 {
	return c.correctionValues.Apply(dataset.AttackPokemonActualValues().C(), CorrectionStatusC, dataset, c.side)
}

func (c StatusCorrections) ApplyD(dataset IPokemonBattleDataSet) uint16 {
	return c.correctionValues.Apply(dataset.AttackPokemonActualValues().D(), CorrectionStatusD, dataset, c.side)
}

func (c StatusCorrections) ApplyS(dataset IPokemonBattleDataSet) uint16 {
	return c.correctionValues.Apply(dataset.AttackPokemonActualValues().S(), CorrectionStatusS, dataset, c.side)
}
