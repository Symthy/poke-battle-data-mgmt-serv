package mock

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model/users"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/test/data"
	"gorm.io/gorm"
)

var _ repository.IUserRepository = (*UserRepositoryMock)(nil)

type UserRepositoryMock struct {
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (mock UserRepositoryMock) FindById(id uint) (*users.User, error) {
	dummyUser := data.DummyUser1()
	if dummyUser.ID != id {
		return &users.User{}, nil
	}
	u := dummyUser.ConvertToDomain()
	return &u, nil
}

func (mock UserRepositoryMock) FindByName(targetName string, filterFields ...string) (*users.User, error) {
	dummyUser := data.DummyUser1(filterFields...)
	if dummyUser.Name != targetName {
		return nil, gorm.ErrRecordNotFound
	}
	u := dummyUser.ConvertToDomain()
	return &u, nil
}

func (mock UserRepositoryMock) Create(user users.User) (*users.User, error) {
	return &users.User{}, nil
}

func (mock UserRepositoryMock) Update(user users.User) (*users.User, error) {
	return &users.User{}, nil
}

func (mock UserRepositoryMock) Delete(user users.User) (*users.User, error) {
	return &users.User{}, nil
}
