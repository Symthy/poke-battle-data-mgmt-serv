package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
)

type TrainedPokemonAttackTarget struct {
	ID                        uint `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId          uint
	TrainedPokemon            TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	MoveId                    int            // M:1 <- Move
	TargetPokemonId           int
	TargetPokemonNature       enum.Nature
	TargetPokemonEffortValueH int `gorm:"default:0"`
	TargetPokemonEffortValueB int `gorm:"default:0"`
	TargetPokemonEffortValueD int `gorm:"default:0"`
}

func (TrainedPokemonAttackTarget) TableName() string {
	return "trained_pokemon_attack_targets"
}

func (t TrainedPokemonAttackTarget) ConvertToDomain() trainings.TrainedPokemonAttackTarget {
	return trainings.NewTrainedPokemonAttackTarget(
		t.ID, t.TrainedPokemonId, t.MoveId, t.TargetPokemonId, t.TargetPokemonNature.String(),
		t.TargetPokemonEffortValueD, t.TargetPokemonEffortValueB, t.TargetPokemonEffortValueD)
}
