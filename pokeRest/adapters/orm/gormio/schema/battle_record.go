package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"gorm.io/gorm"
)

type BattleRecord struct {
	gorm.Model
	PartyId               uint
	Party                 Party `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // belongs to
	UserId                uint
	User                  User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Generation            int
	Series                int
	Season                int
	Result                enum.BattleResult
	BattleOpponentPartyId uint                // belongs to (感覚的に向き逆だがデータ流用のため許容)
	BattleOpponentParty   BattleOpponentParty `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`

	// 自身選出
	SelfElectionPokemonId1 *uint
	SelfTrainedPokemonId1  *uint
	SelfElectionPokemonId2 *uint
	SelfTrainedPokemonId2  *uint
	SelfElectionPokemonId3 *uint
	SelfTrainedPokemonId3  *uint
	SelfElectionPokemonId4 *uint
	SelfTrainedPokemonId4  *uint
	// 相手選出
	OpponentElectionPokemonId1 *uint
	OpponentElectionPokemonId2 *uint
	OpponentElectionPokemonId3 *uint
	OpponentElectionPokemonId4 *uint
}

func (BattleRecord) TableName() string {
	return "battle_records"
}
