package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type TrainedPokemonAdjustment struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	PokemonId    uint16
	Pokemon      Pokemon // belongs to
	Nature       enum.Nature
	AbilityId    *uint16 // M:1 <- Ability
	HeldItemId   *uint16 // M:1 <- HeldItem
	EffortValueH uint8   `gorm:"default:0"`
	EffortValueA uint8   `gorm:"default:0"`
	EffortValueB uint8   `gorm:"default:0"`
	EffortValueC uint8   `gorm:"default:0"`
	EffortValueD uint8   `gorm:"default:0"`
	EffortValueS uint8   `gorm:"default:0"`
	MoveId1      *uint16 // M:1 <- Move
	MoveId2      *uint16 // M:1 <- Move
	MoveId3      *uint16 // M:1 <- Move
	MoveId4      *uint16 // M:1 <- Move
	mixin.UpdateTimes
}

func (TrainedPokemonAdjustment) TableName() string {
	return "trained_pokemon_adjustments"
}
