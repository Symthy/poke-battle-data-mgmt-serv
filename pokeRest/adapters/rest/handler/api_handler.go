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
	Lock               sync.Mutex
	pokemonController  *controller.PokemonController
	abilityController  *controller.AbilityController
	moveController     *controller.MoveController
	itemController     *controller.ItemController
	typeController     *controller.TypeController
	partyTagController *controller.PartyTagController
	partyController    *controller.PartyController
	userController     *controller.UserController
}

func NewPokeRestHandler(
	pokemonCtrl *controller.PokemonController,
	abilityCtrl *controller.AbilityController,
	moveCtrl *controller.MoveController,
	itemCtrl *controller.ItemController,
	typeCtrl *controller.TypeController,
	partyTagCtrl *controller.PartyTagController,
	partyCtrl *controller.PartyController,
	userCtrl *controller.UserController,
) *PokeRestHandler {
	return &PokeRestHandler{
		pokemonController:  pokemonCtrl,
		abilityController:  abilityCtrl,
		moveController:     moveCtrl,
		itemController:     itemCtrl,
		typeController:     typeCtrl,
		partyTagController: partyTagCtrl,
		partyController:    partyCtrl,
		userController:     userCtrl,
	}
}

// get pokemons
// (GET /pokemons)
func (h *PokeRestHandler) GetPokemons(ctx echo.Context, params server.GetPokemonsParams) error {
	return nil
}

// get pokemon details by Pokedex No
// (GET /pokemons/{pokedexNo}/details)
func (h *PokeRestHandler) GetPokemonDetailsByPokedexNo(ctx echo.Context, pokedexNo int) error {
	return nil
}

// get pokemon details by Pokedex No and Form No
// (GET /pokemons/{pokedexNo}/{formNo}/details)
func (h *PokeRestHandler) GetPokemonDetailsByPokedexNoAndFormNo(ctx echo.Context, pokedexNo int, formNo int) error {
	return nil
}

// get pokemons of specify ability
// (GET /pokemons/abilities/{abilityId})
func (h *PokeRestHandler) GetPokemonsBySpecifiedAbility(ctx echo.Context, abilityId int) error {
	return nil
}

// get pokemons with specified move
// (GET /pokemons/moves/{moveId})
func (h *PokeRestHandler) GetPokemonsBySpecifiedMove(ctx echo.Context, moveId int) error {
	return nil
}

// get pokemon by No
// (GET /pokemons/{pokedexNo}/{formNo})
func (h *PokeRestHandler) GetPokemonByPokedexNoAndFormNo(ctx echo.Context, pokedexNo int, formNo int) error {
	// Todo: unsupported?
	return nil
}

// get pokemons details
// (GET /pokemons/details)
func (h *PokeRestHandler) GetPokemonsDetails(ctx echo.Context, params server.GetPokemonsDetailsParams) error {
	// Todo: unsupported?
	return nil
}

// get abilities
// (GET /abilities)
func (h *PokeRestHandler) GetAbilities(ctx echo.Context, params server.GetAbilitiesParams) error {
	return h.abilityController.GetAbilities(ctx, int(*params.Next), 0)
}

// get ability by id
// (GET /abilities/{avilityId})
func (h *PokeRestHandler) GetAbilityById(ctx echo.Context, avilityId int) error {
	return h.abilityController.GetAbilityById(ctx, avilityId)
}

// get abilities by specified pokemon id
// (GET /abilities/pokemons/{pokemonId})
func (h *PokeRestHandler) GetAbilitiesBySpecifedPokemonId(ctx echo.Context, pokemonId int) error {
	return h.abilityController.GetAbilityByPokemonId(ctx, pokemonId)
}

// get moves
// (GET /moves)
func (h *PokeRestHandler) GetMoves(ctx echo.Context, params server.GetMovesParams) error {
	return h.moveController.GetMoves(ctx, *params.Next, 0)
}

// get move by id
// (GET /moves/{moveId})
func (h *PokeRestHandler) GetMoveById(ctx echo.Context, moveId int) error {
	return h.moveController.GetMoveById(ctx, moveId)
}

// get moves by specified pokemon id
// (GET /moves/pokemons/{pokemonId})
func (h *PokeRestHandler) GetMovesBySpecifiedPokemonId(ctx echo.Context, pokemonId int) error {
	return h.moveController.GetMoveByPokemonId(ctx, pokemonId)
}

// get items
// (GET /items)
func (h *PokeRestHandler) GetHeldItems(ctx echo.Context, params server.GetHeldItemsParams) error {
	return h.itemController.GetHeldItems(ctx, *params.Next, 0)
}

// get held items by id
// (GET /helditems/{id})
func (h *PokeRestHandler) GetHeldItemsId(ctx echo.Context, id int) error {
	return h.itemController.GetHeldItemById(ctx, id)
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
	return h.typeController.GetAttackTypeCompatibility(ctx, pType)
}

// get type compability of defense side
// (GET /types/compability/defense/{type})
func (h *PokeRestHandler) GetTypeCompabilityOfDefenseSide(ctx echo.Context, pType string) error {
	return h.typeController.GetDeffenceTypeCompatibility(ctx, pType)
}

// GET Parties
// (GET /parties)
func (h *PokeRestHandler) GetParties(ctx echo.Context, params server.GetPartiesParams) error {
	return nil
}

// PUT Parties
// (PUT /parties)
func (h *PokeRestHandler) PutParties(ctx echo.Context) error {
	return h.partyController.SaveParty(ctx)
}

