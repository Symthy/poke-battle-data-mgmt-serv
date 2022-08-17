package command

import "github.com/Symthy/PokeRest/internal/domain/value"

type CreateUserCommand struct {
	name     string
	password string
	role     value.Role
}

func NewCreateUserCommand(name string, password string, role value.Role) CreateUserCommand {
	return CreateUserCommand{
		name:     name,
		password: password,
		role:     role,
	}
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
