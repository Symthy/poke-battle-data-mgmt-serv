package schema

import "github.com/Symthy/PokeRest/pokeRest/domain/model/parties"

type PartySeasonResult struct {
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

func (PartySeasonResult) TableName() string {
	return "party_season_result"
}

func (p PartySeasonResult) ConvertToDomain() parties.PartySeasonResult {
	return parties.NewPartySeasonResult(p.ID)
}
