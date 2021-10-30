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
	Pokemon             []*Pokemon         `gorm:"many2many:pokemon_moves;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`           // Many To Many
	TrainedPokemonBase1 TrainedPokemonBase `gorm:"foreignKey:MoveId1;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // has one
	TrainedPokemonBase2 TrainedPokemonBase `gorm:"foreignKey:MoveId2;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // has one
	TrainedPokemonBase3 TrainedPokemonBase `gorm:"foreignKey:MoveId3;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // has one
	TrainedPokemonBase4 TrainedPokemonBase `gorm:"foreignKey:MoveId4;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // has one
}

func (Move) TableName() string {
	return "move"
}
