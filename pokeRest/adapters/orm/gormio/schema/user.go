package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type User struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	DisplayName *string
	Email       *string
	Profile     *string
	Role        enum.Role
	mixin.UpdateTimes

	// relation
	TrainedPokemon []TrainedPokemon `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
	Party          []Party          `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
}

func (User) TableName() string {
	return "users"
}
