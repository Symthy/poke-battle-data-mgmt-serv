package types

import "github.com/Symthy/PokeRest/pokeRest/domain/repository"

type TypeReadService struct {
	repository repository.ITypeRepository
}
