package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"gorm.io/gorm"
)

type BattleRecord struct {
	gorm.Model
	PartyId               uint
	Party                 Party `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // belongs to
	Generation            int
	Series                int
	Season                int
	Result                enum.BattleResult
	BattleOpponentPartyId uint // has one (感覚的に向き逆だがデータ流用のため許容)
	// 自身選出
	SelfElectionPokemonId1 *int
	SelfTrainedPokemonId1  *uint
	SelfElectionPokemonId2 *int
	SelfTrainedPokemonId2  *uint
	SelfElectionPokemonId3 *int
	SelfTrainedPokemonId3  *uint
	SelfElectionPokemonId4 *int
	SelfTrainedPokemonId4  *uint
	// 相手選出
	OpponentElectionPokemonId1 *int
	OpponentElectionPokemonId2 *int
	OpponentElectionPokemonId3 *int
	OpponentElectionPokemonId4 *int
}

func (BattleRecord) TableName() string {
	return "battle_records"
}
