package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/repository"

type ActualValueResolver struct {
	pokemonRepo  repository.IPokemonRepository
	heldItemRepo repository.IHeldItemRepository
}
