package dto

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
	"github.com/Symthy/PokeRest/internal/pkg/lists"
)

func toSchemaBattleEffects(effects *battles.BattleEffects) *mixin.BattleEffects {
	schemaCorrections := lists.Map(
		effects.Corrections().ToSlice(),
		func(corrction *battles.BattleCorrectionValue) *mixin.Correction {
			builder := NewCorrectionValueSchemaBuilder()
			corrction.Notify(builder)
			return builder.Build()
		},
	)
	schemaOverrides := lists.Map(
		effects.Overrides().ToSlice(),
		func(override *battles.BattleOverrideValue) *mixin.Override {
			builder := NewOverrideValueSchemaBuilder()
			override.Notify(builder)
			return builder.Build()
		},
	)

	return &mixin.BattleEffects{
		Corrections: schemaCorrections,
		Overrides:   schemaOverrides,
	}
}

var _ battles.ICorrectionValueNotification = (*CorrectionValueSchemaBuilder)(nil)

type CorrectionValueSchemaBuilder struct {
	*mixin.Correction
	*TriggerConditionSchemaBuilder
}

func NewCorrectionValueSchemaBuilder() *CorrectionValueSchemaBuilder {
	return &CorrectionValueSchemaBuilder{
		Correction:                    &mixin.Correction{},
		TriggerConditionSchemaBuilder: NewTriggerConditionSchemaBuilder(),
	}
}
func (c *CorrectionValueSchemaBuilder) SetTarget(target battles.CorrectionTarget) {
	c.Target = enum.CorrectionTarget(target.ToString())
}
func (c *CorrectionValueSchemaBuilder) SetCorrectionValue(value uint16) {
	c.Value = value
}
func (c *CorrectionValueSchemaBuilder) SetDecimalCalcType(decimalCalcType battles.DecimalCalcType) {
	// Todo
}
func (c CorrectionValueSchemaBuilder) Build() *mixin.Correction {
	c.Correction.TriggerCondition = c.TriggerConditionSchemaBuilder.Build()
	return c.Correction
}

var _ battles.IOverrideValueNotification = (*OverrideValueSchemaBuilder)(nil)

type OverrideValueSchemaBuilder struct {
	*mixin.Override
	*TriggerConditionSchemaBuilder
}

func NewOverrideValueSchemaBuilder() *OverrideValueSchemaBuilder {
	return &OverrideValueSchemaBuilder{
		Override:                      &mixin.Override{},
		TriggerConditionSchemaBuilder: NewTriggerConditionSchemaBuilder(),
	}
}

func (o *OverrideValueSchemaBuilder) SetTarget(target battles.OverrideTarget) {
	o.Target = enum.OverrideTarget(target.ToString())
}
func (o *OverrideValueSchemaBuilder) SetOverrideValue(value string) {
	o.Value = value
}
func (o OverrideValueSchemaBuilder) Build() *mixin.Override {
	o.Override.TriggerCondition = o.TriggerConditionSchemaBuilder.Build()
	return o.Override
}

var _ battles.ITriggerConditionNotification = (*TriggerConditionSchemaBuilder)(nil)

type TriggerConditionSchemaBuilder struct {
	*mixin.TriggerCondition
}

func NewTriggerConditionSchemaBuilder() *TriggerConditionSchemaBuilder {
	return &TriggerConditionSchemaBuilder{
		TriggerCondition: &mixin.TriggerCondition{},
	}
}
func (t *TriggerConditionSchemaBuilder) SetEntry(entry battles.ConditionEntry) {
	t.Entry = enum.ConditionEntry(entry.ToString())
}

func (t *TriggerConditionSchemaBuilder) SetConditionValue(value string) {
	t.Value = value
}

func (t TriggerConditionSchemaBuilder) Build() *mixin.TriggerCondition {
	return t.TriggerCondition
}
