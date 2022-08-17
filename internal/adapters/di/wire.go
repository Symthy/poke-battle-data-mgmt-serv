//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/application/service/abilities"
	"github.com/Symthy/PokeRest/internal/application/service/battles"
	"github.com/Symthy/PokeRest/internal/application/service/items"
	"github.com/Symthy/PokeRest/internal/application/service/moves"
	"github.com/Symthy/PokeRest/internal/application/service/parties"
	"github.com/Symthy/PokeRest/internal/application/service/pokemons"
	"github.com/Symthy/PokeRest/internal/application/service/trainings"
	"github.com/Symthy/PokeRest/internal/application/service/types"
	"github.com/Symthy/PokeRest/internal/application/service/users"
	i_repository "github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository"
	"github.com/Symthy/PokeRest/internal/presentation/auth"
	"github.com/Symthy/PokeRest/internal/presentation/controller"
	m_repo "github.com/Symthy/PokeRest/test/mock/repository"
	"github.com/google/wire"
)

// Bug: Can't auto generation, because wire unsupported generics?

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
		m_repo.NewUserRepositoryMock,
		wire.Bind(new(i_repository.IUserRepository), new(*m_repo.UserRepositoryMock)),
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
		m_repo.NewPokemonRepositoryMock,
		wire.Bind(new(i_repository.IPokemonRepository), new(*m_repo.PokemonRepositoryMock)),
	)
	return nil
}

/* Ability */
func InitAbilityController(dbClient orm.IDbClient) *controller.AbilityController {
	wire.Build(
		controller.NewAbilityController,
		abilities.NewAbilityReadService,
		repository.NewAbilityRepository,
		wire.Bind(new(i_repository.IAbilityRepository), new(*repository.AbilityRepository)),
	)
	return nil
}

/* Move */
func InitMoveController(dbClient orm.IDbClient) *controller.MoveController {
	wire.Build(
		controller.NewMoveController,
		moves.NewMoveReadService,
		repository.NewMoveRepository,
		wire.Bind(new(i_repository.IMoveRepository), new(*repository.MoveRepository)),
	)
	return nil
}

/* Item */
func InitItemController(dbClient orm.IDbClient) *controller.ItemController {
	wire.Build(
		controller.NewItemController,
		items.NewItemReadService,
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
func InitPartyTagController(dbClient orm.IDbClient) *controller.PartyTagController {
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
func InitPartyController(dbClient orm.IDbClient) *controller.PartyController {
	wire.Build(
		controller.NewPartyController,
		auth.NewAccessUserResolver,
		users.NewUserReadService,
		parties.NewPartyWriteService,
		repository.NewUserRepository,
		wire.Bind(new(i_repository.IUserRepository), new(*repository.UserRepository)),
		repository.NewPartyRepository,
		wire.Bind(new(i_repository.IPartyRepository), new(*repository.PartyRepository)),
	)
	return nil
}

/* TrainedPokemon */
func InitTrainedPokemonController(dbClient orm.IDbClient) *controller.TrainedPokemonController {
	wire.Build(
		controller.NewTrainedPokemonController,
		auth.NewAccessUserResolver,
		users.NewUserReadService,
		trainings.NewTrainedPokemonWriteService,
		repository.NewUserRepository,
		wire.Bind(new(i_repository.IUserRepository), new(*repository.UserRepository)),
		repository.NewTrainedPokemonRepository,
		wire.Bind(new(i_repository.ITrainedPokemonRepository), new(*repository.TrainedPokemonRepository)),
	)
	return nil
}

/* BattleRecord */
func InitBattleRecordController(dbClient orm.IDbClient) *controller.BattleRecordController {
	wire.Build(
		controller.NewBattleRecordController,
		auth.NewAccessUserResolver,
		users.NewUserReadService,
		battles.NewBattleRecordWriteService,
		repository.NewUserRepository,
		wire.Bind(new(i_repository.IUserRepository), new(*repository.UserRepository)),
		repository.NewPartyRepository,
		wire.Bind(new(i_repository.IPartyRepository), new(*repository.PartyRepository)),
		repository.NewBattleRecordRepository,
		wire.Bind(new(i_repository.IBattleRecordRepository), new(*repository.BattleRecordRepository)),
	)
	return nil
}
