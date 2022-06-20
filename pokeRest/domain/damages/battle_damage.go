package damages

import "github.com/Symthy/PokeRest/pokeRest/common/fmath"

type BattleDamage struct {
	damageCalcPartialValues DamageCalcElements
	damages                 []uint16
}

func NewBattleDamage(damageCalcPartialValues DamageCalcElements) BattleDamage {
	return BattleDamage{
		damageCalcPartialValues: damageCalcPartialValues,
	}
}

func (d BattleDamage) calculate() int {
	if d.damageCalcPartialValues.dataSet.IsNoDamage() {
		return 0
	}

	// 威力
	powerValue := float64(d.damageCalcPartialValues.ResolvePowerValue())
	// 最終攻撃
	attackValue := float64(d.damageCalcPartialValues.ResolveAttackValue())
	// 最終防御
	defenseValue := float64(d.damageCalcPartialValues.ResolveDefenseValue())
	// ダメージ補正

	// 最終ダメージ
	level := 50.0
	levelCorrection := float64(fmath.RoundDown[uint16](level*2.0/5.0 + 2.0))
	damageBaseValue := float64(fmath.RoundDown[uint16](levelCorrection * powerValue * attackValue / defenseValue))
	damage := fmath.RoundDown[uint16](damageBaseValue/50 + 2)

	// ダブル補正

	// 天候

	// タイプ一致補正

	// タイプ相性 切り捨て

	// やけど

	// 乱数（最大/最小）切り捨て

	// 1未満は1にする
	if damage < 1 {
		return 1
	}

	return int(damage)
}
