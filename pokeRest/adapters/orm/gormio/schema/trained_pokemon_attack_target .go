package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
)

type TrainedPokemonAttackTarget struct {
	ID                          uint `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId            uint
	TrainedPokemon              TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	MoveId                      uint           // M:1 <- Move
	OpponentPokemonId           uint
	OpponentPokemonNature       enum.Nature
	OpponentPokemonEffortValueH int
	OpponentPokemonEffortValueB int
	OpponentPokemonEffortValueD int
}

func (TrainedPokemonAttackTarget) TableName() string {
	return "trained_pokemon_attack_targets"
}

func (t TrainedPokemonAttackTarget) ConvertToDomain() pokemons.TrainedPokemonAttackTarget {
	return pokemons.NewTrainedPokemonAttackTarget(t.ID)
}
