package schema

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/mixin"
)

type Ability struct {
	AbilitySchema

	// relations
	Pokemon1          []Pokemon                  `gorm:"foreignKey:AbilityID1;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> Pokemon
	Pokemon2          []Pokemon                  `gorm:"foreignKey:AbilityID2;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pokemon3          []Pokemon                  `gorm:"foreignKey:HiddenAbilityID;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TrainedAdjustment []TrainedPokemonAdjustment `gorm:"foreignKey:AbilityID;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M
}

type AbilitySchema struct {
	ID            uint16 `gorm:"primaryKey;autoIncrement:true"`
	Name          string `gorm:"unique"`
	Description   string
	BattleEffects *mixin.BattleEffects `gorm:"embedded"`
}

func (Ability) TableName() string {
	return "abilities"
}
