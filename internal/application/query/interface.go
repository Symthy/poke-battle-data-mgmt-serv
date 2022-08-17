package query

type IPokemonQueryService interface {
	QueryPokemonDetails()
}

type ITrainedPokemonQueryService interface {
	QueryTrainedPokemons(userId uint)
}

type IPartyQueryService interface {
	QueryParties(userId uint)
}

type IBattleRecordQueryService interface {
	QueryBattleRecords(partyId uint)
}
