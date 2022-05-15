package schema

import (
	"time"
)

type BattleSeason struct {
	Generation    uint16 `gorm:"primaryKey;autoIncrement:false"`
	Series        uint16 `gorm:"primaryKey;autoIncrement:false"`
	Season        uint16 `gorm:"primaryKey;autoIncrement:false"`
	StartDateTime time.Time
	EndDateTime   time.Time
}

func (BattleSeason) TableName() string {
	return "battle_seasons"
}
