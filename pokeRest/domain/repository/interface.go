package repository

import "github.com/Symthy/PokeRest/pokeRest/domain/model"

type IBasicRepository interface {
	FindById(id uint)
	FindAll()
}

type IPokemonRepository interface {
	FindById(id uint) (model.Pokemon, error)
	FindAll() (model.PokemonList, error)
	Create(pokemon *model.Pokemon) (model.Pokemon, error)
}

type IAbilityRepository interface {
	IBasicRepository
}

type IMoveRepository interface {
	IBasicRepository
}

type ITypeRepository interface {
	IBasicRepository
}

type IItemRepository interface {
	IBasicRepository
}

type IUserRepository interface {
	FindById(id uint) (model.User, error)
	FindByName(targetName string, filterFields ...string) (model.User, error)
	Create(user *model.User) (model.User, error)
}
