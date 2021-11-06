package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/db/orm/gormio/enum"

type TypeCompatibility struct {
	AttackType    enum.PokemonType `gorm:"primaryKey"`
	DefenceType   enum.PokemonType `gorm:"primaryKey"`
	Compatibility int              `gorm:"default:1"`
}

func (TypeCompatibility) TableName() string {
	return "type_compatibility"
}
