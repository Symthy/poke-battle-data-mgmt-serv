package factory

import (
	"github.com/Symthy/PokeRest/internal/domain/entity/users"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type UserInput struct {
	id   uint64
	name string
	role string
	// Todo
}

func NewUserInput(id uint64, name string, role string) UserInput {
	return UserInput{id, name, role}
}

func NewUserBuilder() *UserInput {
	return &UserInput{}
}

func (b *UserInput) Id(id uint64) {
	b.id = id
}
func (b *UserInput) Name(name string) {
	b.name = name
}
func (b *UserInput) Role(role string) {
	b.role = role
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
	domain := users.NewUser(*id, *name, nil, nil, value.Role(i.role))
	return domain, nil
}
