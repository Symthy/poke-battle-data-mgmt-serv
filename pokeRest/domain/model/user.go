package model

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type User struct {
	id          uint
	name        value.Name
	displayName *string
	email       *value.Email
	profile     *string
	role        value.Role
}

// Todo: builder
func NewUser(
	id uint,
	name value.Name,
	displayName string,
	email value.Email,
	profile string,
	role value.Role) User {
	return User{
		id:          id,
		name:        name,
		displayName: &displayName,
		email:       &email,
		profile:     &profile,
		role:        role,
	}
}

func (u User) Id() uint {
	return u.id
}

func (u User) Name() value.Name {
	return u.name
}

func (u User) DisplayName() *string {
	return u.displayName
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
