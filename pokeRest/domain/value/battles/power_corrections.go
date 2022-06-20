package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type PowerCorrections struct {
	targets []CorrectionTarget
	side    BattleSideType
	*BattleCorrectionValues
}

func NewPowerCorrections(values *BattleCorrectionValues) *PowerCorrections {
	targets := GetPowerCorrectionTargets()
	return &PowerCorrections{
		targets:                targets,
		side:                   BattleAttackSide,
		BattleCorrectionValues: values.get(targets...),
	}
}

func (c PowerCorrections) SupplyPowerCorrectionApplier(
	species value.MoveSpecies, data IPokemonBattleDataSet) correctionApplier {
	target := CorrectionNone
	if species == value.MoveSpeciesPhysical {
		target = CorrectionPhysicalPower
	}
	if species == value.MoveSpeciesSpecial {
		target = CorrectionSpecialPower
	}
	return func(value uint16) uint16 {
		return c.Apply(value, target, data, c.side)
	}
}
