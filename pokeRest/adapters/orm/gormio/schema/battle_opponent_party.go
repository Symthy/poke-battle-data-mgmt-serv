package schema

type BattleOpponentParty struct {
	ID                 uint  `gorm:"primaryKey;autoIncrement:true"`
	OpponentPokemonId1 *uint // M:1 <- Pokemon
	OpponentPokemonId2 *uint
	OpponentPokemonId3 *uint
	OpponentPokemonId4 *uint
	OpponentPokemonId5 *uint
	OpponentPokemonId6 *uint
	// relation
	BattleRecord []BattleRecord `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"` // 1:M -> BattleRecord
}

func (BattleOpponentParty) TableName() string {
	return "battle_opponent_party"
}
