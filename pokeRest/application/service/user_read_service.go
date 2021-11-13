package service

import "github.com/Symthy/PokeRest/pokeRest/domain/repository"

type UserReadService struct {
	repository repository.IUserRepository
}
