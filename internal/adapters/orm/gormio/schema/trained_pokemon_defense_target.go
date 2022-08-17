package schema

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
)

type TrainedPokemonDefenseTarget struct {
	TrainedPokemonDefenseTargetSchema

	// relations
	TrainedPokemon TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Move           Move           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TargetPokemon  Pokemon        `gorm:"foreignKey:TargetPokemonId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ability        Ability        `gorm:"foreignKey:TargetPokemonAbilityId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	HeldItem       HeldItem       `gorm:"foreignKey:TargetPokemonHeldItemId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TrainedPokemonDefenseTargetSchema struct {
	ID                        uint64 `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId          uint64 // ref: TrainedPokemon
	MoveId                    uint16 // ref: Move
	TargetPokemonId           uint16 // ref: Pokemon
	TargetPokemonNature       enum.Nature
	TargetPokemonHeldItemId   uint16 // ref: HeldItem
	TargetPokemonAbilityId    uint16 // ref: Ability
	TargetPokemonEffortValueA uint8  `gorm:"default:0"`
	TargetPokemonEffortValueC uint8  `gorm:"default:0"`
}

func (TrainedPokemonDefenseTarget) TableName() string {
	return "trained_pokemon_defense_targets"
}
