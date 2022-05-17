package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type BattleRecord struct {
	ID                    uint64 `gorm:"primaryKey;autoIncrement:true"`
	PartyID               uint64
	Party                 Party `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // belongs to
	UserID                uint64
	User                  User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SeasonID              uint16
	BattleSeason          BattleSeason `gorm:"foreignKey:SeasonID;references:id;"`
	Result                enum.BattleResult
	BattleOpponentPartyId uint64              // belongs to (感覚的に向き逆だがデータ流用のため許容)
	BattleOpponentParty   BattleOpponentParty `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	// 自身選出
	SelfElectionPokemonId1 *uint16
	SelfTrainedPokemonId1  *uint64
	SelfElectionPokemonId2 *uint16
	SelfTrainedPokemonId2  *uint64
	SelfElectionPokemonId3 *uint16
	SelfTrainedPokemonId3  *uint64
	SelfElectionPokemonId4 *uint16
	SelfTrainedPokemonId4  *uint64
	// 相手選出
	OpponentElectionPokemonId1 *uint16
	OpponentElectionPokemonId2 *uint16
	OpponentElectionPokemonId3 *uint16
	OpponentElectionPokemonId4 *uint16
	mixin.UpdateTimes
}

func (BattleRecord) TableName() string {
	return "battle_records"
}
