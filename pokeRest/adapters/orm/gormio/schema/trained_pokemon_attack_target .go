package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
)

type TrainedPokemonAttackTarget struct {
	ID                          uint `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId            uint
	TrainedPokemon              TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	MoveId                      uint           // M:1 <- Move
	OpponentPokemonId           uint
	OpponentPokemonNature       enum.Nature
	OpponentPokemonEffortValueH int `gorm:"default:0"`
	OpponentPokemonEffortValueB int `gorm:"default:0"`
	OpponentPokemonEffortValueD int `gorm:"default:0"`
}

func (TrainedPokemonAttackTarget) TableName() string {
	return "trained_pokemon_attack_targets"
}

func (t TrainedPokemonAttackTarget) ConvertToDomain() trainings.TrainedPokemonAttackTarget {
	return trainings.NewTrainedPokemonAttackTarget(t.ID)
}
