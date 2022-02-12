package handler

import (
	"sync"

	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/presentation/controller"
	"github.com/labstack/echo/v4"
)

// Todo: autogen
// required: delegate to controller

var _ server.ServerInterface = (*PokeRestHandler)(nil)

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
func (h *PokeRestHandler) GetAbilities(ctx echo.Context, params server.GetAbilitiesParams) error {
	return nil
}

// get ability by id
// (GET /abilities/{avilityId})
func (h *PokeRestHandler) GetAbilityById(ctx echo.Context, avilityId float32) error {
	return nil
}

// get items
// (GET /items)
func (h *PokeRestHandler) GetHeldItems(ctx echo.Context, params server.GetHeldItemsParams) error {
	return nil
}

// get moves
// (GET /moves)
func (h *PokeRestHandler) GetMoves(ctx echo.Context, params server.GetMovesParams) error {
	return nil
}

// get move by id
// (GET /moves/{moveId})
func (h *PokeRestHandler) GetMoveById(ctx echo.Context, moveId float32) error {
	return nil
}

// GET Parties
// (GET /parties)
func (h *PokeRestHandler) GetParties(ctx echo.Context) error {
	return nil
}

// PUT Parties
// (PUT /parties)
func (h *PokeRestHandler) PutParties(ctx echo.Context) error {
	return nil
}

// GET Party
// (GET /parties/{id})
func (h *PokeRestHandler) GetPartiesId(ctx echo.Context, id string) error {
	return nil
}

// POST party
// (POST /parties/{id})
func (h *PokeRestHandler) PostPartiesId(ctx echo.Context, id string) error {
	return nil
}

// DELETE party
// (DELETE /parties/{id})
func (h *PokeRestHandler) DeletePartiesId(ctx echo.Context, id string) error {
	return nil
}

// GET party tags
// (GET /parties/tags)
func (h *PokeRestHandler) GetPartiesTags(ctx echo.Context) error {
	return nil
}

// PUT party tags
// (PUT /parties/tags)
func (h *PokeRestHandler) PutPartiesTags(ctx echo.Context) error {
	return nil
}

// DELETE party tag
// (DELETE /parties/tags/{id})
func (h *PokeRestHandler) DeletePartiesTagsId(ctx echo.Context, id string) error {
	return nil
}

// get pokemons
// (GET /pokemons)
func (h *PokeRestHandler) GetPokemons(ctx echo.Context, params server.GetPokemonsParams) error {
	return nil
}

// get pokemon details by Pokedex No
// (GET /pokemons/{pokedexNo}/details)
func (h *PokeRestHandler) GetPokemonDetailsByPokedexNo(ctx echo.Context, pokedexNo float32) error {
	return nil
}

// get pokemon details by Pokedex No and Form No
// (GET /pokemons/{pokedexNo}/{formNo}/details)
func (h *PokeRestHandler) GetPokemonDetailsByPokedexNoAndFormNo(ctx echo.Context, pokedexNo float32, formNo float32) error {
	return nil
}

// get pokemons of specify ability
// (GET /pokemons/abilities/{abilityId})
func (h *PokeRestHandler) GetPokemonsBySpecifiedAbility(ctx echo.Context, abilityId float32) error {
	return nil
}

// get pokemons with specified move
// (GET /pokemons/moves/{moveId})
func (h *PokeRestHandler) GetPokemonsBySpecifiedMove(ctx echo.Context, moveId float32) error {
	return nil
}

// get pokemon by No
// (GET /pokemons/{pokedexNo}/{formNo})
func (h *PokeRestHandler) GetPokemonByPokedexNoAndFormNo(ctx echo.Context, pokedexNo float32, formNo float32) error {
	// Todo: unsupported?
	return nil
}

// get pokemons details
// (GET /pokemons/details)
func (h *PokeRestHandler) GetPokemonsDetails(ctx echo.Context, params server.GetPokemonsDetailsParams) error {
	// Todo: unsupported?
	return nil
}

