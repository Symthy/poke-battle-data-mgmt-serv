package service

import (
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
func (s UserReadService) CreateUser(user *model.User) (model.User, error) {
	created, err := s.repository.Create(user)
	// password delete
	return created, err
}
