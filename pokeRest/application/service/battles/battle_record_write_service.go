package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/repository"

type BattleRecordWriteServie struct {
	// Todo: 個別に分けるかどうか
	pr  repository.IPartyRepository
	br  repository.IBattleRecordRepository
	bor repository.IBattleOpponentPartyRepository
}

// UC: 戦績登録 (パーティ戦績も更新)
func (s BattleRecordWriteServie) AddBattleRecord() {
}

// UC: 戦績編集 (パーティ戦績も更新)
func (s BattleRecordWriteServie) EditBattleRecord() {
}

// UC: 戦績削除 (パーティ戦績も更新)
func (s BattleRecordWriteServie) DeleteBattleRecord() {
}
