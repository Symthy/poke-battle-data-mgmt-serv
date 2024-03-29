package users

import (
	"github.com/Symthy/PokeRest/internal/application/service/users/command"
	"github.com/Symthy/PokeRest/internal/domain/entity/users"
	"github.com/Symthy/PokeRest/internal/domain/repository"
)

type UserWriteService struct {
	repository repository.IUserRepository
}

func NewUserWriteService(repository repository.IUserRepository) UserWriteService {
	return UserWriteService{repository: repository}
}

// Todo: validate email or name duplicate
func (s UserReadService) SaveUser(command command.CreateUserCommand) (*users.User, error) {
	user := users.NewUserFromCommand(command)
	createdUser, err := s.repository.Create(user)
	return createdUser, err
}
