//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/application/service/abilities"
	"github.com/Symthy/PokeRest/pokeRest/application/service/moves"
	"github.com/Symthy/PokeRest/pokeRest/application/service/parties"
	"github.com/Symthy/PokeRest/pokeRest/application/service/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/application/service/types"
	"github.com/Symthy/PokeRest/pokeRest/application/service/users"
	i_repository "github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository"
	"github.com/Symthy/PokeRest/pokeRest/presentation/controller"
	"github.com/Symthy/PokeRest/test/mock"
	"github.com/google/wire"
)

// Bug: Don't auto generation, because wire unsupported generics

/* User */
func InitUserController(dbClient orm.IDbClient) *controller.UserController {
	wire.Build(
		controller.NewUserController,
		users.NewUserReadService,
		repository.NewUserRepository,
		wire.Bind(new(i_repository.IUserRepository), new(*repository.UserRepository)),
	)
	return nil
}

func InitUserControllerByRepoMock() *controller.UserController {
	wire.Build(
		controller.NewUserController,
		users.NewUserReadService,
		mock.NewUserRepositoryMock,
		wire.Bind(new(i_repository.IUserRepository), new(*mock.UserRepositoryMock)),
	)
	return nil
}

/* Pokemon */
func InitPokemonController(dbClient orm.IDbClient) *controller.PokemonController {
	wire.Build(
		controller.NewPokemonController,
		pokemons.NewPokemonSummaryReadService,
		repository.NewPokemonRepository,
		wire.Bind(new(i_repository.IPokemonRepository), new(*repository.PokemonRepository)),
	)
	return nil
}

func InitPokemonControllerByRepoMock() *controller.PokemonController {
	wire.Build(
		controller.NewPokemonController,
		pokemons.NewPokemonSummaryReadService,
		mock.NewPokemonRepositoryMock,
		wire.Bind(new(i_repository.IPokemonRepository), new(*mock.PokemonRepositoryMock)),
	)
	return nil
}

/* Ability */
func InitAbilityController() *controller.AbilityController {
	wire.Build(
		controller.NewAbilityController,
		abilities.NewAbilityReadService,
		repository.NewAbilityRepository,
		wire.Bind(new(i_repository.IAbilityRepository), new(*repository.AbilityRepository)),
	)
	return nil
}

/* Move */
func InitMoveController() *controller.MoveController {
	wire.Build(
		controller.NewMoveController,
		moves.NewMoveReadService,
		repository.NewMoveRepository,
		wire.Bind(new(i_repository.IMoveRepository), new(*repository.MoveRepository)),
	)
	return nil
}

/* Item */
func InitItemController() *controller.ItemController {
	wire.Build(
		controller.NewItemController,
		moves.NewItemReadService,
		repository.NewHeldItemRepository,
		wire.Bind(new(i_repository.IHeldItemRepository), new(*repository.HeldItemRepository)),
	)
	return nil
}

/* Types */
func InitTypeController() *controller.TypeController {
	wire.Build(
		controller.NewTypeController,
		types.NewTypeReadService,
	)
	return nil
}

/* PartyTag */
func InitItemController() *controller.ItemController {
	wire.Build(
		controller.NewPartyTagController,
		parties.NewPartyTagReadService,
		parties.NewPartyTagWriteService,
		repository.NewPartyTagRepository,
		wire.Bind(new(i_repository.IPartyTagRepository), new(*repository.PartyTagRepository)),
	)
	return nil
}

/* Party */
func InitPartyController() *controller.PartyController {
	wire.Build(
		controller.NewPartyController,
		parties.NewPartyWriteService,
		repository.NewPartyRepository,
		wire.Bind(new(i_repository.IPartyRepository), new(*repository.PartyRepository)),
	)
	return nil
}

/* TrainedPokemon */
