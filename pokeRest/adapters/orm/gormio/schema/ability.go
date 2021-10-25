package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"

type Ability struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	Description     string
	CorrectionValue mixin.CorrectionValue `gorm:"embedded"`
}

func (Ability) TableName() string {
	return "abilities"
}
