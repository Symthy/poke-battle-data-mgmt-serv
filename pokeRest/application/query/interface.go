package datastore

type IPokemonQueryService interface {
	QueryPokemon()
}

type IAbilityQueryService interface {
	QueryAbility()
}

type IMoveQueryService interface {
	QueryMove()
}
