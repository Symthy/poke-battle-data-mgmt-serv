package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
)

type TrainedPokemonAttackTarget struct {
	ID                        uint `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId          uint
	TrainedPokemon            TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MoveId                    int
	Move                      Move `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemonId           int
	TargetPokemon             Pokemon `gorm:"foreignKey:TargetPokemonId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemonAbilityId    int
	Ability                   Ability `gorm:"foreignKey:TargetPokemonAbilityId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemonNature       enum.Nature
	TargetPokemonEffortValueH int `gorm:"default:0"`
	TargetPokemonEffortValueB int `gorm:"default:0"`
	TargetPokemonEffortValueD int `gorm:"default:0"`
}

func (TrainedPokemonAttackTarget) TableName() string {
	return "trained_pokemon_attack_targets"
}
