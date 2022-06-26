package damages

import "github.com/Symthy/PokeRest/pokeRest/common/fmath"

type BattleDamage struct {
	calcElements DamageCalcElements
	damages      []uint16
}

func NewBattleDamage(calcElements DamageCalcElements) BattleDamage {
	return BattleDamage{
		calcElements: calcElements,
		damages:      calculate(calcElements),
	}
}

// Todo: method split
func calculate(calcElements DamageCalcElements) []uint16 {
	if calcElements.IsNoDamage() {
		return []uint16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	}

	// 最終威力
	powerValue := float64(calcElements.FinalAttackValue())
	// 最終攻撃
	attackValue := float64(calcElements.AttackActualValue())
	// 最終防御
	defenseValue := float64(calcElements.DefenseActualValue())

	// 最終ダメージ
	level := 50.0
	levelCorrection := float64(fmath.RoundDown[uint16](level*2.0/5.0 + 2.0))
	damageBaseValue := float64(fmath.RoundDown[uint16](levelCorrection * powerValue * attackValue / defenseValue))
	damage := fmath.RoundDown[uint16](damageBaseValue/50 + 2)

	// ダブル補正

	// 天候
	damage = fmath.RoundUpIfDecimalGreaterFive[uint16](float64(damage*calcElements.WeatherCorrectedValue()) / 4096.0)

	// タイプ一致補正
	if calcElements.IsTypeMatch() {
		damage = fmath.RoundUpIfDecimalGreaterFive[uint16](float64(damage) * 6144.0 / 4096.0)
	}

	// タイプ相性 切り捨て
	damage = fmath.Round[uint16](float64(damage) * calcElements.TypeCompatibilityDamageRate())

	// やけど

	// ダメージ補正
	damage = fmath.RoundUpIfDecimalGreaterFive[uint16](float64(damage * calcElements.DamageCorrectedValue() / 4096.0))

	// 乱数（最大/最小）切り捨て
	damages := generateRandomDamage(damage)

	// 1未満は1にする
	for i := 0; i < len(damages); i++ {
		if damages[i] < 1 {
			damages[i] = 1
		}
	}
	return damages
}

func generateRandomDamage(maxDamage uint16) []uint16 {
	return []uint16{
		fmath.RoundDown[uint16](float64(maxDamage) * 0.85),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.86),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.87),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.88),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.89),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.90),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.91),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.92),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.93),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.94),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.95),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.96),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.97),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.98),
		fmath.RoundDown[uint16](float64(maxDamage) * 0.99),
		maxDamage,
	}
}
