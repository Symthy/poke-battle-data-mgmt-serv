package schema

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
