package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type TypeCompatibility struct {
	AttackType    enum.PokemonType `gorm:"primaryKey"`
	DefenceType   enum.PokemonType
	Compatibility int `gorm:"default:1"`
}

func (TypeCompatibility) TableName() string {
	return "type_compatibility"
}
