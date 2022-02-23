package dto

import (
	"time"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
)

var _ battles.IBattleSeasonNotification = (*BattleSeasonSchemaBuilder)(nil)

type BattleSeasonSchemaBuilder struct {
	schema.BattleSeason
}

func (b BattleSeasonSchemaBuilder) Generation(value int) {
	b.BattleSeason.Generation = value
}

func (b BattleSeasonSchemaBuilder) Series(value int) {
	b.BattleSeason.Series = value
}

func (b BattleSeasonSchemaBuilder) Season(value int) {
	b.BattleSeason.Season = value
}

func (b BattleSeasonSchemaBuilder) StartDateTime(value time.Time) {
	b.BattleSeason.StartDateTime = value
}

func (b BattleSeasonSchemaBuilder) EndDateTime(value time.Time) {
	b.BattleSeason.EndDateTime = value
}
