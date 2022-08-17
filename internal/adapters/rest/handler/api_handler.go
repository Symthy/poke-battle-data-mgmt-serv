package handler

import (
	"sync"

	"github.com/Symthy/PokeRest/internal/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/internal/presentation/controller"
	"github.com/labstack/echo/v4"
)

// Todo: autogen
// required: delegate to controller

var _ server.ServerInterface = (*PokeRestHandler)(nil)

type PokeRestHandler struct {
	Lock                     sync.Mutex
	pokemonController        *controller.PokemonController
	abilityController        *controller.AbilityController
	moveController           *controller.MoveController
	itemController           *controller.ItemController
	typeController           *controller.TypeController
	partyTagController       *controller.PartyTagController
	partyController          *controller.PartyController
	battleRecordController   *controller.BattleRecordController
	trainedPokemonController *controller.TrainedPokemonController
	userController           *controller.UserController
}

func NewPokeRestHandler(
	pokemonCtrl *controller.PokemonController,
	abilityCtrl *controller.AbilityController,
	moveCtrl *controller.MoveController,
	itemCtrl *controller.ItemController,
	typeCtrl *controller.TypeController,
	partyTagCtrl *controller.PartyTagController,
	partyCtrl *controller.PartyController,
	battleRecordCtrl *controller.BattleRecordController,
	trainedPokeCtrl *controller.TrainedPokemonController,
	userCtrl *controller.UserController,
) *PokeRestHandler {
	return &PokeRestHandler{
		pokemonController:        pokemonCtrl,
		abilityController:        abilityCtrl,
		moveController:           moveCtrl,
		itemController:           itemCtrl,
		typeController:           typeCtrl,
		partyTagController:       partyTagCtrl,
		partyController:          partyCtrl,
		battleRecordController:   battleRecordCtrl,
		trainedPokemonController: trainedPokeCtrl,
		userController:           userCtrl,
	}
}

// get pokemons
// (GET /pokemons)
func (h *PokeRestHandler) GetPokemons(ctx echo.Context, params server.GetPokemonsParams) error {
	return nil
}

// get pokemon details by Pokedex No
// (GET /pokemons/{pokedexNo}/details)
func (h *PokeRestHandler) GetPokemonDetailsByPokedexNo(ctx echo.Context, pokedexNo uint64) error {
	return nil
}

// get pokemon details by Pokedex No and Form No
// (GET /pokemons/{pokedexNo}/{formNo}/details)
func (h *PokeRestHandler) GetPokemonDetailsByPokedexNoAndFormNo(ctx echo.Context, pokedexNo uint64, formNo uint64) error {
	return nil
}

// get pokemons of specify ability
// (GET /pokemons/abilities/{abilityId})
func (h *PokeRestHandler) GetPokemonsBySpecifiedAbility(ctx echo.Context, abilityId uint64) error {
	return nil
}

// get pokemons with specified move
// (GET /pokemons/moves/{moveId})
func (h *PokeRestHandler) GetPokemonsBySpecifiedMove(ctx echo.Context, moveId uint64) error {
	return nil
}

// get pokemon by No
// (GET /pokemons/{pokedexNo}/{formNo})
func (h *PokeRestHandler) GetPokemonByPokedexNoAndFormNo(ctx echo.Context, pokedexNo uint64, formNo uint64) error {
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
	var next uint64 = 0
	if params.Next != nil {
		next = *params.Next
	}
	return h.abilityController.GetAbilities(ctx, next, 0)
}

// get ability by id
// (GET /abilities/{avilityId})
func (h *PokeRestHandler) GetAbilityById(ctx echo.Context, avilityId uint64) error {
	return h.abilityController.GetAbilityById(ctx, avilityId)
}

// get abilities by specified pokemon id
// (GET /abilities/pokemons/{pokemonId})
func (h *PokeRestHandler) GetAbilitiesBySpecifedPokemonId(ctx echo.Context, pokemonId uint64) error {
	return h.abilityController.GetAbilityByPokemonId(ctx, pokemonId)
}

// get moves
// (GET /moves)
func (h *PokeRestHandler) GetMoves(ctx echo.Context, params server.GetMovesParams) error {
	var next uint64 = 0
	if params.Next != nil {
		next = *params.Next
	}
	return h.moveController.GetMoves(ctx, next, 0)
}

