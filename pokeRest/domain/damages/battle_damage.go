package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

type BattleDamage struct {
	calcElements battles.PokemonBattleBaseParams
	damages      []uint16
}

func NewBattleDamage(calcElements battles.PokemonBattleBaseParams) BattleDamage {
	return BattleDamage{
		calcElements: calcElements,
		damages:      calculate(calcElements),
	}
}

// Todo: method split
func calculate(calcElements battles.PokemonBattleBaseParams) []uint16 {
	if calcElements.IsNoDamage() {
		return []uint16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	}

	damage := calcBaseDamage(calcElements)

	// ダブル補正

	// 天候
	damage = calcDamageCorrectedWeather(damage, calcElements)

	// タイプ一致補正
	damage = calcDamageCorrectedTypeMatch(damage, calcElements)

	// タイプ相性 切り捨て
	damage = calcDamageCorrectedTypeCompatibility(damage, calcElements)

	// やけど
	damage = calcDamageCorrectedBurn(damage, calcElements)

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

func calcBaseDamage(calcElements battles.PokemonBattleBaseParams) uint16 {
	// 最終威力
	attackPower := calcElements.PowerCorrectedValue()
	attackPower = fmath.Round[uint16](float64(attackPower*calcElements.FieldCorrectedValue()) / 4096.0)
	movePower := calcElements.MovePowerValue()
	finalPowerValue := fmath.RoundUpIfDecimalGreaterFive[uint16](float64(movePower*attackPower) / 4096.0)
	// 最終攻撃
	finalAttackValue := calcElements.AttackCorrectedActualValue()
	// 最終防御
	finalDefenseValue := calcElements.DefenseCorrectedActualValue()

	level := 50.0
	levelCorrection := float64(fmath.RoundDown[uint16](level*2.0/5.0 + 2.0))
	damageBaseValue := float64(fmath.RoundDown[uint16](
		levelCorrection * float64(finalPowerValue*finalAttackValue) / float64(finalDefenseValue)))
	return fmath.RoundDown[uint16](damageBaseValue/50 + 2)
}

func calcDamageCorrectedWeather(damage uint16, calcElements battles.PokemonBattleBaseParams) uint16 {
	return fmath.RoundUpIfDecimalGreaterFive[uint16](float64(damage * calcElements.WeatherCorrectedValue()))
}

func calcDamageCorrectedTypeMatch(damage uint16, calcElements battles.PokemonBattleBaseParams) uint16 {
	result := damage
	if calcElements.IsTypeMatchAttackSide() {
		result = fmath.RoundUpIfDecimalGreaterFive[uint16](float64(damage) * 6144.0 / 4096.0)
	}
	return result
}

func calcDamageCorrectedTypeCompatibility(damage uint16, calcElements battles.PokemonBattleBaseParams) uint16 {
	return fmath.Round[uint16](float64(damage) * float64(calcElements.TypeCompatibilityDamageRate()))
}

func calcDamageCorrectedBurn(damage uint16, calcElements battles.PokemonBattleBaseParams) uint16 {
	result := damage
	if calcElements.IsBurnAttackSide() {
		correctionValue := calcElements.AbnormalStateAttackSideCorectedValue()
		result = fmath.RoundUpIfDecimalGreaterFive[uint16](float64(damage * correctionValue))
	}
	return result
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
