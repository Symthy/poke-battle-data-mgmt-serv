package datastore

type IBasicRepository interface {
	FindById()
	FindAll()
}

type IPokemonRepository interface {
	IBasicRepository
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

type IPokemonQueryService interface {
	QueryPokemon()
}

type IAbilityQueryService interface {
	QueryAbility()
}

type IMoveQueryService interface {
	QueryMove()
}
