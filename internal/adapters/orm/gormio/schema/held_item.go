package schema

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/mixin"
)

type HeldItem struct {
	HeldItemSchema
	// relation
	TrainedAdjustment []TrainedPokemonAdjustment `gorm:"foreignKey:HeldItemID;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemon
}

type HeldItemSchema struct {
	ID            uint16 `gorm:"primaryKey;autoIncrement:true"`
	Name          string `gorm:"unique"`
	Description   string
	BattleEffects *mixin.BattleEffects `gorm:"embedded"`
}

func (HeldItem) TableName() string {
	return "held_items"
}
