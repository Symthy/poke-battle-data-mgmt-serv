package schema

type PartySeasonResult struct {
	PartyID          uint64 `gorm:"primaryKey;"`
	ResultID         uint64 `gorm:"primaryKey;"`
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
	return "party_battle_results"
}
