package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"gorm.io/gorm"
)

type TrainedPokemon struct {
	gorm.Model
	PokemonId            uint
	Pokemon              Pokemon // belongs to
	Nickname             *string
	TrainedPokemonBaseId uint // has many
	Nature               enum.Nature
	EffortValueH         int  `gorm:"default:0"`
	EffortValueA         int  `gorm:"default:0"`
	EffortValueB         int  `gorm:"default:0"`
	EffortValueC         int  `gorm:"default:0"`
	EffortValueD         int  `gorm:"default:0"`
	EffortValueS         int  `gorm:"default:0"`
	CreateUserId         uint // has many
}

func (TrainedPokemon) TableName() string {
	return "trained_pokemon"
}
