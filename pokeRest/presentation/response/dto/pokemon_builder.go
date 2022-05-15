package dto

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ pokemons.IPokemonNotification = (*PokemonResponseBuilder)(nil)

// Todo: auto gen
type PokemonResponseBuilder struct {
	server.Pokemon
}

func (b *PokemonResponseBuilder) SetId(id identifier.PokemonId) {
	b.Id = id.Value()
}
func (b *PokemonResponseBuilder) SetPokedexId(value value.PokedexId) {
	b.PokedexNo = value.PokedexNo()
	b.FormNo = uint16(value.FormNo())
}
func (b *PokemonResponseBuilder) SetFormName(value string) {
	b.FormName = value
}
func (b *PokemonResponseBuilder) SetName(value string) {
	b.Name = value
}
func (b *PokemonResponseBuilder) SetEnglishName(value string) {
	b.EnglishName = value
}
func (b *PokemonResponseBuilder) SetGeneration(value uint16) {
}
func (b *PokemonResponseBuilder) SetTypeSet(value value.PokemonTypeSet) {
	type1, type2 := value.GetTypeEnglishNames()
	b.Type1 = type1
	b.Type2 = &type2
}
func (b *PokemonResponseBuilder) SetAbilitySet(value value.PokemonAbilityIdSet) {
	abilityId1, abilityId2, hiddenAbilityId := value.GetAbilityIds()
	b.Ability1 = abilityId1.Value()
	ability2 := abilityId2.Value()
	b.Ability2 = &ability2
	hiddenAbility := hiddenAbilityId.Value()
	b.HiddenAbility = &hiddenAbility
}
func (b *PokemonResponseBuilder) SetBaseStatsH(value value.PokemonBaseStats) {
	b.BaseStatsH = value.Value()
}
func (b *PokemonResponseBuilder) SetBaseStatsA(value value.PokemonBaseStats) {
	b.BaseStatsB = value.Value()
}
func (b *PokemonResponseBuilder) SetBaseStatsB(value value.PokemonBaseStats) {
	b.BaseStatsC = value.Value()
}
func (b *PokemonResponseBuilder) SetBaseStatsC(value value.PokemonBaseStats) {
	b.BaseStatsD = value.Value()
}
func (b *PokemonResponseBuilder) SetBaseStatsD(value value.PokemonBaseStats) {
	b.BaseStatsD = value.Value()
}
func (b *PokemonResponseBuilder) SetBaseStatsS(value value.PokemonBaseStats) {
	b.BaseStatsS = value.Value()
}
func (b *PokemonResponseBuilder) SetIsFinalEvolution(value bool) {
	// nothing
}
func (b PokemonResponseBuilder) Build() server.Pokemon {
	return b.Pokemon
}
