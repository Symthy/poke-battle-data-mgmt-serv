package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"

type Ability struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	Description     string
	CorrectionValue mixin.CorrectionValue `gorm:"embedded"`

	// relation
	Pokemon1           []Pokemon            `gorm:"foreignKey:id;references:AbilityId1;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pokemon2           []Pokemon            `gorm:"foreignKey:id;references:AbilityId2;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pokemon3           []Pokemon            `gorm:"foreignKey:id;references:HiddenAbilityId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TrainedPokemonBase []TrainedPokemonBase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Ability) TableName() string {
	return "abilities"
}
