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

func (p PartyBattleResults) TotalWinCount() int {
	total := 0
	for _, result := range p.items {
		total += result.getWinCount()
	}
	return total
}

func (p PartyBattleResults) TotalLoseCount() int {
	total := 0
	for _, result := range p.items {
		total += result.getLoseCount()
	}
	return total
}
