package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"

type HeldItem struct {
	ID                uint `gorm:"primaryKey; autoIncrement"`
	Name              string
	Description       string
	CorrectionValueId uint
	CorrectionValue   mixin.CorrectionValue `gorm:"embedded"`
}

func (HeldItem) TableName() string {
	return "held_item"
}
