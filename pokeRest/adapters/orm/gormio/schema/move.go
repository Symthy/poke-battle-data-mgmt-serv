package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
)

type Move struct {
	ID          uint16           `gorm:"primaryKey;autoIncrement:true"`
	Name        string           `gorm:"unique"`
	Species     enum.MoveSpecies `sql:"type:species"`
	Type        enum.PokemonType `sql:"type:pokemonType"`
	Power       uint16
	Accuracy    float32
	PP          uint8
	IsContained *bool `gorm:"default:false"`
	CanGuard    *bool `gorm:"default:true"`

	// relation
	// M:M
	Pokemon []*Pokemon `gorm:"many2many:pokemons_moves;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// 1:M -> TrainedPokemonAdjustment
	TrainedAdjustment1 []TrainedPokemonAdjustment `gorm:"foreignKey:MoveID1;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TrainedAdjustment2 []TrainedPokemonAdjustment `gorm:"foreignKey:MoveID2;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TrainedAdjustment3 []TrainedPokemonAdjustment `gorm:"foreignKey:MoveID3;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TrainedAdjustment4 []TrainedPokemonAdjustment `gorm:"foreignKey:MoveID4;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Move) TableName() string {
	return "moves"
}
