package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"gorm.io/gorm"
)

type Party struct {
	gorm.Model
	CreateUserId   uint // has many
	Name           string
	BattleFormat   enum.BattleFormat
	IsPrivate      bool                 `gorm:"default:false"`
	PartyResult    []*PartyBattleResult `gorm:"foreignKey:PartyId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`  // 1:M -> PartyResult
	PartyTag       []*PartyTag          `gorm:"many2many:party_tags;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`            // M:M
	TrainedPokemon []*TrainedPokemon    `gorm:"many2many:trained_pokemon_parties;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // M:M

}

func (Party) TableName() string {
	return "parties"
}

func (p Party) ConvertToDomain() parties.Party {
	return parties.NewParty(p.ID)
}
