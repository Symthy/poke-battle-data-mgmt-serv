package ability

import "github.com/Symthy/PokeRest/pokeRest/application/datastore"

type AbilityReadService struct {
	repository datastore.IAbilityRepository
}
