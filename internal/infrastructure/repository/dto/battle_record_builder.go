package dto

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/battles"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

var _ battles.IBattleRecordNotification = (*BattleRecordSchemaBuilder)(nil)

type BattleRecordSchemaBuilder struct {
	*schema.BattleRecord
	*BattleOpponentPartySchemaBuilder
	*BattleSeasonBuilder
}

func NewBattleRecordBuilder() *BattleRecordSchemaBuilder {
	return &BattleRecordSchemaBuilder{
		&schema.BattleRecord{},
		NewBattleOpponentPartySchemaBuilder(),
		NewBattleSeasonBuilder(),
	}
}

func (b *BattleRecordSchemaBuilder) SetId(id identifier.BattleRecordId) {
	b.BattleRecord.ID = id.Value()
}
func (b *BattleRecordSchemaBuilder) SetPartyId(partyId identifier.PartyId) {
	b.BattleRecord.PartyID = partyId.Value()
}
func (b *BattleRecordSchemaBuilder) SetUserId(userId identifier.UserId) {
	b.BattleRecord.UserID = userId.Value()
}
func (b *BattleRecordSchemaBuilder) SetBattleResult(battleResult value.BattleResult) {
	b.BattleRecord.Result = enum.BattleResult(battleResult)
}
func (b *BattleRecordSchemaBuilder) SetBattleOpponentPartyId(id identifier.BattleOpponentPartyId) {
	b.BattleRecord.BattleOpponentPartyId = id.Value()
}

// Todo: refactor
func (b *BattleRecordSchemaBuilder) SetSelfElectionPokemons(ids *battles.ElectionPokemons) {
	b.BattleRecord.SelfElectionPokemonId1 = ids.Get(0)
	b.BattleRecord.SelfElectionPokemonId2 = ids.Get(1)
	b.BattleRecord.SelfElectionPokemonId3 = ids.Get(2)
	if ids.Size() == 4 {
		b.BattleRecord.SelfElectionPokemonId4 = ids.Get(3)
	}
}
func (b *BattleRecordSchemaBuilder) SetSelfTrainedPokemons(ids *battles.ElectionTrainedPokemons) {
	b.BattleRecord.SelfTrainedPokemonId1 = ids.Get(0)
	b.BattleRecord.SelfTrainedPokemonId2 = ids.Get(1)
	b.BattleRecord.SelfTrainedPokemonId3 = ids.Get(2)
	if ids.Size() == 4 {
		b.BattleRecord.SelfTrainedPokemonId4 = ids.Get(3)
	}
}
func (b *BattleRecordSchemaBuilder) SetOpponentElectionPokemons(ids *battles.ElectionPokemons) {
	b.BattleRecord.OpponentElectionPokemonId1 = ids.Get(0)
	b.BattleRecord.OpponentElectionPokemonId2 = ids.Get(1)
	b.BattleRecord.OpponentElectionPokemonId3 = ids.Get(2)
	if ids.Size() == 4 {
		b.OpponentElectionPokemonId4 = ids.Get(3)
	}
}

func (b BattleRecordSchemaBuilder) Build() *schema.BattleRecord {
	return b.BattleRecord
}
