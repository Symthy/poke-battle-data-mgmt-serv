package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type HeldItem struct {
	ID            uint16 `gorm:"primaryKey;autoIncrement:true"`
	Name          string `gorm:"unique"`
	Description   string
	BattleEffects *mixin.BattleEffects `gorm:"embedded"`

	// relation
	TrainedAdjustment []TrainedPokemonAdjustment `gorm:"foreignKey:HeldItemID;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemon
}

func (HeldItem) TableName() string {
	return "held_items"
}
