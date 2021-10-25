package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"gorm.io/gorm"
)

type TrainedPokemon struct {
	gorm.Model
	Nickname                 *string
	Gender                   *enum.Gender `sql:"type:gender"`
	AbilityId                uint
	Ability                  Ability
	HeldItemId               uint
	HeldItem                 HeldItem
	MoveId1                  uint
	Move1                    Move `gorm:"foreignKey:MoveId1"`
	MoveId2                  uint
	Move2                    Move `gorm:"foreignKey:MoveId2"`
	MoveId3                  uint
	Move3                    Move `gorm:"foreignKey:MoveId3"`
	MoveId4                  uint
	Move4                    Move `gorm:"foreignKey:MoveId4"`
	UserId                   uint
	User                     User
	AdjustmentId             uint `gorm:"foreignKey:AdjustmentId"`
	TrainedPokemonAdjustment TrainedPokemonAdjustment
}

func (TrainedPokemon) TableName() string {
	return "trained_pokemon"
}
