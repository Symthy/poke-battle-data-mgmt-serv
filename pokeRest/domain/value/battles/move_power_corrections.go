package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type MovePowerCorrections struct {
	targets []CorrectionTarget
	side    BattleSideType
	*BattleCorrectionValues
}

func NewMovePowerCorrections(values *BattleCorrectionValues) *MovePowerCorrections {
	targets := GetMovePowerCorrectionTargets()
	return &MovePowerCorrections{
		targets:                targets,
		side:                   BattleAttackSide,
		BattleCorrectionValues: values.get(GetMovePowerCorrectionTargets()...),
	}
}

func (c MovePowerCorrections) SupplyMovePowerCorrectionApplier(
	species value.MoveSpecies, data IPokemonBattleDataSet) correctionApplier {
	return func(val uint16) uint16 {
		if species == value.MoveSpeciesPhysical {
			return c.Apply(val, CorrectionPhysicalMove, data, c.side)
		}
		if species == value.MoveSpeciesSpecial {
			return c.Apply(val, CorrectionSpecialMove, data, c.side)
		}
		return identify(val)
	}
}
