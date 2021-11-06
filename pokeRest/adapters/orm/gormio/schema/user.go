package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        *string
	DisplayName *string
	Email       *string
	Profile     *string
	Role        enum.Role

	// relation
	TrainedPokemon     []TrainedPokemon     `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
	TrainedPokemonBase []TrainedPokemonBase `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
	Party              []Party              `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
}

func (User) TableName() string {
	return "users"
}
