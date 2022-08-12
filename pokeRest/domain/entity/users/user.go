package users

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service/users/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"golang.org/x/crypto/bcrypt"
)

var _ entity.IDomain[identifier.UserId, uint64] = (*User)(nil)

type User struct {
	id          identifier.UserId
	name        value.UserName
	displayName *string
	password    []byte
	email       *value.Email
	profile     *string
	role        value.Role
	// Todo: have twitter info only
}

// Todo: factory
func NewUser(
	id identifier.UserId,
	name value.UserName,
	displayName *string,
	email *value.Email,
	profile *string,
	role value.Role) *User {
	return &User{
		id:          id,
		name:        name,
		displayName: displayName,
		email:       email,
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
		nil,
		command.Role(),
	)
}

func (u User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password(), []byte(password))
}

// Todo: refactor Notification
func (u User) Id() identifier.UserId {
	return u.id
}

func (u User) Name() value.UserName {
	return u.name
}

func (u User) DisplayName() *string {
	return u.displayName
}

func (u User) Password() []byte {
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

func (u *User) MaskPassword() {
	u.password = nil
}
