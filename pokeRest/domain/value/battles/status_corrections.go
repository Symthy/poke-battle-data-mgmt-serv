package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type StatusCorrections struct {
	targets []CorrectionTarget
	side    BattleSideType
	*BattleCorrectionValues
}

func NewStatusCorrections(values *BattleCorrectionValues, side BattleSideType) *StatusCorrections {
	targets := GetStatusCorrectionTargets()
	return &StatusCorrections{
		targets:                targets,
		side:                   side,
		BattleCorrectionValues: values.get(targets...),
	}
}

func (c StatusCorrections) SupplyAllStatusCorrectionApplier(
	data IPokemonBattleDataSet) correctionsApplier[*value.PokemonActualValues] {
	return func(actualValues *value.PokemonActualValues) *value.PokemonActualValues {
		aVal := c.Apply(actualValues.A(), CorrectionStatusA, data, c.side)
		bVal := c.Apply(actualValues.B(), CorrectionStatusB, data, c.side)
		cVal := c.Apply(actualValues.C(), CorrectionStatusC, data, c.side)
		dVal := c.Apply(actualValues.D(), CorrectionStatusD, data, c.side)
		sVal := c.Apply(actualValues.S(), CorrectionStatusS, data, c.side)
		return value.NewPokemonActualValues(actualValues.H(), aVal, bVal, cVal, dVal, sVal)
	}
}
