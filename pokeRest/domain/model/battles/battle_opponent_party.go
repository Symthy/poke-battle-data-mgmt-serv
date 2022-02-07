package battles

type BattleOpponentParty struct {
	id uint
}

func (b BattleOpponentParty) Id() uint {
	return b.id
}
