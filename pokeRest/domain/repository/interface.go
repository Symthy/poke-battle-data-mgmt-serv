package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
)

type IBasicRepository interface {
	FindById(id uint)
}

type IListRepository[L model.IDomains[T], T model.IDomain] interface {
	FindAll(page int, pageSize int) (*L, error)
}

type IValueOfPokemonRepository[L model.IDomains[T], T model.IDomain] interface {
	FindOfPokemon(pokemonId uint) (*L, error)
}

type IPokemonRepository interface {
	FindById(id uint) (*pokemons.Pokemon, error)
	FindAll() (pokemons.PokemonList, error)
	Create(pokemon *pokemons.Pokemon) (*pokemons.Pokemon, error)
}

type IAbilityRepository interface {
	IListRepository[abilities.Abilities, abilities.Ability]
	IValueOfPokemonRepository[abilities.Abilities, abilities.Ability]
}

type IMoveRepository interface {
	IListRepository[moves.Moves, moves.Move]
	IValueOfPokemonRepository[moves.Moves, moves.Move]
}

type ITypeRepository interface {
	IBasicRepository
}

type IItemRepository interface {
	IBasicRepository
}

type IUserRepository interface {
	FindById(id uint) (*model.User, error)
	FindByName(targetName string, filterFields ...string) (*model.User, error)
	Create(user model.User) (*model.User, error)
}
