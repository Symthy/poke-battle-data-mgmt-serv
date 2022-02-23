package battles

import (
	"fmt"
	"time"
)

type Season struct {
	generation int
	series     int
	season     int
}

func NewSeason(generation int, series int, season int) Season {
	return Season{generation: generation, series: series, season: season}
}

func (s Season) ApplySeason(from Season) {
	if s.generation == 0 {
		s.generation = from.generation
	}
	if s.series == 0 {
		s.series = from.series
	}
	if s.season == 0 {
		s.season = from.season
	}
}

type SeasonPeriod struct {
	Season
	startDateTime time.Time
	endDateTime   time.Time
}

func NewSeasonPeriod(generation int, series int, season int, startDateTime time.Time,
	endDateTime time.Time) SeasonPeriod {
	return SeasonPeriod{Season{generation, series, season}, startDateTime, endDateTime}
}

// Todo: temporary. change
func (s SeasonPeriod) Id() uint { return 0 }

var datetimeLayout = "2006/01/02 15:04"

func (b SeasonPeriod) BuildSeasonPeriod() string {
	start := b.startDateTime.Format(datetimeLayout)
	end := b.endDateTime.Format(datetimeLayout)
	return fmt.Sprintf("%s ~ %s", start, end)
}