// get move by id
// (GET /moves/{moveId})
func (h *PokeRestHandler) GetMoveById(ctx echo.Context, moveId uint64) error {
	return h.moveController.GetMoveById(ctx, moveId)
}

// get moves by specified pokemon id
// (GET /moves/pokemons/{pokemonId})
func (h *PokeRestHandler) GetMovesBySpecifiedPokemonId(ctx echo.Context, pokemonId uint64) error {
	return h.moveController.GetMoveByPokemonId(ctx, pokemonId)
}

// get items
// (GET /items)
func (h *PokeRestHandler) GetHeldItems(ctx echo.Context, params server.GetHeldItemsParams) error {
	return h.itemController.GetHeldItems(ctx, *params.Next, 0)
}

// get held items by id
// (GET /helditems/{id})
func (h *PokeRestHandler) GetHeldItemsId(ctx echo.Context, id uint64) error {
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
	return h.typeController.GetDeffenseTypeCompatibility(ctx, pType)
}

/* transaction data */
// GET Parties
// (GET /parties)
func (h *PokeRestHandler) GetParties(ctx echo.Context, params server.GetPartiesParams) error {
	return nil
}

// POST party
// (POST /parties)
func (h *PokeRestHandler) PostParties(ctx echo.Context) error {
	return h.partyController.SaveParty(ctx)
}

// GET Party
// (GET /parties/{id})
func (h *PokeRestHandler) GetPartiesId(ctx echo.Context, id uint64) error {
	return nil
}

// PUT party
// (PUT /parties/{id})
func (h *PokeRestHandler) PutPartiesId(ctx echo.Context, id uint64) error {
	return h.partyController.UpdateParty(ctx, id)
}

// DELETE party
// (DELETE /parties/{id})
func (h *PokeRestHandler) DeletePartiesId(ctx echo.Context, id uint64) error {
	return h.partyController.DeleteParty(ctx, id) // Todo: fix
}

// GET party tags
// (GET /parties/tags)
func (h *PokeRestHandler) GetPartiesTags(ctx echo.Context, params server.GetPartiesTagsParams) error {
	return nil
}

// POST party tag
// (POST /parties/tags)
func (h *PokeRestHandler) PostPartiesTags(ctx echo.Context) error {
	return nil
}

// DELETE party tag
// (DELETE /parties/tags/{id})
func (h *PokeRestHandler) DeletePartiesTagsId(ctx echo.Context, id uint64) error {
	return nil
}

// GET battle record of Party
// (GET /battles/parties/{partyId})
func (h *PokeRestHandler) GetBattlesPartyId(ctx echo.Context, partyId uint64) error {
	return nil
}

// GET battle record of User
// (GET /battles/users/{userId})
func (h *PokeRestHandler) GetBattlesUserId(ctx echo.Context, userId uint64) error {
	return nil
}

// GET Battle record
// (GET /battles)
func (h *PokeRestHandler) GetBattles(ctx echo.Context) error {
	return nil
}

// POST Battle record
// (POST /battles)
func (h *PokeRestHandler) PostBattles(ctx echo.Context) error {
	return nil
}

// PUT Battle record
// (PUT /battles/{id})
func (h *PokeRestHandler) PutBattlesId(ctx echo.Context, id uint64) error {
	return nil
}

// DELETE battle record
// (DELETE /battles/{id})
func (h *PokeRestHandler) DeleteBattles(ctx echo.Context, id uint64) error {
	return h.battleRecordController.DeleteBattleRecord(ctx, id)
}

// GET trained pokemons
// (GET /trainedpokemons)
func (h *PokeRestHandler) GetTrainedPokemons(ctx echo.Context, params server.GetTrainedPokemonsParams) error {
	return nil
}

// POST trained pokemon
// (POST /trainedpokemons)
func (h *PokeRestHandler) PostTrainedpokemons(ctx echo.Context) error {
	return h.trainedPokemonController.SaveTrainedPokemon(ctx)
}

// GET trained pokemon
// (GET /trainedpokemons/{id})
func (h *PokeRestHandler) GetTrainedPokemonsId(ctx echo.Context, id uint64) error {
	return nil
}

// PUT trained pokemon
// (PUT /trainedpokemons/{id})
func (h *PokeRestHandler) PutTrainedpokemonsId(ctx echo.Context, id uint64) error {
	return h.trainedPokemonController.UpdateTrainedPokemon(ctx)
}

