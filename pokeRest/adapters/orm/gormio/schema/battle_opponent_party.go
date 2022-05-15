package schema

type BattleOpponentParty struct {
	ID                 uint64  `gorm:"primaryKey;autoIncrement:true"`
	OpponentPokemonId1 *uint16 // M:1 <- Pokemon
	OpponentPokemonId2 *uint16
	OpponentPokemonId3 *uint16
	OpponentPokemonId4 *uint16
	OpponentPokemonId5 *uint16
	OpponentPokemonId6 *uint16
}

func (BattleOpponentParty) TableName() string {
	return "battle_opponent_parties"
}
