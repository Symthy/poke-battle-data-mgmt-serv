package factory

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"

type SeasonInput struct {
	generation uint64
	series     uint64
	season     uint64
}

func NewSeasonBuilder() *SeasonInput {
	return &SeasonInput{}
}

func (b *SeasonInput) Generation(generation uint64) *SeasonInput {
	b.generation = generation
	return b
}

func (b *SeasonInput) Series(series uint64) *SeasonInput {
	b.series = series
	return b
}

func (b *SeasonInput) Season(season uint64) *SeasonInput {
	b.season = season
	return b
}

func (b SeasonInput) BuildDomain() (*battles.Season, error) {
	return battles.NewSeason(b.generation, b.series, b.season)
}
