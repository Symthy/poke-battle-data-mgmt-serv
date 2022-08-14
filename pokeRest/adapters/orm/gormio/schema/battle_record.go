package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type BattleRecord struct {
	BattleRecordSchema
	mixin.UpdateTimes

	// relations
	Party               Party               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // belongs to
	User                User                `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BattleSeason        BattleSeason        `gorm:"foreignKey:SeasonID;references:id;"`
	BattleOpponentParty BattleOpponentParty `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

type BattleRecordSchema struct {
	ID                    uint64 `gorm:"primaryKey;autoIncrement:true"`
	PartyID               uint64 // ref: Party
	UserID                uint64 // ref: User
	SeasonID              uint16 // ref: BattleSeason
	BattleFormat          enum.BattleFormat
	Result                enum.BattleResult
	BattleOpponentPartyId uint64 // ref: BattleOpponentParty belongs to (感覚的に向き逆だがデータ流用のため許容)
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
}

func (BattleRecord) TableName() string {
	return "battle_records"
}
