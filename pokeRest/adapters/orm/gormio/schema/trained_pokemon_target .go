package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type TrainedPokemonTarget struct {
	ID                         uint `gorm:"primaryKey;autoIncrement"`
	SelfPokemonId              uint
	SelfPokemonNature          enum.Nature
	SelfPokemonEffortValue     int
	MoveId                     uint
	OpponentPokemonId          uint
	OpponentPokemonNature      enum.Nature
	OpponentPokemonEffortValue int
}
