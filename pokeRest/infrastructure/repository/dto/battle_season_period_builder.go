package dto

import (
	"time"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
)

var _ battles.IBattleSeasonPeriodNotification = (*BattleSeasonPeriodSchemaBuilder)(nil)

type BattleSeasonPeriodSchemaBuilder struct {
	*schema.BattleSeason
}

func (b *BattleSeasonPeriodSchemaBuilder) SetGeneration(value uint16) {
	b.Generation = value
}

func (b *BattleSeasonPeriodSchemaBuilder) SetSeries(value uint16) {
	b.Series = value
}

func (b *BattleSeasonPeriodSchemaBuilder) SetSeason(value uint16) {
	b.Season = value
}

func (b *BattleSeasonPeriodSchemaBuilder) SetStartDateTime(value time.Time) {
	b.StartDateTime = value
}

func (b *BattleSeasonPeriodSchemaBuilder) SetEndDateTime(value time.Time) {
	b.EndDateTime = value
}

func (b BattleSeasonPeriodSchemaBuilder) Build() *schema.BattleSeason {
	return b.BattleSeason
}
