package fmath

import "math"

func RoundDown(value float64) float64 {
	// 小数点切り捨て
	return math.Trunc(value)
}

func RoundUp(value float64) float64 {
	// 小数点切り上げ
	return math.Ceil(value)
}

func Round(value float64) int {
	return int(math.Round(value))
}

// 五捨五超入
func RoundUpIfDecimalGreaterFive(value float64) int {
	integerValue := RoundDown(value)
	if value-integerValue > 0.5 {
		return int(RoundUp(value))
	}
	return int(RoundDown(value))
}
