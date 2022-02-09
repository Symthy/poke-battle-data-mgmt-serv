package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
)

type TrainedPokemonDeffenceTarget struct {
	ID                          uint `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId            uint
	TrainedPokemon              TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	MoveId                      uint           // M:1 <- Move
	OpponentPokemonId           uint
	OpponentPokemonNature       enum.Nature
	OpponentPokemonEffortValueA int
	OpponentPokemonEffortValueC int
}

func (TrainedPokemonDeffenceTarget) TableName() string {
	return "trained_pokemon_deffence_targets"
}

func (t TrainedPokemonDeffenceTarget) ConvertToDomain() pokemons.TrainedPokemonDeffenceTarget {
	return pokemons.NewTrainedPokemonDeffenceTarget(t.ID)
}
