package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
)

type HeldItem struct {
	ID                uint `gorm:"primaryKey;autoIncrement:true"`
	Name              string
	Description       string
	CorrectionValueId uint
	CorrectionValue   mixin.CorrectionValue `gorm:"embedded"`

	// relation
	TrainedPokemon []TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemon
}

func (HeldItem) TableName() string {
	return "held_items"
}

func (i HeldItem) ConvertToDomain() items.HeldItem {
	return items.NewHeldItem(i.ID)
}
