package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
)

type Move struct {
	ID          uint16 `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Species     enum.MoveSpecies `sql:"type:species"`
	Type        enum.PokemonType `sql:"type:pokemonType"`
	Power       uint16
	Accuracy    float32
	PP          uint8
	IsContained *bool `gorm:"default:false"`
	CanGuard    *bool `gorm:"default:true"`

	// relation
	// M:M
	Pokemon []*Pokemon `gorm:"many2many:pokemon_moves;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// 1:M -> TrainedPokemon
	TrainedPokemon []TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Move) TableName() string {
	return "moves"
}
