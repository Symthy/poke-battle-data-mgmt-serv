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

	damage := resolveBaseDamage(calcElements)

	// ダブル補正

	// 天候
	damage = calcElements.ApplyWeatherCorrection(damage)

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

func resolveBaseDamage(calcElements DamageCalcElements) uint16 {
	// 最終威力
	powerValue := float64(calcElements.FinalAttackValue())
	// 最終攻撃
	attackValue := float64(calcElements.AttackActualValue())
	// 最終防御
	defenseValue := float64(calcElements.DefenseActualValue())

	level := 50.0
	levelCorrection := float64(fmath.RoundDown[uint16](level*2.0/5.0 + 2.0))
	damageBaseValue := float64(fmath.RoundDown[uint16](levelCorrection * powerValue * attackValue / defenseValue))
	return fmath.RoundDown[uint16](damageBaseValue/50 + 2)
}

func generateRandomDamage(maxDamage uint16) []uint16 {
	randamRange := []float64{
		0.85, 0.86, 0.87, 0.88, 0.89, 0.90, 0.91, 0.92, 0.93, 0.94, 0.95, 0.96, 0.97, 0.98, 0.99, 100.0,
	}
	damages := make([]uint16, len(randamRange))
	for i, randam := range randamRange {
		damages[i] = fmath.RoundDown[uint16](float64(maxDamage) * randam)
	}
	return damages
}
