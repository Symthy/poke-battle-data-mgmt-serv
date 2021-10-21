package user

import "github.com/Symthy/PokeRest/pokeRest/application/datastore"

type UserReadService struct {
	repository datastore.IUserRepository
}
