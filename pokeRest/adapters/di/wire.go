//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/database"
	"github.com/Symthy/PokeRest/pokeRest/presentation/controller"
	"github.com/Symthy/PokeRest/test/mock"
	"github.com/google/wire"
)

/* User */
func InitUserController(dbClient orm.IDbClient) *controller.UserController {
	wire.Build(
		controller.NewUserController,
		service.NewUserReadService,
		database.NewUserRepository,
		wire.Bind(new(repository.IUserRepository), new(*database.UserRepository)),
	)
	return nil
}

func InitUserControllerByRepoMock() *controller.UserController {
	wire.Build(
		controller.NewUserController,
		service.NewUserReadService,
		mock.NewUserRepositoryMock,
		wire.Bind(new(repository.IUserRepository), new(*mock.UserRepositoryMock)),
	)
	return nil
}

/* Pokemon */
func InitPokemonController(dbClient orm.IDbClient) *controller.PokemonController {
	wire.Build(
		controller.NewPokemonController,
		service.NewPokemonReadService,
		database.NewPokemonRepository,
		wire.Bind(new(repository.IPokemonRepository), new(*database.PokemonRepository)),
	)
	return nil
}

func InitPokemonControllerByRepoMock() *controller.PokemonController {
	wire.Build(
		controller.NewPokemonController,
		service.NewPokemonReadService,
		mock.NewPokemonRepositoryMock,
		wire.Bind(new(repository.IPokemonRepository), new(*mock.PokemonRepositoryMock)),
	)
	return nil
}
