package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type Pokemon struct {
	ID               uint `gorm:"primaryKey; autoIncrement"`
	PokedexNo        int
	FormId           uint
	FormName         string
	Form             Form
	Name             string
	EnglishName      string
	Type1            enum.PokemonType
	Type2            enum.PokemonType
	AbilityId1       uint
	Ability1         Ability `gorm:"foreignKey:AbilityId1"`
	AbilityId2       uint
	Ability2         Ability `gorm:"foreignKey:AbilityId2"`
	HiddenAbilityId  uint
	Ability3         Ability `gorm:"foreignKey:HiddenAbilityId"`
	IsFinalEvolution *bool   `gorm:"default:false"`
}

func (Pokemon) TableName() string {
	return "pokemon"
}
