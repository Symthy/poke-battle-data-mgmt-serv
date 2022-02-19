package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
)

type TrainedPokemonDeffenceTarget struct {
	ID                        uint `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId          uint
	TrainedPokemon            TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	MoveId                    int            // M:1 <- Move
	TargetPokemonId           int
	TargetPokemonNature       enum.Nature
	TargetPokemonEffortValueA int `gorm:"default:0"`
	TargetPokemonEffortValueC int `gorm:"default:0"`
}

func (TrainedPokemonDeffenceTarget) TableName() string {
	return "trained_pokemon_deffence_targets"
}

func (t TrainedPokemonDeffenceTarget) ConvertToDomain() trainings.TrainedPokemonDefenceTarget {
	return trainings.NewTrainedPokemonDefenceTarget(
		t.ID, t.TrainedPokemonId, t.MoveId, t.TargetPokemonId, t.TargetPokemonNature.String(),
		t.TargetPokemonEffortValueA, t.TargetPokemonEffortValueC)
}
