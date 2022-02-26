package dto

import (
	"time"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
)

var _ battles.IBattleSeasonPeriodNotification = (*BattleSeasonPeriodSchemaBuilder)(nil)

type BattleSeasonPeriodSchemaBuilder struct {
	schema.BattleSeason
}

func (b BattleSeasonPeriodSchemaBuilder) SetGeneration(value int) {
	b.Generation = value
}

func (b BattleSeasonPeriodSchemaBuilder) SetSeries(value int) {
	b.Series = value
}

func (b BattleSeasonPeriodSchemaBuilder) SetSeason(value int) {
	b.Season = value
}

func (b BattleSeasonPeriodSchemaBuilder) SetStartDateTime(value time.Time) {
	b.StartDateTime = value
}

func (b BattleSeasonPeriodSchemaBuilder) SetEndDateTime(value time.Time) {
	b.EndDateTime = value
}
