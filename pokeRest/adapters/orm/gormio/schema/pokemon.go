package schema

import (
	"database/sql"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/sqltype"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
)

type Pokemon struct {
	ID               uint `gorm:"primaryKey;autoIncrement:true"`
	PokedexNo        int
	FormNo           int
	FormName         string
	Name             string
	EnglishName      string
	Generation       int
	Type1            enum.PokemonType
	Type2            enum.PokemonType
	AbilityId1       sql.NullInt16 // has one
	AbilityId2       sql.NullInt16 // has one
	HiddenAbilityId  sql.NullInt16 // has one
	IsFinalEvolution bool          `gorm:"default:false"`

	// relation
	Move                  []*Move               `gorm:"many2many:pokemon_moves;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`        // M:M
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

func (Pokemon) TableName() string {
	return "pokemons"
}

func (p Pokemon) ConvertToDomain() model.Pokemon {
	return model.NewPokemon(
		p.ID,
		p.PokedexNo,
		p.FormNo,
		p.FormName,
		p.Name,
		p.EnglishName,
		p.Generation,
		p.Type1.String(),
		p.Type2.String(),
		sqltype.ResolveNullInt16(p.AbilityId1),
		sqltype.ResolveNullInt16(p.AbilityId2),
		sqltype.ResolveNullInt16(p.HiddenAbilityId),
		p.IsFinalEvolution,
	)
}

func (p Pokemon) ConvertToDomainNonId() model.Pokemon {
	return model.NewPokemonNonId(
		p.PokedexNo,
		p.FormNo,
		p.FormName,
		p.Name,
		p.EnglishName,
		p.Generation,
		p.Type1.String(),
		p.Type2.String(),
		sqltype.ResolveNullInt16(p.AbilityId1),
		sqltype.ResolveNullInt16(p.AbilityId2),
		sqltype.ResolveNullInt16(p.HiddenAbilityId),
		p.IsFinalEvolution,
	)
}