// GET Party
// (GET /parties/{id})
func (h *PokeRestHandler) GetPartiesId(ctx echo.Context, id int) error {
	return nil
}

// POST party
// (POST /parties/{id})
func (h *PokeRestHandler) PostPartiesId(ctx echo.Context, id int) error {
	return h.partyController.UpdateParty(ctx, id)
}

// DELETE party
// (DELETE /parties/{id})
func (h *PokeRestHandler) DeletePartiesId(ctx echo.Context, id int) error {
	return h.partyController.DeleteParty(ctx, id)
}

// GET party tags
// (GET /parties/tags)
func (h *PokeRestHandler) GetPartiesTags(ctx echo.Context, params server.GetPartiesTagsParams) error {
	return nil
}

// PUT party tags
// (PUT /parties/tags)
func (h *PokeRestHandler) PutPartiesTags(ctx echo.Context) error {
	return nil
}

// DELETE party tag
// (DELETE /parties/tags/{id})
func (h *PokeRestHandler) DeletePartiesTagsId(ctx echo.Context, id int) error {
	return nil
}

/* transaction data */
// GET trained pokemons
// (GET /trainedpokemons)
func (h *PokeRestHandler) GetTrainedPokemons(ctx echo.Context, params server.GetTrainedPokemonsParams) error {
	return nil
}

// PUT trained pokemon
// (PUT /trainedpokemons)
func (h *PokeRestHandler) PutTrainedPokemons(ctx echo.Context) error {
	return nil
}

// GET trained pokemon
// (GET /trainedpokemons/{id})
func (h *PokeRestHandler) GetTrainedPokemonsId(ctx echo.Context, id int) error {
	return nil
}

// POST trained pokemon
// (POST /trainedpokemons/{id})
func (h *PokeRestHandler) PostTrainedPokemonsId(ctx echo.Context, id int) error {
	return nil
}

// DELETE trained pokemon
// (DELETE /trainedpokemons/{id})
func (h *PokeRestHandler) DeleteTrainedPokemonsId(ctx echo.Context, id int) error {
	return nil
}

// GET trained pokemon attack adjustments
// (GET /trainedpokemons/{id}/attacks)
func (h *PokeRestHandler) GetTrainedPokemonsIdAttacks(ctx echo.Context, id int) error {
	return nil
}

// PUT trained pokemon attack adjustments
// (PUT /trainedpokemons/{id}/attacks)
func (h *PokeRestHandler) PutTrainedPokemonsIdAttacks(ctx echo.Context, id int) error {
	return nil
}

// GET trained pokemon attack adjustment
// (GET /trainedpokemons/{trainedPokemonId}/attacks/{attackId})
func (h *PokeRestHandler) GetTrainedPokemonsIdAttacksId(ctx echo.Context, trainedPokemonId int, attackId int) error {
	return nil
}

// POST trained pokemon attack adjustment
// (POST /trainedpokemons/{trainedPokemonId}/attacks/{attackId})
func (h *PokeRestHandler) PostTrainedPokemonsIdAttacksId(ctx echo.Context, trainedPokemonId int, attackId int) error {
	return nil
}

// DELETE trained pokemon attack adjustment
// (DELETE /trainedpokemons/{trainedPokemonId}/attacks/{attackId})
func (h *PokeRestHandler) DeleteTrainedPokemonsIdAttacksId(ctx echo.Context, trainedPokemonId int, attackId int) error {
	return nil
}

// (GET /trainedpokemons/{id}/deffences)
func (h *PokeRestHandler) GetTrainedPokemonsIdDeffences(ctx echo.Context, id int) error {
	return nil
}

// PUT trained pokemon deffence adjustments
// (PUT /trainedpokemons/{id}/deffences)
func (h *PokeRestHandler) PutTrainedPokemonsIdDeffences(ctx echo.Context, id int) error {
	return nil
}

// GET trained pokemon deffence adjustment
// (GET /trainedpokemons/{trainedPokemonId}/deffences/{deffenceId})
func (h *PokeRestHandler) GetTrainedPokemonsIdDeffencesId(ctx echo.Context, trainedPokemonId int, deffenceId int) error {
	return nil
}

// POST trained pokemon deffence adjustment
// (POST /trainedpokemons/{trainedPokemonId}/deffences/{deffenceId})
func (h *PokeRestHandler) PostTrainedPokemonsIdDeffencesId(ctx echo.Context, trainedPokemonId int, deffenceId int) error {
	return nil
}

// DELETE trained pokemon deffence adjustment
// (DELETE /trainedpokemons/{trainedPokemonId}/deffences/{deffenceId})
func (h *PokeRestHandler) DeleteTrainedPokemonsIdDeffencesId(ctx echo.Context, trainedPokemonId int, deffenceId int) error {
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
func (h *PokeRestHandler) GetUsersId(ctx echo.Context, id string) error {
	return h.userController.GetUserById(ctx, id)
}

// GET user parties
// (GET /users/{id}/parties)
func (h *PokeRestHandler) GetUsersIdParties(ctx echo.Context, id string, params server.GetUsersIdPartiesParams) error {
	return nil
}

// GET user trained pokemons
// (GET /users/{id}/trainedpokemons)
func (h *PokeRestHandler) GetUsersIdTrainedPokemons(ctx echo.Context, id string, params server.GetUsersIdTrainedPokemonsParams) error {
	return nil
}
