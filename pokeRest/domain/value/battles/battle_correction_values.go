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
	data IPokemonBattleDataSet) correctionsApplier[value.PokemonActualValues] {
	return func(actualValues value.PokemonActualValues) value.PokemonActualValues {
		aVal := c.Apply(actualValues.A(), AttackCorrection, data, c.side)
		bVal := c.Apply(actualValues.B(), DefenseCorrection, data, c.side)
		cVal := c.Apply(actualValues.C(), SpecialAttackCorrection, data, c.side)
		dVal := c.Apply(actualValues.D(), SpecialDefenseCorrection, data, c.side)
		sVal := c.Apply(actualValues.S(), SpeedCorrection, data, c.side)
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
	target := NoneCorrection
	if species == value.PhysicalMove {
		target = PhysicalPowerCorrection
	}
	if species == value.SpecialMove {
		target = SpecialPowerCorrection
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
		if species == value.PhysicalMove {
			return c.Apply(val, PhysicalMoveCorrection, data, c.side)
		}
		if species == value.SpecialMove {
			return c.Apply(val, SpecialMoveCorrection, data, c.side)
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
		return c.Apply(value, DamageCorrection, data, c.side)
	}
}

type BattleCorrectionValues struct {
	targets []CorrectionTarget
	items   []*BattleCorrectionValue
}

func NewBattleCorrectionValues(items ...*BattleCorrectionValue) *BattleCorrectionValues {
	targets := []CorrectionTarget{}
	for _, correction := range items {
		targets = append(targets, correction.target)
	}
	return &BattleCorrectionValues{
		targets: targets,
		items:   items,
	}
}

func (b BattleCorrectionValues) get(targets ...CorrectionTarget) *BattleCorrectionValues {
	values := lists.Filter(b.items, func(correctionValue *BattleCorrectionValue) bool {
		return correctionValue.AnyEqualTarget(targets...)
	})
	return NewBattleCorrectionValues(values...)
}

func (c BattleCorrectionValues) Apply(
	value uint16, target CorrectionTarget, data IPokemonBattleDataSet, side BattleSideType) uint16 {
	result := value
	for _, correction := range c.items {
		if target != correction.target {
			continue
		}
		result = correction.Apply(result, data, side)
	}
	return result
}

func (c *BattleCorrectionValues) Merge(corrections *BattleCorrectionValues) {
	c.items = append(c.items, corrections.items...)
	c.targets = append(c.targets, corrections.targets...)
}
