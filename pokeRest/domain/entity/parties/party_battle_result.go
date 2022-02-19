package parties

type PartyBattleResult struct {
	id               uint
	generation       int
	series           int
	season           int
	maxRate          int
	seasonRanking    int
	maxSeasonRanking int
	winCount         int
	loseCount        int
}

func NewPartyBattleResult(id uint, generation int, series int, season int, maxRate int, seasonRanking int, maxSeasonRanking int, winCount int, loseCount int) PartyBattleResult {
	return PartyBattleResult{
		id:               id,
		generation:       generation,
		series:           series,
		season:           season,
		maxRate:          maxRate,
		seasonRanking:    seasonRanking,
		maxSeasonRanking: maxSeasonRanking,
		winCount:         winCount,
		loseCount:        loseCount,
	}
}

func (p PartyBattleResult) Id() uint {
	return p.id
}
