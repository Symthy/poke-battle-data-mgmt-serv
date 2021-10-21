package _type

import "github.com/Symthy/PokeRest/pokeRest/application/datastore"

type TypeReadService struct {
	repository datastore.ITypeRepository
}
