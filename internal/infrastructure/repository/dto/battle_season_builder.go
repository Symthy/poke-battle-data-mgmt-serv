package dto

import "github.com/Symthy/PokeRest/internal/domain/entity/battles"

var _ battles.IBattleSeasonNotification = (*BattleSeasonBuilder)(nil)

type BattleSeasonBuilder struct {
	generation uint16
	series     uint16
	season     uint16
}

func NewBattleSeasonBuilder() *BattleSeasonBuilder {
	return &BattleSeasonBuilder{}
}

func (b *BattleSeasonBuilder) SetGeneration(generation uint16) {
	b.generation = generation
}
func (b *BattleSeasonBuilder) SetSeries(series uint16) {
	b.series = series
}
func (b *BattleSeasonBuilder) SetSeason(season uint16) {
	b.season = season
}

func (b BattleSeasonBuilder) Build() (uint16, uint16, uint16) {
	return b.generation, b.series, b.season
}
