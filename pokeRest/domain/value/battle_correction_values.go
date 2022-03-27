package value

import "github.com/Symthy/PokeRest/pokeRest/common/lists"

func identify(value float32) float32 {
	return value
}

type StatusCorrections struct {
	targets []CorrectionTarget
	BattleCorrectionValues
}

func NewStatusCorrections(values BattleCorrectionValues) StatusCorrections {
	targets := GetStatusCorrectionTargets()
	return StatusCorrections{
		targets:                targets,
		BattleCorrectionValues: values.get(targets...),
	}
}

func (c StatusCorrections) SupplyStatusCorrectionApplier(param PokemonParam) func(value float32) float32 {
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
		return c.Apply(value, target)
	}
}

type PowerCorrections struct {
	targets []CorrectionTarget
	BattleCorrectionValues
}

func NewPowerCorrections(values BattleCorrectionValues) PowerCorrections {
	targets := GetPowerCorrectionTargets()
	return PowerCorrections{
		BattleCorrectionValues: values.get(targets...),
	}
}

func (c PowerCorrections) SupplyPowerCorrectionApplier(species MoveSpecies) func(value float32) float32 {
	target := NoneCorrection
	if species == Physical {
		target = PhysicalPowerCorrection
	}
	if species == Special {
		target = SpecialPowerCorrection
	}
	return func(value float32) float32 {
		return c.Apply(value, target)
	}
}

type MovePowerCorrections struct {
	targets []CorrectionTarget
	BattleCorrectionValues
}

func NewMovePowerCorrections(values BattleCorrectionValues) MovePowerCorrections {
	return MovePowerCorrections{
		BattleCorrectionValues: values.get(GetMovePowerCorrectionTargets()...),
	}
}

func (c MovePowerCorrections) SupplyMovePowerCorrectionApplier(species MoveSpecies) func(value float32) float32 {
	return func(value float32) float32 {
		if species == Physical {
			return c.Apply(value, PhysicalMoveCorrection)
		}
		if species == Special {
			return c.Apply(value, SpecialMoveCorrection)
		}
		return identify(value)
	}
}

type DamageCorrections struct {
	targets []CorrectionTarget
	BattleCorrectionValues
}

func NewDamageCorrections(values BattleCorrectionValues) DamageCorrections {
	return DamageCorrections{
		BattleCorrectionValues: values.get(GetDamageCorrectionTargets()...),
	}
}

func (c DamageCorrections) SupplyDamageCorrectionApplier() func(value float32) float32 {
	return func(value float32) float32 {
		return c.Apply(value, DamageCorrection)
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

func (c BattleCorrectionValues) Apply(value float32, target CorrectionTarget) float32 {
	result := value
	for _, correction := range c.items {
		if target != correction.target {
			continue
		}
		result = correction.Apply(result)
	}
	return result
}

func (c *BattleCorrectionValues) Merge(corrections BattleCorrectionValues) {
	c.items = append(c.items, corrections.items...)
	c.targets = append(c.targets, corrections.targets...)
}
