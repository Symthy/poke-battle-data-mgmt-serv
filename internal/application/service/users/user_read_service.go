package users

import (
	"github.com/Symthy/PokeRest/internal/application/service/users/command"
	"github.com/Symthy/PokeRest/internal/domain/entity/users"
	"github.com/Symthy/PokeRest/internal/domain/repository"
)

type UserReadService struct {
	repository repository.IUserRepository
}

func NewUserReadService(repository repository.IUserRepository) UserReadService {
	return UserReadService{repository: repository}
}

func (s UserReadService) GetUserById(id uint64) (*users.User, error) {
	user, err := s.repository.FindById(id)
	return user, err
}

func (s UserReadService) GetUser(command command.GetUserCommand) (*users.User, error) {
	user, err := s.repository.FindByName(command.TargetName(), command.FilterFields()...)
	return user, err
}
