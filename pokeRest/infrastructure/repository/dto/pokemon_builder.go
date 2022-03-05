package dto

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ pokemons.IPokemonNotification = (*PokemonSchemaBuilder)(nil)

type PokemonSchemaBuilder struct {
	schema.Pokemon
}

func (b *PokemonSchemaBuilder) SetId(id identifier.PokemonId) {
	b.ID = id.Value()
}
func (b *PokemonSchemaBuilder) SetPokedexNo(value int) {
	b.PokedexNo = value
}
func (b *PokemonSchemaBuilder) SetFormNo(value int) {
	b.FormNo = value
}
func (b *PokemonSchemaBuilder) SetFormName(value string) {
	b.FormName = value
}
func (b *PokemonSchemaBuilder) SetName(value string) {
	b.Name = value
}
func (b *PokemonSchemaBuilder) SetEnglishName(value string) {
	b.EnglishName = value
}
func (b *PokemonSchemaBuilder) SetGeneration(value int) {
	b.Generation = value
}
func (b *PokemonSchemaBuilder) SetTypeSet(value value.PokemonTypeSet) {
	type1, type2 := value.GetTypeEnglishNames()
	b.Type1 = enum.PokemonType(type1)
	b.Type2 = enum.PokemonType(type2)
}
func (b *PokemonSchemaBuilder) SetAbilitySet(value value.PokemonAbilityIdSet) {
	abilityId1, abilityId2, hiddenAbilityId := value.GetAbilityIds()
	b.AbilityId1 = convertIdToNullInt16(abilityId1)
	b.AbilityId2 = convertIdToNullInt16(abilityId2)
	b.HiddenAbilityId = convertIdToNullInt16(hiddenAbilityId)
}
func (b *PokemonSchemaBuilder) SetBaseStatsH(value value.PokemonBaseStats) {
	b.BaseStatsH = value.Value()
}
func (b *PokemonSchemaBuilder) SetBaseStatsA(value value.PokemonBaseStats) {
	b.BaseStatsB = value.Value()
}
func (b *PokemonSchemaBuilder) SetBaseStatsB(value value.PokemonBaseStats) {
	b.BaseStatsC = value.Value()
}
func (b *PokemonSchemaBuilder) SetBaseStatsC(value value.PokemonBaseStats) {
	b.BaseStatsD = value.Value()
}
func (b *PokemonSchemaBuilder) SetBaseStatsD(value value.PokemonBaseStats) {
	b.BaseStatsD = value.Value()
}
func (b *PokemonSchemaBuilder) SetBaseStatsS(value value.PokemonBaseStats) {
	b.BaseStatsS = value.Value()
}
func (b *PokemonSchemaBuilder) SetIsFinalEvolution(value bool) {
	b.IsFinalEvolution = value
}

func (b PokemonSchemaBuilder) Build() schema.Pokemon {
	return b.Pokemon
}
