package dto

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ parties.IPartyNotification = (*PartySchemaBuilder)(nil)

type PartySchemaBuilder struct {
	*schema.Party
}

func NewPartySchemaBuilder() *PartySchemaBuilder {
	return &PartySchemaBuilder{&schema.Party{}}
}

func (b *PartySchemaBuilder) SetId(id identifier.PartyId) {
	b.ID = id.Value()
}
func (b *PartySchemaBuilder) SetName(name string) {
	b.Name = name
}
func (b *PartySchemaBuilder) SetBattleFormat(format value.BattleFormat) {
	b.BattleFormat = enum.BattleFormat(format.String())
}
func (b *PartySchemaBuilder) SetIsPrivate(isPrivate bool) {
	b.IsPrivate = isPrivate
}

func (b PartySchemaBuilder) Build() *schema.Party {
	return b.Party
}
