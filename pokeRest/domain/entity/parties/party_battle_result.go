package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.PartyBattleResultId] = (*PartyBattleResult)(nil)

type PartyBattleResult struct {
	id               identifier.PartyBattleResultId
	maxRate          int
	seasonRanking    int
	maxSeasonRanking int
	winCount         int
	loseCount        int
	battles.Season
}

func NewPartyBattleResult(id identifier.PartyBattleResultId, generation int, series int, season int, maxRate int, seasonRanking int, maxSeasonRanking int, winCount int, loseCount int) PartyBattleResult {
	return PartyBattleResult{
		id:               id,
		maxRate:          maxRate,
		seasonRanking:    seasonRanking,
		maxSeasonRanking: maxSeasonRanking,
		winCount:         winCount,
		loseCount:        loseCount,
		Season:           battles.NewSeason(generation, series, season),
	}
}

func (p PartyBattleResult) Id() identifier.PartyBattleResultId {
	return p.id
}
