package battles

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/errs"
)

type Season struct {
	generation int
	series     int
	season     int
}

func NewSeason(generation int, series int, season int) (*Season, error) {
	isNotSatisfy := (generation == 0 && (series != 0 || season != 0)) ||
		(generation != 0 && (series == 0 || season == 0))

	if isNotSatisfy {
		return nil, errs.ThrowErrorInvalidValue("Season", "generation,series,season",
			fmt.Sprintf("%d,%d,%d", generation, series, season))
	}
	return &Season{generation: generation, series: series, season: season}, nil
}

func (s Season) UnsolvedSeason() bool {
	return s.generation == 0 && s.series == 0 && s.season == 0
}

func (s *Season) ApplySeason(from Season) {
	s.generation = from.generation
	s.series = from.series
	s.season = from.season
}

func (s Season) Notify(note IBattleSeasonNotification) {
	note.SetGeneration(s.generation)
	note.SetSeries(s.series)
	note.SetSeason(s.season)
}
