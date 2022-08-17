package battles

import (
	"fmt"

	"github.com/Symthy/PokeRest/internal/errs"
)

type Season struct {
	generation uint16
	series     uint16
	season     uint16
}

func NewSeason(generation, series, season uint64) (*Season, error) {
	// Todo: validate upper limit
	isNotSatisfy := (generation == 0 && (series != 0 || season != 0)) ||
		(generation != 0 && (series == 0 || season == 0))

	if isNotSatisfy {
		return nil, errs.ThrowErrorInvalidValue("Season", "generation,series,season",
			fmt.Sprintf("%d,%d,%d", generation, series, season))
	}
	return &Season{generation: uint16(generation), series: uint16(series), season: uint16(season)}, nil
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
