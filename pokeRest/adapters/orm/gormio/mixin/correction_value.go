package mixin

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type CorrectionValue struct {
	// 技威力
	PhysicalMovePowerCorrectionValue int `gorm:"default:1"`
	SpecialMovePowerCorrectionValue  int `gorm:"default:1"`
	// 攻撃力
	AttackCorrectionValue        int `gorm:"default:1"`
	SpecialAttackCorrectionValue int `gorm:"default:1"`
	// 攻撃威力
	AttackPowerCorrectionValue        int `gorm:"default:1"`
	SpecialAttackPowerCorrectionValue int `gorm:"default:1"`
	// 防御力
	DeffenseCorrectionValue        int `gorm:"default:1"`
	SpecialDeffenseCorrectionValue int `gorm:"default:1"`
	// ダメージ
	DamageCorrectionType1  enum.PokemonType `sql:"type:pokemon_type"`
	DamageCorrectionValue1 int              `gorm:"default:1"`
	DamageCorrectionType2  enum.PokemonType `sql:"type:pokemon_type"`
	DamageCorrectionValue2 int              `gorm:"default:1"`
	// 重さ
	WeightCorrectionValue int `gorm:"default:1"`
}
