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
func InitUserController() *controller.UserController {
	wire.Build(
		controller.NewUserController,
		service.NewUserReadService,
		database.NewUserRepository,
		orm.NewGormDbClient,
		wire.Bind(new(repository.IUserRepository), new(*database.UserRepository)),
		wire.Bind(new(orm.IDbClient), new(*orm.GormDbClient)),
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
func InitPokemonController() *controller.PokemonController {
	wire.Build(
		controller.NewPokemonController,
		service.NewPokemonReadService,
		database.NewPokemonRepository,
		orm.NewGormDbClient,
		wire.Bind(new(repository.IPokemonRepository), new(*database.PokemonRepository)),
		wire.Bind(new(orm.IDbClient), new(*orm.GormDbClient)),
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
