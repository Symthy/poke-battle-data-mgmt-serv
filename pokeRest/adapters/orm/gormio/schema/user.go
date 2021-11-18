package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
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
	TrainedPokemon     []TrainedPokemon     `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
	TrainedPokemonBase []TrainedPokemonBase `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
	Party              []Party              `gorm:"foreignKey:CreateUserId;references:id"` // 1:M
}

func (User) TableName() string {
	return "users"
}

// Todo: error
func (u User) ConvertToDomain() model.User {
	name, _ := value.NewName(u.Name)
	email, _ := value.NewEmail(*u.Email)
	role := value.Role(u.Role.String())
	return model.NewUser(
		u.ID,
		*name,
		*u.DisplayName,
		*email,
		*u.Profile,
		role,
	)
}

func (u User) ConvertToDomainNonId() model.User {
	name, _ := value.NewName(u.Name)
	email, _ := value.NewEmail(*u.Email)
	role := value.Role(u.Role.String())
	return model.NewUser(
		0,
		*name,
		*u.DisplayName,
		*email,
		*u.Profile,
		role,
	)
}
