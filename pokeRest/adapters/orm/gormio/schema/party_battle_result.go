package schema

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"

type PartyBattleResult struct {
	ID         uint `gorm:"primaryKey;autoIncrement:true"` // has many
	PartyId    uint
	Generation int
	Series     int
	Season     int
	MaxRate    int
	MaxRanking int
	WinCount   int
	LoseCount  int
}

func (PartyBattleResult) TableName() string {
	return "party_battle_results"
}

func (p PartyBattleResult) ConvertToDomain() parties.PartyBattleResult {
	return parties.NewPartyBattleResult(p.ID)
}
