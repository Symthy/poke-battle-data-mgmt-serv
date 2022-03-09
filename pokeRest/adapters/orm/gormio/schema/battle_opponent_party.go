package schema

type BattleOpponentParty struct {
	ID                 uint  `gorm:"primaryKey;autoIncrement:true"`
	OpponentPokemonId1 *uint // M:1 <- Pokemon
	OpponentPokemonId2 *uint
	OpponentPokemonId3 *uint
	OpponentPokemonId4 *uint
	OpponentPokemonId5 *uint
	OpponentPokemonId6 *uint
}

func (BattleOpponentParty) TableName() string {
	return "battle_opponent_parties"
}
