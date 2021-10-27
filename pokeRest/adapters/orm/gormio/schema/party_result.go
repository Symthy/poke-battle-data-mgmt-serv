package schema

type PartyResult struct {
	ID         uint `gorm:"primaryKey; autoIncrement"`
	Generation int
	Seasons    int
	MaxRate    int
	MaxRanking int
	PartyId    uint // has many
}

func (PartyResult) TableName() string {
	return "party_result"
}
