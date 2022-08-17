package dto

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/battles"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

var _ battles.IBattleOpponentPartyNotification = (*BattleOpponentPartySchemaBuilder)(nil)

type BattleOpponentPartySchemaBuilder struct {
	*schema.BattleOpponentParty
}

func NewBattleOpponentPartySchemaBuilder() *BattleOpponentPartySchemaBuilder {
	return &BattleOpponentPartySchemaBuilder{
		BattleOpponentParty: &schema.BattleOpponentParty{},
	}
}

func (b BattleOpponentPartySchemaBuilder) SetOpponentPartyId(id identifier.BattleOpponentPartyId) {
	b.BattleOpponentParty.ID = id.Value()
}

func (b BattleOpponentPartySchemaBuilder) SetPokemonIds(ids *value.PartyPokemonIds) {
	b.BattleOpponentParty.OpponentPokemonId1 = ids.Get(0)
	b.BattleOpponentParty.OpponentPokemonId2 = ids.Get(1)
	b.BattleOpponentParty.OpponentPokemonId3 = ids.Get(2)
	b.BattleOpponentParty.OpponentPokemonId4 = ids.Get(3)
	b.BattleOpponentParty.OpponentPokemonId5 = ids.Get(4)
	b.BattleOpponentParty.OpponentPokemonId6 = ids.Get(5)
}

func (b BattleOpponentPartySchemaBuilder) Build() *schema.BattleOpponentParty {
	return b.BattleOpponentParty
}
