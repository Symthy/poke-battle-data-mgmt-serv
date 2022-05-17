package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type Party struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name         string
	BattleFormat enum.BattleFormat
	IsPrivate    bool    `gorm:"default:false"`
	CreateUserId *uint64 // has many
	mixin.UpdateTimes

	PartyResult    []PartySeasonResult `gorm:"foreignKey:PartyID;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`   // 1:M
	PartyTag       []PartyTag          `gorm:"many2many:parties_party_tags;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`     // M:M
	TrainedPokemon []TrainedPokemon    `gorm:"many2many:trained_pokemons_parties;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // M:M
}

func (Party) TableName() string {
	return "parties"
}
