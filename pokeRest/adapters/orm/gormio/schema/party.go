package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"gorm.io/gorm"
)

type Party struct {
	gorm.Model
	CreateUserId uint // has many
	Name         string
	BattleFormat enum.BattleFormat
	PartyResult  []PartyResult `gorm:"foreignKey:PartyId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 1:M -> PartyResult
	Tags         []Tag         `gorm:"many2many:party_tags;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`           // M:M

}

func (Party) TableName() string {
	return "party"
}
