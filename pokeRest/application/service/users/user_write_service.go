package users

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service/users/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type UserWriteService struct {
	repository repository.IUserRepository
}

func NewUserWriteService(repository repository.IUserRepository) UserWriteService {
	return UserWriteService{repository: repository}
}

func (s UserReadService) SaveUser(command command.CreateUserCommand) (model.User, error) {
	user := model.NewUserFromCommand(command)
	createdUser, err := s.repository.Save(user)
	createdUser.MaskPassword()
	return *createdUser, err
}
