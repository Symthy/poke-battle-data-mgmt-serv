package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/moves"
)

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
	// M:M
	Pokemon []*Pokemon `gorm:"many2many:pokemon_moves;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// 1:M -> TrainedPokemon
	TrainedPokemon []TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// 1:M -> TrainedPokemonAttackTarget
	TrainedPokemonAttackTarget TrainedPokemonAttackTarget `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// 1:M -> TrainedPokemonDeffenceTarget
	TrainedPokemonDeffenceTarget TrainedPokemonDeffenceTarget `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Move) TableName() string {
	return "moves"
}

func (m Move) ConvertToDomain() moves.Move {
	return moves.NewMove(m.ID)
}
