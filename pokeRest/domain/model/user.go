package model

import (
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id          uint
	name        value.UserName
	displayName *string
	password    *[]byte
	email       *value.Email
	profile     *string
	role        value.Role
}

// Todo: builder
func NewUser(
	id uint,
	name value.UserName,
	displayName *string,
	email *value.Email,
	profile *string,
	role value.Role) User {
	return User{
		id:          id,
		name:        name,
		displayName: displayName,
		email:       email,
		profile:     profile,
		role:        role,
	}
}

func NewUserFromCommand(command command.CreateUserCommand) User {
	name, _ := value.NewUserName(command.Name())
	return NewUser(
		0,
		*name,
		nil,
		nil,
		nil,
		command.Role(),
	)
}

func (u User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword(*u.Password(), []byte(password))
}

func (u User) Id() uint {
	return u.id
}

func (u User) Name() value.UserName {
	return u.name
}

func (u User) DisplayName() *string {
	return u.displayName
}

func (u User) Password() *[]byte {
	return u.password
}

func (u User) Email() *value.Email {
	return u.email
}

func (u User) Profile() *string {
	return u.profile
}

func (u User) Role() value.Role {
	return u.role
}

func (u *User) ResetPassword() {
	u.password = nil
}
