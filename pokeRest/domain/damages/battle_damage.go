package damages

import "math"

type BattleDamage struct {
	battlePokemons BattlePokemons
	maxDamage      int
	minDamage      int
}

func NewBattleDamage(battlePokemons BattlePokemons) BattleDamage {
	return BattleDamage{
		battlePokemons: battlePokemons,
	}
}

func (d BattleDamage) calculate() int {
	// タイプ相性 ×0 なら return 0
	if d.battlePokemons.IsNoDamage() {
		return 0
	}

	// 威力
	powerValue := d.battlePokemons.attackSide.AttackPower()

	// 攻撃補正

	// 最終攻撃
	attackValue := 0.0

	// 防御補正

	// 最終防御
	defenceValue := 0.0

	// ダメージ補正

	// 最終ダメージ
	level := 50.0
	levelCorrection := roundDown(level*2.0/5.0 + 2.0)
	baseDamage = roundDown(roundDown(levelCorrection*powerValue*attackValue/defenceValue)/50 + 2)

	// ダブル補正

	// 天候

	// タイプ一致補正

	// タイプ相性 切り捨て

	// やけど

	// 乱数（最大/最小）切り捨て

	// 1未満は1にする
}

func resolvePower() {

}

func roundDown(value float64) float64 {
	// 小数点切り捨て
	return math.Trunc(value)
}

func roundUp(value float64) float64 {
	// 小数点切り上げ
	return math.Ceil(value)
}

func roundUpIfDecimalGreaterFive(value float64) float64 {
	integerValue := roundDown(value)
	if value-integerValue > 0.5 {
		return roundUp(value)
	}
	return roundDown(value)
}
