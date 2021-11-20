package service

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type UserReadService struct {
	repository repository.IUserRepository
}

func NewUserReadService(repository repository.IUserRepository) UserReadService {
	return UserReadService{repository: repository}
}

func (s UserReadService) GetUser(id uint) (model.User, error) {
	user, err := s.repository.FindById(id)
	return user, err
}
