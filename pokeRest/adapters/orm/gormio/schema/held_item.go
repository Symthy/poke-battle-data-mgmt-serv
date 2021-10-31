package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"

type HeldItem struct {
	ID                uint `gorm:"primaryKey; autoIncrement"`
	Name              string
	Description       string
	CorrectionValueId uint
	CorrectionValue   mixin.CorrectionValue `gorm:"embedded"`

	// relation
	TrainedPokemonBase []TrainedPokemonBase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemonBase
}

func (HeldItem) TableName() string {
	return "held_items"
}
