package schema

type BattleOpponentParty struct {
	ID                 uint `gorm:"primaryKey;autoIncrement"`
	OpponentPokemonId1 int  // M:1 <- Pokemon
	OpponentPokemonId2 *int
	OpponentPokemonId3 *int
	OpponentPokemonId4 *int
	OpponentPokemonId5 *int
	OpponentPokemonId6 *int
	// relation
	BattleRecord []BattleRecord `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"` // 1:M -> BattleRecord
}

func (BattleOpponentParty) TableName() string {
	return "battle_opponent_party"
}
