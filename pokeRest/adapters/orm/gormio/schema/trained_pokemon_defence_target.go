package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
)

type TrainedPokemonDefenceTarget struct {
	ID                        uint `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId          uint
	TrainedPokemon            TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	MoveId                    int            // M:1 <- Move
	TargetPokemonId           int
	TargetPokemonNature       enum.Nature
	TargetPokemonEffortValueA int `gorm:"default:0"`
	TargetPokemonEffortValueC int `gorm:"default:0"`
}

func (TrainedPokemonDefenceTarget) TableName() string {
	return "trained_pokemon_defence_targets"
}