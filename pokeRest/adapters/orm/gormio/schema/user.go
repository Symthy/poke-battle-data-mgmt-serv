package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type User struct {
	UserSchema
	mixin.UpdateTimes

	// relation
	TrainedPokemon []TrainedPokemon `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
	Party          []Party          `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
}

type UserSchema struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	TwitterID   string `gorm:"unique"`
	DisplayName *string
	Email       *string
	Profile     *string
	Role        enum.Role
}

func (User) TableName() string {
	return "users"
}
