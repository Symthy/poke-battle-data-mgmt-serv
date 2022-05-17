package schema

type PartySeasonResult struct {
	ID               uint64 `gorm:"primaryKey;autoIncrement:true"`
	PartyID          uint64
	Generation       uint16
	Series           uint16
	Season           uint16
	MaxRate          uint16
	SeasonRanking    uint16
	MaxSeasonRanking uint16
	WinCount         uint16
	LoseCount        uint16
}

func (PartySeasonResult) TableName() string {
	return "party_season_results"
}
