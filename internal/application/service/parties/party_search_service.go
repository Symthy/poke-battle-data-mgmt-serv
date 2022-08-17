package parties

import (
	"github.com/Symthy/PokeRest/internal/application/query"
	"github.com/Symthy/PokeRest/internal/application/service/parties/command"
)

type PartySearchService struct {
	qs query.IPartyQueryService
}

// UC: パーティ検索
func SearchParties(cmd command.SearchPartyCommand) {
}
