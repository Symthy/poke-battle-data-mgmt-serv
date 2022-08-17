package battles

import (
	"github.com/Symthy/PokeRest/internal/application/query"
	"github.com/Symthy/PokeRest/internal/application/service/battles/command"
)

type BattleRecordReadService struct {
	qs query.IBattleRecordQueryService
}

// UC: パーティの戦績取得（フィルター付き)
func (s BattleRecordReadService) FindBattleRecordByParty(cmd command.SearchBattleRecordByPartyCommand) {
}

// UC: ユーザの戦績取得（フィルター付き)
func (s BattleRecordReadService) FindBattleRecordByUser(cmd command.SearchBattleRecordByUserCommand) {
}
