package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type TrainedPokemonDeffenceTarget struct {
	ID                          uint `gorm:"primaryKey;autoIncrement:true"`
	TrainedPokemonId            uint
	TrainedPokemon              TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	TrainedPokemonBaseId        uint
	TrainedPokemonBase          TrainedPokemonBase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	MoveId                      uint               // M:1 <- Move
	OpponentPokemonId           uint
	OpponentPokemonNature       enum.Nature
	OpponentPokemonEffortValueA int
	OpponentPokemonEffortValueC int
}
