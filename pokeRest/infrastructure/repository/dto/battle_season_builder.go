package dto

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"

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
