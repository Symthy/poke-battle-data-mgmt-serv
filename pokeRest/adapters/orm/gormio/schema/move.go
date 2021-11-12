package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type Move struct {
	ID          uint `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Species     enum.MoveSpecies `sql:"type:species"`
	Power       int
	Accuracy    int
	PP          int
	IsContained *bool `gorm:"default:false"`
	IsCanGuard  *bool `gorm:"default:true"`

	// relation
	Pokemon             []*Pokemon           `gorm:"many2many:pokemon_moves;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`           // M:M
	TrainedPokemonBase1 []TrainedPokemonBase `gorm:"foreignKey:MoveId1;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemonBase
	TrainedPokemonBase2 []TrainedPokemonBase `gorm:"foreignKey:MoveId2;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemonBase
	TrainedPokemonBase3 []TrainedPokemonBase `gorm:"foreignKey:MoveId3;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemonBase
	TrainedPokemonBase4 []TrainedPokemonBase `gorm:"foreignKey:MoveId4;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemonBase
}

func (Move) TableName() string {
	return "moves"
}
