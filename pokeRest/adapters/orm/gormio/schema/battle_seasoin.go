package schema

import (
	"time"

	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
)

type BattleSeason struct {
	Generation    int `gorm:"primaryKey;autoIncrement:false"`
	Series        int `gorm:"primaryKey;autoIncrement:false"`
	Season        int `gorm:"primaryKey;autoIncrement:false"`
	StartDateTime time.Time
	EndDateTime   time.Time
}

func (BattleSeason) TableName() string {
	return "battle_seasons"
}

func (b BattleSeason) ConvertToDomain() battles.SeasonPeriod {
	return battles.NewSeasonPeriod(
		b.Generation,
		b.Series,
		b.Season,
		b.StartDateTime,
		b.EndDateTime,
	)
}

func (b BattleSeason) ConvertToDomainSeason() battles.Season {
	return battles.NewSeason(
		b.Generation,
		b.Series,
		b.Season,
	)
}
