package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/abilities"
)

type Ability struct {
	ID              uint   `gorm:"primaryKey;autoIncrement:true"`
	Name            string `gorm:"unique"`
	Description     string
	CorrectionValue mixin.CorrectionValue `gorm:"embedded"`

	// relation
	Pokemon1           []Pokemon            `gorm:"foreignKey:AbilityId1;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> Pokemon
	Pokemon2           []Pokemon            `gorm:"foreignKey:AbilityId2;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pokemon3           []Pokemon            `gorm:"foreignKey:HiddenAbilityId;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TrainedPokemonBase []TrainedPokemonBase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemonBase
}

func (Ability) TableName() string {
	return "abilities"
}

func (a Ability) ConvertToDomain() abilities.Ability {
	return abilities.NewAbility(a.ID)
}
