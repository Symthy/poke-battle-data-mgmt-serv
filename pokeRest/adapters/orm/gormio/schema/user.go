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
}

func (User) TableName() string {
	return "user"
}
