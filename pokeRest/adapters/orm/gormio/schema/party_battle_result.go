package schema

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"

type PartyBattleResult struct {
	ID               uint `gorm:"primaryKey;autoIncrement:true"` // has many
	Generation       int
	Series           int
	Season           int
	MaxRate          int
	SeasonRanking    int
	MaxSeasonRanking int
	WinCount         int
	LoseCount        int
}

func (PartyBattleResult) TableName() string {
	return "party_battle_results"
}

func (p PartyBattleResult) ConvertToDomain() parties.PartyBattleResult {
	return parties.NewPartyBattleResult(p.ID, p.Generation, p.Series, p.Season, p.MaxRate, p.SeasonRanking, p.MaxSeasonRanking, p.WinCount, p.LoseCount)
}
