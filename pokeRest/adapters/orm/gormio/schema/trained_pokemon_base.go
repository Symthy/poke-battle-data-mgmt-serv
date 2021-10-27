package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"gorm.io/gorm"
)

type TrainedPokemonBase struct {
	gorm.Model
	Gender     *enum.Gender `sql:"type:gender"`
	AbilityId  *uint        // has one
	HeldItemId *uint        // has one
	MoveId1    *uint        // has one
	MoveId2    *uint        // has one
	MoveId3    *uint        // has one
	MoveId4    *uint        // has one
	UserId     *uint        // has one
	// relation
	TrainedPokemon []TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"` // has many
}

func (TrainedPokemonBase) TableName() string {
	return "trained_pokemon_base"
}
