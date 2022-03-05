package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type UserInput struct {
	id   uint
	name string
	role string
	// Todo
}

func NewUserInput(id uint, name string, role string) UserInput {
	return UserInput{id, name, role}
}

func (i UserInput) BuildDomain() (*users.User, error) {
	id, err := identifier.NewUserId(i.id)
	if err != nil {
		return nil, err
	}
	name, err := value.NewUserName(i.name)
	if err != nil {
		return nil, err
	}
	domain := users.NewUser(*id, *name, nil, nil, nil, value.Role(i.role))
	return &domain, nil
}
