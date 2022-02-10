package parties

type PartyBattleResult struct {
	id uint
}

func NewPartyBattleResult(id uint) PartyBattleResult {
	return PartyBattleResult{
		id: id,
	}
}

func (p PartyBattleResult) Id() uint {
	return p.id
}
