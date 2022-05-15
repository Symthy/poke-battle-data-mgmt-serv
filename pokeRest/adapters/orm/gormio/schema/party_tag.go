package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type PartyTag struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name         string
	IsGeneration *bool   `gorm:"default:false"`
	IsSeason     *bool   `gorm:"default:false"`
	Party        []Party `gorm:"many2many:parties_party_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // M:M
	mixin.UpdateTimes
}

func (PartyTag) TableName() string {
	return "party_tags"
}
