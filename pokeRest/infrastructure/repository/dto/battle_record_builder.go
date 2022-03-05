package dto

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ battles.IBattleRecordNotification = (*BattleRecordSchemaBuilder)(nil)

type BattleRecordSchemaBuilder struct {
	schema.BattleRecord
	BattleSeasonBuilder
}

func NewBattleRecordBuilder() BattleRecordSchemaBuilder {
	return BattleRecordSchemaBuilder{}
}

func (b *BattleRecordSchemaBuilder) SetId(id identifier.BattleRecordId) {
	b.ID = id.Value()
}
func (b *BattleRecordSchemaBuilder) SetPartyId(partyId identifier.PartyId) {
	b.PartyId = partyId.Value()
}

func (b *BattleRecordSchemaBuilder) SetBattleResult(battleResult value.BattleResult) {
	b.Result = enum.BattleResult(battleResult)
}
func (b *BattleRecordSchemaBuilder) SetBattleOpponentPartyId(id identifier.BattleOpponentPartyId) {
	b.BattleOpponentPartyId = id.Value()
}

// Todo: refactor
func (b *BattleRecordSchemaBuilder) SetSelfElectionPokemons(ids battles.ElectionPokemons) {
	b.SelfElectionPokemonId1 = ids.Get(0)
	b.SelfElectionPokemonId2 = ids.Get(1)
	b.SelfElectionPokemonId3 = ids.Get(2)
	if ids.Size() == 4 {
		b.SelfElectionPokemonId4 = ids.Get(3)
	}
}
func (b *BattleRecordSchemaBuilder) SetSelfTrainedPokemons(ids battles.ElectionPokemons) {
	b.SelfTrainedPokemonId1 = ids.Get(0)
	b.SelfTrainedPokemonId2 = ids.Get(1)
	b.SelfTrainedPokemonId3 = ids.Get(2)
	if ids.Size() == 4 {
		b.SelfTrainedPokemonId4 = ids.Get(3)
	}
}
func (b *BattleRecordSchemaBuilder) SetOpponentElectionPokemons(ids battles.ElectionPokemons) {
	b.OpponentElectionPokemonId1 = ids.Get(0)
	b.OpponentElectionPokemonId2 = ids.Get(1)
	b.OpponentElectionPokemonId3 = ids.Get(2)
	if ids.Size() == 4 {
		b.OpponentElectionPokemonId4 = ids.Get(3)
	}
}

func (b BattleRecordSchemaBuilder) Build() schema.BattleRecord {
	return b.BattleRecord
}

var _ battles.IBattleSeasonNotification = (*BattleSeasonBuilder)(nil)

type BattleSeasonBuilder struct {
	generation int
	series     int
	season     int
}

func (b *BattleSeasonBuilder) SetGeneration(generation int) {
	b.generation = generation
}
func (b *BattleSeasonBuilder) SetSeries(series int) {
	b.series = series
}
func (b *BattleSeasonBuilder) SetSeason(season int) {
	b.season = season
}

func (b BattleSeasonBuilder) Build() (int, int, int) {
	return b.generation, b.series, b.season
}
