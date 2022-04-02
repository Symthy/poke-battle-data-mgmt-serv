package damages

import "github.com/Symthy/PokeRest/pokeRest/common/fmath"

type BattleDamage struct {
	dataset PokemonBattleDataSet
	damages []int
}

func NewBattleDamage(battlePokemons PokemonBattleDataSet) BattleDamage {
	return BattleDamage{
		dataset: battlePokemons,
	}
}

func (d BattleDamage) calculate() int {
	if d.dataset.IsNoDamage() {
		return 0
	}

	// 威力
	powerValue := float64(d.dataset.ResolvePowerValue())
	// 最終攻撃
	attackValue := float64(d.dataset.ResolveAttackValue())
	// 最終防御
	defenseValue := float64(d.dataset.ResolveDefenseValue())
	// ダメージ補正

	// 最終ダメージ
	level := 50.0
	levelCorrection := fmath.RoundDown(level*2.0/5.0 + 2.0)
	baseDamage := fmath.RoundDown(fmath.RoundDown(levelCorrection*powerValue*attackValue/defenseValue)/50 + 2)

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