// get types
// (GET /types)
func (h *PokeRestHandler) GetTypes(ctx echo.Context) error {
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

/* transaction data */
// GET trained pokemons
// (GET /trainedpokemons)
func (h *PokeRestHandler) GetTrainedPokemons(ctx echo.Context) error {
	return nil
}

// PUT trained pokemon
// (PUT /trainedpokemons)
func (h *PokeRestHandler) PutTrainedPokemons(ctx echo.Context) error {
	return nil
}

// GET trained pokemon
// (GET /trainedpokemons/{id})
func (h *PokeRestHandler) GetTrainedPokemonsId(ctx echo.Context, id string) error {
	return nil
}

// POST trained pokemon
// (POST /trainedpokemons/{id})
func (h *PokeRestHandler) PostTrainedPokemonsId(ctx echo.Context, id string) error {
	return nil
}

// DELETE trained pokemon
// (DELETE /trainedpokemons/{id})
func (h *PokeRestHandler) DeleteTrainedPokemonsId(ctx echo.Context, id string) error {
	return nil
}

// GET trained pokemon attack adjustments
// (GET /trainedpokemons/{id}/attacks)
func (h *PokeRestHandler) GetTrainedPokemonsIdAttacks(ctx echo.Context, id string) error {
	return nil
}

// PUT trained pokemon attack adjustments
// (PUT /trainedpokemons/{id}/attacks)
func (h *PokeRestHandler) PutTrainedPokemonsIdAttacks(ctx echo.Context, id string) error {
	return nil
}

// GET trained pokemon attack adjustment
// (GET /trainedpokemons/{trainedPokemonId}/attacks/{attackId})
func (h *PokeRestHandler) GetTrainedPokemonsIdAttacksId(ctx echo.Context, trainedPokemonId string, attackId string) error {
	return nil
}

// POST trained pokemon attack adjustment
// (POST /trainedpokemons/{trainedPokemonId}/attacks/{attackId})
func (h *PokeRestHandler) PostTrainedPokemonsIdAttacksId(ctx echo.Context, trainedPokemonId string, attackId string) error {
	return nil
}

// DELETE trained pokemon attack adjustment
// (DELETE /trainedpokemons/{trainedPokemonId}/attacks/{attackId})
func (h *PokeRestHandler) DeleteTrainedPokemonsIdAttacksId(ctx echo.Context, trainedPokemonId string, attackId string) error {
	return nil
}

// (GET /trainedpokemons/{id}/deffences)
func (h *PokeRestHandler) GetTrainedPokemonsIdDeffences(ctx echo.Context, id string) error {
	return nil
}

// PUT trained pokemon deffence adjustments
// (PUT /trainedpokemons/{id}/deffences)
func (h *PokeRestHandler) PutTrainedPokemonsIdDeffences(ctx echo.Context, id string) error {
	return nil
}

// GET trained pokemon deffence adjustment
// (GET /trainedpokemons/{trainedPokemonId}/deffences/{deffenceId})
func (h *PokeRestHandler) GetTrainedPokemonsIdDeffencesId(ctx echo.Context, trainedPokemonId string, deffenceId string) error {
	return nil
}

// POST trained pokemon deffence adjustment
// (POST /trainedpokemons/{trainedPokemonId}/deffences/{deffenceId})
func (h *PokeRestHandler) PostTrainedPokemonsIdDeffencesId(ctx echo.Context, trainedPokemonId string, deffenceId string) error {
	return nil
}

// DELETE trained pokemon deffence adjustment
// (DELETE /trainedpokemons/{trainedPokemonId}/deffences/{deffenceId})
func (h *PokeRestHandler) DeleteTrainedPokemonsIdDeffencesId(ctx echo.Context, trainedPokemonId string, deffenceId string) error {
	return nil
}

// damage calculation
// (POST /damages)
func (h *PokeRestHandler) PostDamages(ctx echo.Context) error {
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

// Your GET endpoint
// (GET /users/{id})
func (h *PokeRestHandler) GetUsersId(ctx echo.Context, id float32) error {
	return h.userController.GetUserById(ctx, id)
}

// GET user parties
// (GET /users/{name}/parties)
func (h *PokeRestHandler) GetUsersIdParties(ctx echo.Context, name string) error {
	return nil
}

// GET user trained pokemons
// (GET /users/{name}/trainedpokemons)
func (h *PokeRestHandler) GetUsersIdTrainedPokemons(ctx echo.Context, name string) error {
	return nil
}
