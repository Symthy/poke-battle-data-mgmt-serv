package battles

import (
	"github.com/Symthy/PokeRest/pokeRest/common/lists"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type correctionApplier func(value uint16) uint16

type correctionsApplier[T any] func(values T) T

func identify(value uint16) uint16 {
	return value
}

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

type BattleCorrectionValues struct {
	values  []*BattleCorrectionValue
	targets []CorrectionTarget
}

func NewBattleCorrectionValues(values ...*BattleCorrectionValue) *BattleCorrectionValues {
	if len(values) == 0 {
		return &BattleCorrectionValues{
			values:  []*BattleCorrectionValue{},
			targets: []CorrectionTarget{},
		}
	}
	return &BattleCorrectionValues{
		values: values,
		targets: lists.Map(values, func(v *BattleCorrectionValue) CorrectionTarget {
			return v.target
		}),
	}
}

func (b BattleCorrectionValues) get(targets ...CorrectionTarget) *BattleCorrectionValues {
	values := lists.Filter(b.values, func(correctionValue *BattleCorrectionValue) bool {
		return correctionValue.AnyEqualTarget(targets...)
	})
	return NewBattleCorrectionValues(values...)
}

func (c BattleCorrectionValues) Apply(
	value uint16, target CorrectionTarget, data IPokemonBattleDataSet, side BattleSideType) uint16 {
	result := value
	for _, correction := range c.values {
		if target != correction.target {
			continue
		}
		result = correction.Apply(result, data, side)
	}
	return result
}

func (c *BattleCorrectionValues) Merge(corrections *BattleCorrectionValues) {
	c.values = append(c.values, corrections.values...)
	c.targets = append(c.targets, corrections.targets...)
}