// DELETE trained pokemon
// (DELETE /trainedpokemons/{id})
func (h *PokeRestHandler) DeleteTrainedPokemonsId(ctx echo.Context, id uint64) error {
	return h.trainedPokemonController.DeleteTrainedPokemon(ctx, id)
}

// GET trained pokemon attack adjustments
// (GET /trainedpokemons/{id}/attacks)
func (h *PokeRestHandler) GetTrainedPokemonsIdAttacks(ctx echo.Context, id uint64) error {
	return nil
}

// POST trained pokemon attack adjustment
// (POST /trainedpokemons/{id}/attacks)
func (h *PokeRestHandler) PostTrainedpokemonsIdAttacks(ctx echo.Context, id uint64) error {
	return nil
}

// GET trained pokemon attack adjustment
// (GET /trainedpokemons/{trainedPokemonId}/attacks/{attackId})
func (h *PokeRestHandler) GetTrainedPokemonsIdAttacksId(ctx echo.Context, trainedPokemonId uint64, attackId uint64) error {
	return nil
}

// PUT trained pokemon attack adjustment
// (PUT /trainedpokemons/{trainedPokemonId}/attacks/{attackId})
func (h *PokeRestHandler) PutTrainedpokemonsTrainedPokemonIdAttacksAttackId(ctx echo.Context, trainedPokemonId uint64, attackId uint64) error {
	return nil
}

// DELETE trained pokemon attack adjustment
// (DELETE /trainedpokemons/{trainedPokemonId}/attacks/{attackId})
func (h *PokeRestHandler) DeleteTrainedPokemonsIdAttacksId(ctx echo.Context, trainedPokemonId uint64, attackId uint64) error {
	return nil
}

// (GET /trainedpokemons/{id}/deffenses)
func (h *PokeRestHandler) GetTrainedPokemonsIdDeffenses(ctx echo.Context, id uint64) error {
	return nil
}

// POST trained pokemon deffense adjustment
// (POST /trainedpokemons/{id}/deffenses)
func (h *PokeRestHandler) PostTrainedpokemonsIdDeffenses(ctx echo.Context, id uint64) error {
	return nil
}

// GET trained pokemon deffense adjustment
// (GET /trainedpokemons/{trainedPokemonId}/deffenses/{deffenseId})
func (h *PokeRestHandler) GetTrainedPokemonsIdDeffensesId(ctx echo.Context, trainedPokemonId uint64, deffenseId uint64) error {
	return nil
}

// PUT trained pokemon deffense adjustment
// (PUT /trainedpokemons/{trainedPokemonId}/deffenses/{deffenseId})
func (h *PokeRestHandler) PutTrainedpokemonsTrainedPokemonIdDeffensesDeffenseId(ctx echo.Context, trainedPokemonId uint64, deffenseId uint64) error {
	return nil
}

// DELETE trained pokemon deffense adjustment
// (DELETE /trainedpokemons/{trainedPokemonId}/deffenses/{deffenseId})
func (h *PokeRestHandler) DeleteTrainedPokemonsIdDeffensesId(ctx echo.Context, trainedPokemonId uint64, deffenseId uint64) error {
	return nil
}

// damage calculation (multi)
// (POST /damages/multi)
func (h *PokeRestHandler) PostDamagesMulti(ctx echo.Context) error {
	return nil
}

// damage calculation (single)
// (POST /damages/single)
func (h *PokeRestHandler) PostDamagesSingle(ctx echo.Context) error {
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

// GET user
// (GET /users/{id})
func (h *PokeRestHandler) GetUsersId(ctx echo.Context, id uint64) error {
	return h.userController.GetUserById(ctx, id)
}

// GET user parties
// (GET /users/{id}/parties)
func (h *PokeRestHandler) GetUsersIdParties(ctx echo.Context, id uint64, params server.GetUsersIdPartiesParams) error {
	return nil
}

// GET user battle records
// (GET /users/{id}/battles)
func (h *PokeRestHandler) GetUsersIdBattles(ctx echo.Context, id uint64, params server.GetUsersIdBattlesParams) error {
	return nil
}

// GET user trained pokemons
// (GET /users/{id}/trainedpokemons)
func (h *PokeRestHandler) GetUsersIdTrainedPokemons(ctx echo.Context, id uint64, params server.GetUsersIdTrainedPokemonsParams) error {
	return nil
}

// GET health check
// (GET /health)
func (h *PokeRestHandler) GetHealth(ctx echo.Context) error {
	return nil
}
