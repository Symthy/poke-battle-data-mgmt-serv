package service

import (
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type UserWriteService struct {
	repository repository.IUserRepository
}

func NewUserWriteService(repository repository.IUserRepository) UserWriteService {
	return UserWriteService{repository: repository}
}

// Todo: input command
func (s UserReadService) CreateUser(command command.CreateUserCommand) (model.User, error) {
	user := model.NewUserFromCommand(command)
	createdUser, err := s.repository.Create(user)
	createdUser.ResetPassword()
	return createdUser, err
}
