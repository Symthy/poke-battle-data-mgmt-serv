package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
)

type TrainedPokemonAttackTarget struct {
	ID                        uint64 `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId          uint64
	TrainedPokemon            TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MoveId                    uint16
	Move                      Move `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemonId           uint16
	TargetPokemon             Pokemon `gorm:"foreignKey:TargetPokemonId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemonNature       enum.Nature
	TargetPokemonAbilityId    uint16
	Ability                   Ability `gorm:"foreignKey:TargetPokemonAbilityId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemonHeldItemId   uint16
	HeldItem                  HeldItem `gorm:"foreignKey:TargetPokemonHeldItemId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemonEffortValueH uint8    `gorm:"default:0"`
	TargetPokemonEffortValueB uint8    `gorm:"default:0"`
	TargetPokemonEffortValueD uint8    `gorm:"default:0"`
}

func (TrainedPokemonAttackTarget) TableName() string {
	return "trained_pokemon_attack_targets"
}
