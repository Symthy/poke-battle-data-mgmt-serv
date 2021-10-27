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
	TrainedPokemon     []*TrainedPokemon     `gorm:"foreignKey:id;references:CreateUserId"` // has many
	TrainedPokemonBase []*TrainedPokemonBase `gorm:"foreignKey:id;references:CreateUserId"` // has many
	Party              []*Party              `gorm:"foreignKey:id;references:CreateUserId"` // has many
}

func (User) TableName() string {
	return "user"
}
