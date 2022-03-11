package command

import "github.com/Symthy/PokeRest/pokeRest/domain/factory"

type DeleteBattleRecord struct {
	factory.BattleRecordInput
}

func NewDeleteBattleRecord(id uint, userId uint) DeleteBattleRecord {
	builder := factory.NewBattleRecordBuilder()
	return DeleteBattleRecord{builder}
}
