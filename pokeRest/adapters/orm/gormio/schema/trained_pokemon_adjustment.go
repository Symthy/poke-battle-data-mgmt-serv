package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type TrainedPokemonAdjustment struct {
	TrainedPokemonAdjustmentSchema
	mixin.UpdateTimes
}

type TrainedPokemonAdjustmentSchema struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	PokemonID    uint16
	Pokemon      Pokemon // belongs to
	Nature       enum.Nature
	AbilityID    *uint16 // M:1 <- Ability
	HeldItemID   *uint16 // M:1 <- HeldItem
	EffortValueH uint8   `gorm:"default:0"`
	EffortValueA uint8   `gorm:"default:0"`
	EffortValueB uint8   `gorm:"default:0"`
	EffortValueC uint8   `gorm:"default:0"`
	EffortValueD uint8   `gorm:"default:0"`
	EffortValueS uint8   `gorm:"default:0"`
	MoveID1      *uint16 // M:1 <- Move
	MoveID2      *uint16 // M:1 <- Move
	MoveID3      *uint16 // M:1 <- Move
	MoveID4      *uint16 // M:1 <- Move
}

func (TrainedPokemonAdjustment) TableName() string {
	return "trained_pokemon_adjustments"
}
