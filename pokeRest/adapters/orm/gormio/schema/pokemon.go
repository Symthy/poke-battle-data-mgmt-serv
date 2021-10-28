package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type Pokemon struct {
	ID               uint `gorm:"primaryKey; autoIncrement"`
	PokedexNo        int
	FormNo           uint
	FormName         string
	Form             Form `gorm:"foreignKey:id,name;references:FormNo,FormName"` // belongs to
	Name             string
	EnglishName      string
	Generation       int
	Type1            enum.PokemonType
	Type2            enum.PokemonType
	AbilityId1       uint  // has one
	AbilityId2       *uint // has one
	HiddenAbilityId  *uint // has one
	IsFinalEvolution *bool `gorm:"default:false"`

	// relation
	Move []Move `gorm:"many2many:pokemon_moves;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Many To Many
}

func (Pokemon) TableName() string {
	return "pokemon"
}
