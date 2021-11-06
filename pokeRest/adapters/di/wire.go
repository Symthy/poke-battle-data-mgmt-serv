//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/application/service/pokemon"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/database"
	"github.com/Symthy/PokeRest/pokeRest/presentation/controller"
	"github.com/google/wire"
)

func initPokemonController() *controller.PokemonController {
	wire.Build(
		controller.NewPokemonController,
		pokemon.NewPokemonReadService,
		database.NewPokemonRepository,
		orm.NewGormDbClient,
		wire.Bind(new(repository.IPokemonRepository), new(*database.PokemonRepository)),
		wire.Bind(new(orm.IDbClient), new(*orm.GormDbClient)),
	)
	return nil
}
