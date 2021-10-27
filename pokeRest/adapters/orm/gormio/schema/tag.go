package schema

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name            string
	IsGenerationTag *bool    `gorm:"default:false"`
	IsSeasonTag     *bool    `gorm:"default:false"`
	Party           []*Party `gorm:"many2many:party_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Many To Many
}

func (Tag) TableName() string {
	return "tag"
}
