package schema

import (
	"gorm.io/gorm"
)

type PartyTag struct {
	gorm.Model
	Name         string
	IsGeneration *bool   `gorm:"default:false"`
	IsSeason     *bool   `gorm:"default:false"`
	Party        []Party `gorm:"many2many:parties_party_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // M:M
}

func (PartyTag) TableName() string {
	return "party_tags"
}
