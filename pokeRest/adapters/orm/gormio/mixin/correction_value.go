package mixin

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type CorrectionValue struct {
	// 技威力
	PhysicalMovePowerCorrectionValue float32 `gorm:"default:1"`
	SpecialMovePowerCorrectionValue  float32 `gorm:"default:1"`
	// 攻撃力
	AttackCorrectionValue        float32 `gorm:"default:1"`
	SpecialAttackCorrectionValue float32 `gorm:"default:1"`
	// 攻撃威力
	AttackPowerCorrectionValue        float32 `gorm:"default:1"`
	SpecialAttackPowerCorrectionValue float32 `gorm:"default:1"`
	// 防御力
	DeffenceCorrectionValue        float32 `gorm:"default:1"`
	SpecialDeffenceCorrectionValue float32 `gorm:"default:1"`
	// ダメージ
	DamageCorrectionType1  enum.PokemonType `sql:"type:pokemon_type"`
	DamageCorrectionValue1 float32          `gorm:"default:1"`
	DamageCorrectionType2  enum.PokemonType `sql:"type:pokemon_type"`
	DamageCorrectionValue2 float32          `gorm:"default:1"`
	// 重さ
	WeightCorrectionValue float32 `gorm:"default:1"`
}
