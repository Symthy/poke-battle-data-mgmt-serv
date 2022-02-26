package battles

type Season struct {
	generation int
	series     int
	season     int
}

func NewSeason(generation int, series int, season int) Season {
	return Season{generation: generation, series: series, season: season}
}

func (s Season) isSolvedSeason() bool {
	return s.generation != 0 && s.series != 0 && s.season == 0
}

func (s *Season) ApplySeason(from Season) {
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

func (s Season) Notify(note IBattleSeasonNotification) {
	note.SetGeneration(s.generation)
	note.SetSeries(s.series)
	note.SetSeason(s.season)
}
