package schema

type PartyResult struct {
	ID         uint `gorm:"primaryKey;autoIncrement"` // has many
	PartyId    uint
	Generation int
	Series     int
	Season     int
	MaxRate    int
	MaxRanking int
	WinCount   int
	LoseCount  int
}

func (PartyResult) TableName() string {
	return "party_result"
}
