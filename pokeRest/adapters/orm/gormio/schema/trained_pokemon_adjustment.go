package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"gorm.io/gorm"
)

type TrainedPokemonAdjustment struct {
	gorm.Model
	PokemonId    uint
	Pokemon      Pokemon
	CreateUserId *uint // 意図的に外部キーにしない
	Nature       enum.Nature
	EffortValueH int `gorm:"default:0"`
	EffortValueA int `gorm:"default:0"`
	EffortValueB int `gorm:"default:0"`
	EffortValueC int `gorm:"default:0"`
	EffortValueD int `gorm:"default:0"`
	EffortValueS int `gorm:"default:0"`
}

func (TrainedPokemonAdjustment) TableName() string {
	return "trained_pokemon_adjustment"
}
