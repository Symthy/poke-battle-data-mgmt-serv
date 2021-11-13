package service

import "github.com/Symthy/PokeRest/pokeRest/domain/repository"

type MoveReadService struct {
	repository repository.IMoveRepository
}
