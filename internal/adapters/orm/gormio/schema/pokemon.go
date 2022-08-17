package schema

import (
	"database/sql"

	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
)

type Pokemon struct {
	PokemonSchema
	// relations
	Move                  []*Move               `gorm:"many2many:pokemons_moves;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`       // M:M
	BattleOpponentParty1  []BattleOpponentParty `gorm:"foreignKey:OpponentPokemonId1;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleOpponentParty
	BattleOpponentParty2  []BattleOpponentParty `gorm:"foreignKey:OpponentPokemonId2;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BattleOpponentParty3  []BattleOpponentParty `gorm:"foreignKey:OpponentPokemonId3;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BattleOpponentParty4  []BattleOpponentParty `gorm:"foreignKey:OpponentPokemonId4;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BattleOpponentParty5  []BattleOpponentParty `gorm:"foreignKey:OpponentPokemonId5;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BattleOpponentParty6  []BattleOpponentParty `gorm:"foreignKey:OpponentPokemonId6;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BattleRecordSelf1     []BattleRecord        `gorm:"foreignKey:SelfElectionPokemonId1;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecordSelf2     []BattleRecord        `gorm:"foreignKey:SelfElectionPokemonId2;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BattleRecordSelf3     []BattleRecord        `gorm:"foreignKey:SelfElectionPokemonId3;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BattleRecordSelf4     []BattleRecord        `gorm:"foreignKey:SelfElectionPokemonId4;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BattleRecordOpponent1 []BattleRecord        `gorm:"foreignKey:OpponentElectionPokemonId1;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecordOpponent2 []BattleRecord        `gorm:"foreignKey:OpponentElectionPokemonId2;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BattleRecordOpponent3 []BattleRecord        `gorm:"foreignKey:OpponentElectionPokemonId3;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BattleRecordOpponent4 []BattleRecord        `gorm:"foreignKey:OpponentElectionPokemonId4;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PokemonSchema struct {
	ID               uint16 `gorm:"primaryKey;autoIncrement:true"`
	PokedexNo        uint16
	FormNo           uint16
	FormName         string
	Name             string
	EnglishName      string
	Generation       uint16
	Type1            enum.PokemonType
	Type2            enum.PokemonType
	AbilityID1       sql.NullInt16 // has one
	AbilityID2       sql.NullInt16 // has one
	HiddenAbilityID  sql.NullInt16 // has one
	BaseStatsH       uint8
	BaseStatsA       uint8
	BaseStatsB       uint8
	BaseStatsC       uint8
	BaseStatsD       uint8
	BaseStatsS       uint8
	IsFinalEvolution bool `gorm:"default:false"`
}

func (Pokemon) TableName() string {
	return "pokemons"
}
