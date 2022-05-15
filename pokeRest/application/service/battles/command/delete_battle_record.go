package command

import "github.com/Symthy/PokeRest/pokeRest/domain/factory"

type DeleteBattleRecord struct {
	factory.BattleRecordInput
}

func NewDeleteBattleRecord(id uint64, userId uint64) DeleteBattleRecord {
	builder := factory.NewBattleRecordBuilder()
	return DeleteBattleRecord{builder}
}
