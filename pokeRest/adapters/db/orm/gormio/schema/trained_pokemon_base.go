package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/db/orm/gormio/enum"
	"gorm.io/gorm"
)

type TrainedPokemonBase struct {
	gorm.Model
	Nature       enum.Nature
	AbilityId    *uint // M:1 <- Ability
	HeldItemId   *uint // M:1 <- HeldItem
	MoveId1      *uint // M:1 <- Move
	MoveId2      *uint // M:1 <- Move
	MoveId3      *uint // M:1 <- Move
	MoveId4      *uint // M:1 <- Move
	EffortValueH int   `gorm:"default:0"`
	EffortValueA int   `gorm:"default:0"`
	EffortValueB int   `gorm:"default:0"`
	EffortValueC int   `gorm:"default:0"`
	EffortValueD int   `gorm:"default:0"`
	EffortValueS int   `gorm:"default:0"`
	CreateUserId *uint // M:1 from User
}

func (TrainedPokemonBase) TableName() string {
	return "trained_pokemon_base"
}
