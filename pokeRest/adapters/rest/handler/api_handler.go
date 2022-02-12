package handler

import (
	"sync"

	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/presentation/controller"
	"github.com/labstack/echo/v4"
)

// Todo: autogen
// required: delegate to controller

type PokeRestHandler struct {
	Lock              sync.Mutex
	pokemonController *controller.PokemonController
	userController    *controller.UserController
	typeController    *controller.TypeController
	abilityController *controller.AbilityController
}

func NewPokeRestHandler(
	pokemonCtrl *controller.PokemonController,
	abilityCtrl *controller.AbilityController,
	typeCtrl *controller.TypeController,
	userCtrl *controller.UserController,
) *PokeRestHandler {
	return &PokeRestHandler{
		pokemonController: pokemonCtrl,
		abilityController: abilityCtrl,
		typeController:    typeCtrl,
		userController:    userCtrl,
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
	return h.pokemonController.GetPokemon(ctx, id)
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

// SignIn
// (POST /signin)
func (h *PokeRestHandler) PostSignin(ctx echo.Context) error {
	return nil
}

// SignUp
// (POST /signup)
func (h *PokeRestHandler) PostSignup(ctx echo.Context) error {
	return nil
}

// get types
// (GET /types)
func (h *PokeRestHandler) GetTypes(ctx echo.Context, params server.GetTypesParams) error {
	return h.typeController.GetTypes(ctx)
}

// get type compability
// (GET /types/compability)
func (h *PokeRestHandler) GetTypesCompability(ctx echo.Context) error {
	return h.typeController.GetTypeCompatibility(ctx)
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
	return h.userController.GetUserById(ctx, id)
}
