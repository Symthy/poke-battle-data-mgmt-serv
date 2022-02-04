//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/application/service/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/application/service/users"
	i_repository "github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository"
	"github.com/Symthy/PokeRest/pokeRest/presentation/controller"
	"github.com/Symthy/PokeRest/test/mock"
	"github.com/google/wire"
)

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
		pokemons.NewPokemonReadService,
		repository.NewPokemonRepository,
		wire.Bind(new(i_repository.IPokemonRepository), new(*repository.PokemonRepository)),
	)
	return nil
}

func InitPokemonControllerByRepoMock() *controller.PokemonController {
	wire.Build(
		controller.NewPokemonController,
		pokemons.NewPokemonReadService,
		mock.NewPokemonRepositoryMock,
		wire.Bind(new(i_repository.IPokemonRepository), new(*mock.PokemonRepositoryMock)),
	)
	return nil
}
