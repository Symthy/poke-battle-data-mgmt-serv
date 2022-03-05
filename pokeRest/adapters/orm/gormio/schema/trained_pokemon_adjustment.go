package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"gorm.io/gorm"
)

type TrainedPokemonAdjustment struct {
	gorm.Model
	PokemonId    uint
	Pokemon      Pokemon // belongs to
	Nature       enum.Nature
	AbilityId    *int // M:1 <- Ability
	HeldItemId   *int // M:1 <- HeldItem
	EffortValueH int  `gorm:"default:0"`
	EffortValueA int  `gorm:"default:0"`
	EffortValueB int  `gorm:"default:0"`
	EffortValueC int  `gorm:"default:0"`
	EffortValueD int  `gorm:"default:0"`
	EffortValueS int  `gorm:"default:0"`
	MoveId1      *int // M:1 <- Move
	MoveId2      *int // M:1 <- Move
	MoveId3      *int // M:1 <- Move
	MoveId4      *int // M:1 <- Move
}

func (TrainedPokemonAdjustment) TableName() string {
	return "trained_pokemon_adjustments"
}
