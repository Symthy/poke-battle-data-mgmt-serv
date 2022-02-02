package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name            string
	IsGenerationTag *bool   `gorm:"default:false"`
	IsSeasonTag     *bool   `gorm:"default:false"`
	Party           []Party `gorm:"many2many:party_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // M:M
}

func (Tag) TableName() string {
	return "tags"
}

func (t Tag) ConvertToDomain() *model.Tag {
	return &model.Tag{
		id:              t.Id,
		name:            t.Name,
		isGenerationTag: t.IsGeneration,
		isSeasonTag:     t.IsSeason,
	}
}
