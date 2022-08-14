package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
)

type TrainedPokemonAttackTarget struct {
	TrainedPokemonAttackTargetSchema

	// relations
	TrainedPokemon TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Move           Move           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemon  Pokemon        `gorm:"foreignKey:TargetPokemonId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ability        Ability        `gorm:"foreignKey:TargetPokemonAbilityId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	HeldItem       HeldItem       `gorm:"foreignKey:TargetPokemonHeldItemId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TrainedPokemonAttackTargetSchema struct {
	ID                        uint64 `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId          uint64 // ref: TrainedPokemon
	MoveId                    uint16 // ref: Move
	TargetPokemonId           uint16 // ref: Pokemon
	TargetPokemonNature       enum.Nature
	TargetPokemonAbilityId    uint16 // ref: Ability
	TargetPokemonHeldItemId   uint16 // ref: HeldItem
	TargetPokemonEffortValueH uint8  `gorm:"default:0"`
	TargetPokemonEffortValueB uint8  `gorm:"default:0"`
	TargetPokemonEffortValueD uint8  `gorm:"default:0"`
}

func (TrainedPokemonAttackTarget) TableName() string {
	return "trained_pokemon_attack_targets"
}
