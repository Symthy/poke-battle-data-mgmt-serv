package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type HeldItem struct {
	ID            uint16 `gorm:"primaryKey;autoIncrement:true"`
	Name          string
	Description   string
	BattleEffects *mixin.BattleEffects `gorm:"embedded"`

	// relation
	TrainedPokemon []TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemon
}

func (HeldItem) TableName() string {
	return "held_items"
}
