package schema

import (
	"time"
)

type BattleSeason struct {
	BattleSeasonSchema
}

type BattleSeasonSchema struct {
	ID            uint16 `gorm:"primaryKey;autoIncrement:false;"`
	Generation    uint16
	Series        uint16
	Season        uint16
	StartDateTime time.Time
	EndDateTime   time.Time
}

func (BattleSeason) TableName() string {
	return "battle_seasons"
}
