package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model/tags"
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

func (t Tag) ConvertToDomain() tags.Tag {
	return tags.NewTag(
		t.ID,
		t.Name,
		*t.IsGenerationTag,
		*t.IsSeasonTag,
	)
}
