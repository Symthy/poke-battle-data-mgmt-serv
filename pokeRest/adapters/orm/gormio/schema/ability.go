package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"

type Ability struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"unique"`
	Description     string
	CorrectionValue mixin.CorrectionValue `gorm:"embedded"`

	// relation
	Pokemon1           []Pokemon            `gorm:"foreignKey:AbilityId1;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pokemon2           []Pokemon            `gorm:"foreignKey:AbilityId2;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pokemon3           []Pokemon            `gorm:"foreignKey:HiddenAbilityId;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TrainedPokemonBase []TrainedPokemonBase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Ability) TableName() string {
	return "abilities"
}
