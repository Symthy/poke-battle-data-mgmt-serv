package battles

import (
	"fmt"
	"time"
)

type SeasonPeriod struct {
	Season
	startDateTime time.Time
	endDateTime   time.Time
}

func NewSeasonPeriod(generation, series, season uint64, startDateTime time.Time,
	endDateTime time.Time) (*SeasonPeriod, error) {
	s, err := NewSeason(generation, series, season)
	if err != nil {
		return nil, err
	}
	return &SeasonPeriod{*s, startDateTime, endDateTime}, nil
}

// Todo: temporary. change
func (s SeasonPeriod) Id() uint { return 0 }

var datetimeLayout = "2006/01/02 15:04"

func (b SeasonPeriod) BuildPeriodString() string {
	start := b.startDateTime.Format(datetimeLayout)
	end := b.endDateTime.Format(datetimeLayout)
	return fmt.Sprintf("%s ~ %s", start, end)
}

func (s SeasonPeriod) Notify(note IBattleSeasonPeriodNotification) {
	s.Season.Notify(note)
	note.SetStartDateTime(s.startDateTime)
	note.SetEndDateTime(s.endDateTime)
}
