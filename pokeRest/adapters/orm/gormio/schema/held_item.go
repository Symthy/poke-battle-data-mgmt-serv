package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/items"
)

type HeldItem struct {
	ID                uint `gorm:"primaryKey;autoIncrement:true"`
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

func (i HeldItem) ConvertToDomain() items.HeldItem {
	return items.NewHeldItem(i.ID)
}
