package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type Move struct {
	ID          uint `gorm:"primaryKey; autoIncrement"`
	Name        string
	Species     enum.MoveSpecies `sql:"type:species"`
	Power       int
	Accuracy    int
	PP          int
	IsContained *bool `gorm:"default:false"`
	IsCanGuard  *bool `gorm:"default:true"`

	// relation
	Pokemon         []Pokemon      `gorm:"many2many:pokemon_moves;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Many To Many
	TrainedPokemon1 TrainedPokemon `gorm:"foreignKey:MoveId1;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`     // has one
	TrainedPokemon2 TrainedPokemon `gorm:"foreignKey:MoveId2;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`     // has one
	TrainedPokemon3 TrainedPokemon `gorm:"foreignKey:MoveId3;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`     // has one
	TrainedPokemon4 TrainedPokemon `gorm:"foreignKey:MoveId4;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`     // has one
}

func (Move) TableName() string {
	return "move"
}
