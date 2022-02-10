package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model/parties"
	"gorm.io/gorm"
)

type PartyTag struct {
	gorm.Model
	Name            string
	IsGenerationTag *bool   `gorm:"default:false"`
	IsSeasonTag     *bool   `gorm:"default:false"`
	Party           []Party `gorm:"many2many:party_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // M:M
}

func (PartyTag) TableName() string {
	return "party_tags"
}

func (t PartyTag) ConvertToDomain() parties.PartyTag {
	return parties.NewPartyTag(
		t.ID,
		t.Name,
		*t.IsGenerationTag,
		*t.IsSeasonTag,
	)
}
