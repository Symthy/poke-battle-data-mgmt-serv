package mock

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/test/data"
	"gorm.io/gorm"
)

// implements IUserRepository
type UserRepositoryMock struct {
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (mock UserRepositoryMock) FindById(id uint) (*model.User, error) {
	dummyUser := data.DummyUser1()
	if dummyUser.ID != id {
		return &model.User{}, nil
	}
	u := dummyUser.ConvertToDomain()
	return &u, nil
}

func (mock UserRepositoryMock) FindByName(targetName string, filterFields ...string) (*model.User, error) {
	dummyUser := data.DummyUser1(filterFields...)
	if dummyUser.Name != targetName {
		return nil, gorm.ErrRecordNotFound
	}
	u := dummyUser.ConvertToDomain()
	return &u, nil
}

func (mock UserRepositoryMock) Create(user model.User) (*model.User, error) {
	return &model.User{}, nil
}
