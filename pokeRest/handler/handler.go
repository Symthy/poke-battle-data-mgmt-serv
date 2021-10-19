package handler

import (
	"sync"

	"github.com/Symthy/PokeRest/pokeRest/autoGen/server"
	"github.com/labstack/echo/v4"
)

type PokemonHandler struct {
	Lock sync.Mutex
}

func NewPokemonHandler() *PokemonHandler {
	return &PokemonHandler{}
}

// get ability
// (GET /abilities)
func (h *PokemonHandler) GetAbility(ctx echo.Context) error {
	return nil
}

// get avilities of specify pokemon
// (GET /abilities/pokemons/{abilityId})
func (h *PokemonHandler) GetAvilitiesOfSpecifyPokemon(ctx echo.Context, abilityId float32) error {
	return nil
}

// get ability by id
// (GET /abilities/{avilityId})
func (h *PokemonHandler) GetAbilityById(ctx echo.Context, avilityId float32) error {
	return nil
}

// get items
// (GET /items)
func (h *PokemonHandler) GetItems(ctx echo.Context, params server.GetItemsParams) error {
	return nil
}

// get moves
// (GET /moves)
func (h *PokemonHandler) GetMoves(ctx echo.Context, params server.GetMovesParams) error {
	return nil
}

// get move by id
// (GET /moves/{moveId})
func (h *PokemonHandler) GetMovesById(ctx echo.Context, moveId float32) error {
	return nil
}

// get pokemons
// (GET /pokemons)
func (h *PokemonHandler) GetPokemons(ctx echo.Context, params server.GetPokemonsParams) error {
	return nil
}

// get moves of specify pokemon
// (GET /pokemons/moves/{pokedexNo}/{formNo})
func (h *PokemonHandler) GetMovesOfSpecifyPokemon(ctx echo.Context, pokedexNo float32, formNo float32) error {
	return nil
}

// Your GET endpoint
// (GET /pokemons/{pokedexNo}/{formNo})
func (h *PokemonHandler) GetPokemonsPokedexNoFormNo(ctx echo.Context, pokedexNo float32, formNo float32) error {
	return nil
}

// get types
// (GET /types)
func (h *PokemonHandler) GetTypes(ctx echo.Context, params server.GetTypesParams) error {
	return nil
}

// get type compability
// (GET /types/compability)
func (h *PokemonHandler) GetTypesCompability(ctx echo.Context) error {
	return nil
}

// get type compability of attack side
// (GET /types/compability/attack/{type})
func (h *PokemonHandler) GetTypeCompabilityOfAttackSide(ctx echo.Context, pType string) error {
	return nil
}

// get type compability of defense side
// (GET /types/compability/defense/{type})
func (h *PokemonHandler) GetTypeCompabilityOfDefenseSide(ctx echo.Context, pType string) error {
	return nil
}
