package battles

type BattleRecord struct {
	id uint
}

func NewBattleRecord(id uint) BattleRecord {
	return BattleRecord{
		id: id,
	}
}

func (b BattleRecord) Id() uint {
	return b.id
}
