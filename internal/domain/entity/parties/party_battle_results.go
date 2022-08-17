package parties

type PartyBattleResults struct {
	items []PartyBattleResult
}

func NewPartyBattleResults(items []PartyBattleResult) PartyBattleResults {
	return PartyBattleResults{items: items}
}

func (p PartyBattleResults) Items() []PartyBattleResult {
	return p.items
}

func (p PartyBattleResults) TotalWinCount() uint64 {
	total := uint64(0)
	for _, result := range p.items {
		total += uint64(result.getWinCount())
	}
	return total
}

func (p PartyBattleResults) TotalLoseCount() uint64 {
	total := uint64(0)
	for _, result := range p.items {
		total += uint64(result.getLoseCount())
	}
	return total
}
