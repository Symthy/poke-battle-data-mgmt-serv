package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/parties"
	"gorm.io/gorm"
)

type Party struct {
	gorm.Model
	CreateUserId uint // has many
	Name         string
	BattleFormat enum.BattleFormat
	PartyResult  []PartySeasonResult `gorm:"foreignKey:PartyId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 1:M -> PartyResult
	Tags         []Tag               `gorm:"many2many:party_tags;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`           // M:M

}

func (Party) TableName() string {
	return "party"
}

func (p Party) ConvertToDomain() parties.Party {
	return parties.NewParty(p.ID)
}
