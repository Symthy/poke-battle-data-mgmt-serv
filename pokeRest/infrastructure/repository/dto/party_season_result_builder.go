package dto

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ parties.IPartyBattleResultNotification = (*PartySeasonResultSchemaBuilder)(nil)

type PartySeasonResultSchemaBuilder struct {
	*schema.PartySeasonResult
}

func NewPartySeasonResultSchemaBuilder() *PartySeasonResultSchemaBuilder {
	return &PartySeasonResultSchemaBuilder{&schema.PartySeasonResult{}}
}

func (b *PartySeasonResultSchemaBuilder) SetId(id identifier.PartyBattleResultId) {
	b.ID = id.Value()
}
func (b *PartySeasonResultSchemaBuilder) SetMaxRate(maxRate uint16) {
	b.MaxRate = maxRate
}
func (b *PartySeasonResultSchemaBuilder) SetSeasonRanking(ranking uint16) {
	b.SeasonRanking = ranking
}
func (b *PartySeasonResultSchemaBuilder) SetMaxSeasonRanking(maxRanking uint16) {
	b.MaxSeasonRanking = maxRanking
}
func (b *PartySeasonResultSchemaBuilder) SetWinCount(winCount uint16) {
	b.WinCount = winCount
}
func (b *PartySeasonResultSchemaBuilder) SetLoseCount(loseCount uint16) {
	b.LoseCount = loseCount
}

func (b *PartySeasonResultSchemaBuilder) SetGeneration(generation uint16) {
	b.Generation = generation
}
func (b *PartySeasonResultSchemaBuilder) SetSeries(series uint16) {
	b.Series = series
}
func (b *PartySeasonResultSchemaBuilder) SetSeason(season uint16) {
	b.Season = season
}

func (b PartySeasonResultSchemaBuilder) Build() *schema.PartySeasonResult {
	return b.PartySeasonResult
}
