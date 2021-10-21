package item

import "github.com/Symthy/PokeRest/pokeRest/application/datastore"

type ItemReadService struct {
	repository datastore.IItemRepository
}
