package battles

type BattleOpponentParty struct {
	id uint
}

func NewBattleOpponentParty(id uint) BattleOpponentParty {
	return BattleOpponentParty{
		id: id,
	}
}

func (b BattleOpponentParty) Id() uint {
	return b.id
}
