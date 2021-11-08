package schema

import (
	"database/sql"
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
)

type PokemonDto struct {
	ID               int `gorm:"primaryKey;autoIncrement"`
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
}

func (p PokemonDto) ConvertToDomain() model.Pokemon {
	// Todo: refactor
	value1, _ := p.AbilityId1.Value()
	var ability1 *int = nil
	if value1 != nil {
		convertVal := int(value1.(int64))
		ability1 = &convertVal
	}

	value2, _ := p.AbilityId1.Value()
	var ability2 *int = nil
	if value2 != nil {
		convertVal := int(value2.(int64))
		ability2 = &convertVal
	}

	value3, _ := p.AbilityId1.Value()
	var ability3 *int = nil
	if value3 != nil {
		convertVal := int(value3.(int64))
		ability3 = &convertVal
	}
	fmt.Printf("%#v\n", value1)
	fmt.Printf("%#v\n", value2)
	fmt.Printf("%#v\n", value3)
	fmt.Printf("%#v\n", ability1)
	fmt.Printf("%#v\n", ability2)
	fmt.Printf("%#v\n", ability3)
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
		ability1,
		ability2,
		ability3,
		p.IsFinalEvolution,
	)
}

type Pokemon struct {
	PokemonDto `gorm:"embedded"`

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
