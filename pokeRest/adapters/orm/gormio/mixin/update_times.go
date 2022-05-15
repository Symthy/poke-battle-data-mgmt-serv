package mixin

import (
	"time"

	"gorm.io/gorm"
)

type UpdateTimes struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
