package battles

import (
	"github.com/Symthy/PokeRest/internal/application/query"
	"github.com/Symthy/PokeRest/internal/application/service/battles/command"
)

type BattleRecordAggregateService struct {
	qs query.IBattleRecordQueryService
}

// UC: 戦績集計結果取得
func (s BattleRecordAggregateService) AggregateBattleRecord(cmd command.BattleRecordAggregationCommand) {

}
