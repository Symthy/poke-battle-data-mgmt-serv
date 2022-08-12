package dto

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ abilities.IAbilityNotification = (*AbilitySchemaBuilder)(nil)

type AbilitySchemaBuilder struct {
	*schema.Ability
}

func NewAbilitySchemaBuilder() *AbilitySchemaBuilder {
	return &AbilitySchemaBuilder{&schema.Ability{}}
}

func (b *AbilitySchemaBuilder) SetId(id identifier.AbilityId) {
	b.ID = id.Value()
}
func (b *AbilitySchemaBuilder) SetName(name string) {
	b.Name = name
}
func (b *AbilitySchemaBuilder) SetDescription(description string) {
	b.Description = description
}
func (b *AbilitySchemaBuilder) SetBattleEffects(effects *battles.BattleEffects) {
	b.BattleEffects = toSchemaBattleEffects(effects)
}

func (b AbilitySchemaBuilder) Build() *schema.Ability {
	return b.Ability
}
