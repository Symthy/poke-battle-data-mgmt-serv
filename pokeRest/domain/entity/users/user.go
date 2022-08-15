package users

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service/users/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.UserId, uint64] = (*User)(nil)

type User struct {
	id          identifier.UserId
	name        value.UserName
	displayName *string
	// password    []byte
	profile *string
	role    value.Role
	// Todo: have twitter info only
}

// Todo: factory
func NewUser(
	id identifier.UserId,
	name value.UserName,
	displayName *string,
	profile *string,
	role value.Role) *User {
	return &User{
		id:          id,
		name:        name,
		displayName: displayName,
		profile:     profile,
		role:        role,
	}
}

func NewUserFromCommand(command command.CreateUserCommand) *User {
	name, _ := value.NewUserName(command.Name())
	return NewUser(
		identifier.NewEmptyUserId(),
		*name,
		nil,
		nil,
		command.Role(),
	)
}

// func (u User) ValidatePassword(password string) error {
// 	return bcrypt.CompareHashAndPassword(u.Password(), []byte(password))
// }

func (u User) Id() identifier.UserId {
	return u.id
}

func (u User) Name() value.UserName {
	return u.name
}

func (u User) DisplayName() *string {
	return u.displayName
}

func (u User) Profile() *string {
	return u.profile
}

func (u User) Role() value.Role {
	return u.role
}

func (u User) Notify(note IUserNotification) {
	note.SetId(u.id)
	note.SetName(u.name)
	note.SetDisplayName(u.displayName)
	note.SetProfile(u.profile)
	note.SetRole(u.role)
}
