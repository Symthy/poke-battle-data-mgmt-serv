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
