package schema

import (
	"time"
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
