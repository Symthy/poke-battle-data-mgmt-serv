package battles

type BattleRecord struct {
	id uint
}

func (b BattleRecord) Id() uint {
	return b.id
}
