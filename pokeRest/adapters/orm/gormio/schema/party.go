package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/common/collections"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"gorm.io/gorm"
)

type Party struct {
	gorm.Model
	Name           string
	BattleFormat   enum.BattleFormat
	IsPrivate      bool                `gorm:"default:false"`
	CreateUserId   *uint               // has many
	PartyResult    []PartyBattleResult `gorm:"foreignKey:PartyId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`   // 1:M -> PartyResult
	PartyTag       []PartyTag          `gorm:"many2many:parties_party_tags;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`     // M:M
	TrainedPokemon []TrainedPokemon    `gorm:"many2many:trained_pokemons_parties;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // M:M
}

func (Party) TableName() string {
	return "parties"
}

func (p Party) ConvertToDomain() parties.Party {
	var partyResultIds []uint = nil
	if p.PartyResult != nil {
		partyResultIds = collections.MapForList(
			p.PartyResult,
			func(result PartyBattleResult) uint {
				return result.ID
			})
	}
	var partyTagIds []uint = nil
	if p.PartyTag != nil {
		partyTagIds = collections.MapForList(
			p.PartyTag, func(tag PartyTag) uint {
				return tag.ID
			})
	}
	trainedPokemonIds := collections.MapForList(
		p.TrainedPokemon, func(t TrainedPokemon) uint {
			return t.ID
		})

	return parties.NewParty(p.ID, p.Name, p.BattleFormat.String(), p.IsPrivate, p.CreateUserId, &partyResultIds, &partyTagIds, trainedPokemonIds)
}
