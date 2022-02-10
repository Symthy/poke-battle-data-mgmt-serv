package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/application/query"
	"github.com/Symthy/PokeRest/pokeRest/application/service/parties/command"
)

type PartySearchService struct {
	qs query.IPartyQueryService
}

// UC: パーティ検索
func SearchParties(cmd command.PartySearchCommand) {
}
