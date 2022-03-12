package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
)

type TrainedPokemonDefenceTarget struct {
	ID                        uint `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId          uint
	TrainedPokemon            TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MoveId                    int
	Move                      Move `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemonId           int
	TargetPokemon             Pokemon `gorm:"foreignKey:TargetPokemonId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemonNature       enum.Nature
	TargetPokemonAbilityId    int
	Ability                   Ability `gorm:"foreignKey:TargetPokemonAbilityId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemonEffortValueA int     `gorm:"default:0"`
	TargetPokemonEffortValueC int     `gorm:"default:0"`
}

func (TrainedPokemonDefenceTarget) TableName() string {
	return "trained_pokemon_defence_targets"
}
