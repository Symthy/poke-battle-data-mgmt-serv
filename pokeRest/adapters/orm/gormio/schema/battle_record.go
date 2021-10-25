package schema

import "gorm.io/gorm"

type BattleRecord struct {
	gorm.Model
	PartyId               uint
	Party                 Party
	BattleOpponentPartyID uint
	BattleOpponentParty   BattleOpponentParty
	// 自身選出
	SelfElectionPokemonId1 uint
	SelfElectionPokemonId2 uint
	SelfElectionPokemonId3 uint
	SelfElectionPokemonId4 uint
	// 相手選出
	OpponentElectionPokemonId1 uint
	OpponentElectionPokemonId2 uint
	OpponentElectionPokemonId3 uint
	OpponentElectionPokemonId4 uint
}

func (BattleRecord) TableName() string {
	return "battle_record"
}
