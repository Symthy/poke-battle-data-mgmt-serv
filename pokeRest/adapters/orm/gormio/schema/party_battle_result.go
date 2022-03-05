package schema

type PartySeasonResult struct {
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

func (PartySeasonResult) TableName() string {
	return "party_battle_results"
}
