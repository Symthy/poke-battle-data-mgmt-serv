package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"gorm.io/gorm"
)

type Party struct {
	gorm.Model
	UserId        uint
	User          User
	Name          string
	BattleFormat  enum.BattleFormat
	PartyResultId uint
	PartyResult   PartyResult
	Tags          []*Tag `gorm:"many2many:party_tags;"`
}

func (Party) TableName() string {
	return "party"
}
