package mock

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/test/data"
)

// implements IUserRepository
type UserRepositoryMock struct {
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (mock UserRepositoryMock) FindById(id uint) (model.User, error) {
	dummyUser := data.DummyUser1()
	if dummyUser.ID != id {
		return model.User{}, nil
	}
	return dummyUser.ConvertToDomain(), nil
}
