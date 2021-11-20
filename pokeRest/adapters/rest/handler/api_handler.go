package handler

import (
	"net/http"
	"sync"

	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/presentation/controller"
	"github.com/labstack/echo/v4"
)

type PokeRestHandler struct {
	Lock              sync.Mutex
	pokemonController *controller.PokemonController
	userController    *controller.UserController
}

func NewPokeRestHandler(
	pokeCon *controller.PokemonController,
	userCon *controller.UserController,
) *PokeRestHandler {
	return &PokeRestHandler{
		pokemonController: pokeCon,
		userController:    userCon,
	}
}

// get abilities
// (GET /abilities)
func (h *PokeRestHandler) GetAbilities(ctx echo.Context) error {
	return nil
}

// get ability by id
// (GET /abilities/{avilityId})
func (h *PokeRestHandler) GetAbilityById(ctx echo.Context, avilityId float32) error {
	return nil
}

// get items
// (GET /items)
func (h *PokeRestHandler) GetItems(ctx echo.Context, params server.GetItemsParams) error {
	return nil
}

// get moves
// (GET /moves)
func (h *PokeRestHandler) GetMoves(ctx echo.Context, params server.GetMovesParams) error {
	return nil
}

// get move with specified pokemon
// (GET /moves/pokemons/{pokedexNo}/{formNo})
func (h *PokeRestHandler) GetMovesWithSpecifiedPokemon(ctx echo.Context, pokedexNo float32, formNo float32) error {
	return nil
}

// get move by id
// (GET /moves/{moveId})
func (h *PokeRestHandler) GetMoveById(ctx echo.Context, moveId float32) error {
	return nil
}

// get pokemon by Id
// (GET /pokemons/{id})
func (h *PokeRestHandler) GetPokemonById(ctx echo.Context, id float32) error {
	user, err := h.pokemonController.GetPokemon(id)
	if err != nil {
		ctx.JSON(http.StatusOK, user)
		return nil
	}
	return err
}

// get pokemons
// (GET /pokemons)
func (h *PokeRestHandler) GetPokemons(ctx echo.Context, params server.GetPokemonsParams) error {
	return nil
}

// get pokemons details
// (GET /pokemons/details)
func (h *PokeRestHandler) GetPokemonsDetails(ctx echo.Context, params server.GetPokemonsDetailsParams) error {
	return nil
}

// get pokemon detail by Id
// (GET /pokemons/details/{id})
func (h *PokeRestHandler) GetPokemonDetailById(ctx echo.Context, id float32) error {
	return nil
}

// get pokemons of specify ability
// (GET /pokemons/abilities/{abilityId})
func (h *PokeRestHandler) GetPokemonsOfSpecifiedAbility(ctx echo.Context, abilityId float32) error {
	return nil
}

// get pokemons with specified move
// (GET /pokemons/moves/{moveId})
func (h *PokeRestHandler) GetPokemonsWithSpecifiedMove(ctx echo.Context, moveId float32) error {
	return nil
}

// get pokemon by No
// (GET /pokemons/{pokedexNo}/{formNo})
func (h *PokeRestHandler) GetPokemonByPokedexNoAndFormNo(ctx echo.Context, pokedexNo float32, formNo float32) error {
	return nil
}

// get types
// (GET /types)
func (h *PokeRestHandler) GetTypes(ctx echo.Context, params server.GetTypesParams) error {
	return nil
}

// get type compability
// (GET /types/compability)
func (h *PokeRestHandler) GetTypesCompability(ctx echo.Context) error {
	return nil
}

// get type compability of attack side
// (GET /types/compability/attack/{type})
func (h *PokeRestHandler) GetTypeCompabilityOfAttackSide(ctx echo.Context, pType string) error {
	return nil
}

// get type compability of defense side
// (GET /types/compability/defense/{type})
func (h *PokeRestHandler) GetTypeCompabilityOfDefenseSide(ctx echo.Context, pType string) error {
	return nil
}

// Your GET endpoint
// (GET /users/{id})
func (h *PokeRestHandler) GetUsersId(ctx echo.Context, id float32) error {
	user, err := h.userController.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusOK, user)
		return nil
	}
	return err
}
