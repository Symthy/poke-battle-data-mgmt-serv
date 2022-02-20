package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/pokeRest/common/collections"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type HeldItem struct {
	ID              uint `gorm:"primaryKey;autoIncrement:true"`
	Name            string
	Description     string
	CorrectionValue mixin.CorrectionValue `gorm:"embedded"`

	// relation
	TrainedPokemon []TrainedPokemon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> TrainedPokemon
}

func (HeldItem) TableName() string {
	return "held_items"
}

func (i HeldItem) ConvertToDomain() items.HeldItem {
	correctionValues := []value.CorrectionValue{}
	collections.AddsToList(&correctionValues,
		value.PhysicalMoveCorrectionValue(nil, i.CorrectionValue.PhysicalMovePowerCorrectionValue),
		value.SpecialMoveCorrectionValue(nil, i.CorrectionValue.SpecialMovePowerCorrectionValue),
		value.AttackCorrectionValue(nil, i.CorrectionValue.AttackCorrectionValue),
		value.SpecialAttackCorrectionValue(nil, i.CorrectionValue.SpecialAttackCorrectionValue),
		value.AttackPowerCorrectionValue(nil, i.CorrectionValue.AttackPowerCorrectionValue),
		value.SpecialAttackPowerCorrectionValue(nil, i.CorrectionValue.SpecialAttackPowerCorrectionValue),
		value.DamageCorrectionValue(i.CorrectionValue.DamageCorrectionType1.ConvertToDomain(), i.CorrectionValue.DamageCorrectionValue1),
		value.DamageCorrectionValue(i.CorrectionValue.DamageCorrectionType2.ConvertToDomain(), i.CorrectionValue.DamageCorrectionValue2),
		value.WeightCorrectionValue(nil, i.CorrectionValue.WeightCorrectionValue),
	)
	return items.NewHeldItem(i.ID, i.Name, i.Description, correctionValues)
}
