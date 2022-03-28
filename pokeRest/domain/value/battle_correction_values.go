package value

import "github.com/Symthy/PokeRest/pokeRest/common/lists"

type correctionApplier func(value float32) float32

func identify(value float32) float32 {
	return value
}

type StatusCorrections struct {
	targets []CorrectionTarget
	side    BattleSideType
	BattleCorrectionValues
}

func NewStatusCorrections(values BattleCorrectionValues, side BattleSideType) StatusCorrections {
	targets := GetStatusCorrectionTargets()
	return StatusCorrections{
		targets:                targets,
		side:                   side,
		BattleCorrectionValues: values.get(targets...),
	}
}

func (c StatusCorrections) SupplyStatusCorrectionApplier(
	param PokemonParam, data IPokemonBattleDataSet) correctionApplier {
	target := NoneCorrection
	if param == A {
		target = AttackCorrection
	}
	if param == B {
		target = DefenceCorrection
	}
	if param == C {
		target = SpecialAttackCorrection
	}
	if param == D {
		target = SpecialDefenceCorrection
	}
	if param == S {
		target = SpeedCorrection
	}
	return func(value float32) float32 {
		return c.Apply(value, target, data, c.side)
	}
}

type PowerCorrections struct {
	targets []CorrectionTarget
	side    BattleSideType
	BattleCorrectionValues
}

func NewPowerCorrections(values BattleCorrectionValues) PowerCorrections {
	targets := GetPowerCorrectionTargets()
	return PowerCorrections{
		targets:                targets,
		side:                   BattleAttackSide,
		BattleCorrectionValues: values.get(targets...),
	}
}

func (c PowerCorrections) SupplyPowerCorrectionApplier(
	species MoveSpecies, data IPokemonBattleDataSet) correctionApplier {
	target := NoneCorrection
	if species == Physical {
		target = PhysicalPowerCorrection
	}
	if species == Special {
		target = SpecialPowerCorrection
	}
	return func(value float32) float32 {
		return c.Apply(value, target, data, c.side)
	}
}

type MovePowerCorrections struct {
	targets []CorrectionTarget
	side    BattleSideType
	BattleCorrectionValues
}

func NewMovePowerCorrections(values BattleCorrectionValues) MovePowerCorrections {
	targets := GetMovePowerCorrectionTargets()
	return MovePowerCorrections{
		targets:                targets,
		side:                   BattleAttackSide,
		BattleCorrectionValues: values.get(GetMovePowerCorrectionTargets()...),
	}
}

func (c MovePowerCorrections) SupplyMovePowerCorrectionApplier(
	species MoveSpecies, data IPokemonBattleDataSet) correctionApplier {
	return func(value float32) float32 {
		if species == Physical {
			return c.Apply(value, PhysicalMoveCorrection, data, c.side)
		}
		if species == Special {
			return c.Apply(value, SpecialMoveCorrection, data, c.side)
		}
		return identify(value)
	}
}

type DamageCorrections struct {
	targets []CorrectionTarget
	side    BattleSideType
	BattleCorrectionValues
}

func NewDamageCorrections(values BattleCorrectionValues, side BattleSideType) DamageCorrections {
	return DamageCorrections{
		BattleCorrectionValues: values.get(GetDamageCorrectionTargets()...),
	}
}

func (c DamageCorrections) SupplyDamageCorrectionApplier(
	data IPokemonBattleDataSet) correctionApplier {
	return func(value float32) float32 {
		return c.Apply(value, DamageCorrection, data, c.side)
	}
}

type BattleCorrectionValues struct {
	targets []CorrectionTarget
	items   []BattleCorrectionValue
}

func NewBattleCorrectionValues(items []BattleCorrectionValue) BattleCorrectionValues {
	targets := []CorrectionTarget{}
	for _, correction := range items {
		targets = append(targets, correction.target)
	}
	return BattleCorrectionValues{
		targets: targets,
		items:   items,
	}
}

func (b BattleCorrectionValues) get(targets ...CorrectionTarget) BattleCorrectionValues {
	values := lists.Filter(b.items, func(value BattleCorrectionValue) bool {
		return value.AnyEqualTarget(targets...)
	})
	return NewBattleCorrectionValues(values)
}

func (c BattleCorrectionValues) Apply(
	value float32, target CorrectionTarget, data IPokemonBattleDataSet, side BattleSideType) float32 {
	result := value
	for _, correction := range c.items {
		if target != correction.target {
			continue
		}
		result = correction.Apply(result, data, side)
	}
	return result
}

func (c *BattleCorrectionValues) Merge(corrections BattleCorrectionValues) {
	c.items = append(c.items, corrections.items...)
	c.targets = append(c.targets, corrections.targets...)
}
