package damages

import "math"

type BattleDamage struct {
	battlePokemons PokemonBattleDataSet
	maxDamage      int
	minDamage      int
}

func NewBattleDamage(battlePokemons PokemonBattleDataSet) BattleDamage {
	return BattleDamage{
		battlePokemons: battlePokemons,
	}
}

func (d BattleDamage) calculate() int {
	if d.battlePokemons.IsNoDamage() {
		return 0
	}

	// 威力
	powerValue := float64(d.battlePokemons.ResolvePowerValue())
	// 最終攻撃
	attackValue := 0.0

	// 防御補正

	// 最終防御
	defenceValue := 0.0

	// ダメージ補正

	// 最終ダメージ
	level := 50.0
	levelCorrection := roundDown(level*2.0/5.0 + 2.0)
	baseDamage := roundDown(roundDown(levelCorrection*powerValue*attackValue/defenceValue)/50 + 2)

	damage := baseDamage
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
