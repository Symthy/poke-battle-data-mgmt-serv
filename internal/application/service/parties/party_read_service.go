package parties

import "github.com/Symthy/PokeRest/internal/application/query"

type PartyReadService struct {
	qs query.IPartyQueryService
}

// UC: ユーザのパーティ一覧取得
func FindPartySummariesByUser(userId uint) {
}

// UC: パーティ詳細取得
func FindPartyById(partyId uint) {
}
