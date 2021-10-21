package move

import "github.com/Symthy/PokeRest/pokeRest/application/datastore"

type MoveReadService struct {
	repository datastore.IMoveRepository
}
