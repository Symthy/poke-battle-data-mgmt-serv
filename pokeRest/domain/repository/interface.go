package repository

import "github.com/Symthy/PokeRest/pokeRest/domain/model"

type IBasicRepository interface {
	FindById(id int)
	FindAll()
}

type IPokemonRepository interface {
	FindById(id int) model.Pokemon
	FindAll() (*model.PokemonList, error)
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
	IBasicRepository
}
