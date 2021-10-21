package pokemon

import "github.com/Symthy/PokeRest/pokeRest/application/datastore"

type PokemonReadService struct {
	repository datastore.IAbilityRepository
}
