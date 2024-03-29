package fmath

import "math"

func RoundDown[T uint16 | uint32 | uint64](value float64) T {
	// 小数点切り捨て
	return T(math.Trunc(value))
}

func RoundUp[T uint16 | uint32 | uint64](value float64) T {
	// 小数点切り上げ
	return T(math.Ceil(value))
}

func Round[T uint16 | uint32 | uint64](value float64) T {
	// 四捨五入
	return T(math.Round(value))
}

func RoundUpIfDecimalGreaterFive[T uint16 | uint32 | uint64](value float64) T {
	// 五捨五超入
	integerValue := float64(RoundDown[T](value))
	if value-integerValue > 0.5 {
		return RoundUp[T](value)
	}
	return RoundDown[T](value)
}
