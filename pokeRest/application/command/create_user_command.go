package command

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type CreateUserCommand struct {
	id       uint
	name     string
	password string
	role     value.Role
}

func NewCreateUserCommand(id uint, name string, password string, role value.Role) CreateUserCommand {
	return CreateUserCommand{
		id:       id,
		name:     name,
		password: password,
		role:     role,
	}
}

func (c CreateUserCommand) Id() uint {
	return c.id
}

func (c CreateUserCommand) Name() string {
	return c.name
}

func (c CreateUserCommand) Password() string {
	return c.password
}

func (c CreateUserCommand) Role() value.Role {
	return c.role
}
