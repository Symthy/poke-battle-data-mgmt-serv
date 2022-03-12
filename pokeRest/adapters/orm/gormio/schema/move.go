package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
)

type Move struct {
	ID          uint `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Species     enum.MoveSpecies `sql:"type:species"`
	Power       int
	Accuracy    float32
	PP          int
	IsContained *bool `gorm:"default:false"`
	IsCanGuard  *bool `gorm:"default:true"`

	// relation
	// M:M
	Pokemon []*Pokemon `gorm:"many2many:pokemon_moves;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// 1:M -> TrainedPokemon
	TrainedPokemon []TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Move) TableName() string {
	return "moves"
}
