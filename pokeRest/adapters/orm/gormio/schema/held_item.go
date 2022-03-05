package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type HeldItem struct {
	ID              uint `gorm:"primaryKey;autoIncrement:true"`
	Name            string
	Description     string
	CorrectionValue mixin.CorrectionValue `gorm:"embedded"`

	// relation
	TrainedPokemon []TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemon
}

func (HeldItem) TableName() string {
	return "held_items"
}
