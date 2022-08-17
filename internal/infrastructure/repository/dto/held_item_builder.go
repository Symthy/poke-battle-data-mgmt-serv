package dto

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/items"
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

var _ items.IHeldItemNotification = (*HeldItemSchemaBuilder)(nil)

type HeldItemSchemaBuilder struct {
	*schema.HeldItem
}

func NewHeldItemSchemaBuilder() *HeldItemSchemaBuilder {
	return &HeldItemSchemaBuilder{&schema.HeldItem{}}
}

func (b *HeldItemSchemaBuilder) SetId(id identifier.HeldItemId) {
	b.ID = id.Value()
}
func (b *HeldItemSchemaBuilder) SetName(name string) {
	b.Name = name
}
func (b *HeldItemSchemaBuilder) SetDescription(description string) {
	b.Description = description
}
func (b *HeldItemSchemaBuilder) SetBattleEffects(effects *battles.BattleEffects) {
	b.BattleEffects = toSchemaBattleEffects(effects)
}

func (b HeldItemSchemaBuilder) Build() *schema.HeldItem {
	return b.HeldItem
}
