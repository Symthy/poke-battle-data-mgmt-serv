package dto

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ parties.IPartyTagNotification = (*PartyTagSchemaBuilder)(nil)

type PartyTagSchemaBuilder struct {
	schema.PartyTag
}

func (b *PartyTagSchemaBuilder) SetId(id identifier.PartyTagId) {
	b.ID = id.Value()
}
func (b *PartyTagSchemaBuilder) SetName(value string) {
	b.Name = value
}
func (b *PartyTagSchemaBuilder) SetIsGeneration(value bool) {
	b.IsGeneration = &value
}
func (b *PartyTagSchemaBuilder) SetIsSeason(value bool) {
	b.IsSeason = &value
}

func (b PartyTagSchemaBuilder) Build() schema.PartyTag {
	return b.PartyTag
}
