package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	DisplayName *string
	Email       *string
	Profile     *string
	Role        enum.Role

	// relation
	TrainedPokemon []TrainedPokemon `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
	Party          []Party          `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
}

func (User) TableName() string {
	return "users"
}

// Todo: error
func (u User) ConvertToDomain() users.User {
	name, _ := value.NewUserName(u.Name)

	var email *value.Email = nil
	if u.Email != nil {
		email, _ = value.NewEmail(*u.Email)
	}
	role := value.Role(u.Role.String())
	return users.NewUser(
		u.ID,
		*name,
		u.DisplayName,
		email,
		u.Profile,
		role,
	)
}

// Todo: error log output
func (u User) ConvertToDomainNonId() users.User {
	name, _ := value.NewUserName(u.Name)
	email, _ := value.NewEmail(*u.Email)
	role := value.Role(u.Role.String())
	return users.NewUser(
		0,
		*name,
		u.DisplayName,
		email,
		u.Profile,
		role,
	)
}
