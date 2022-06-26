package damages

type IDamageClacElements interface {
	IsNoDamage() bool
	FinalAttackValue() uint16
	AttackActualValue() uint16
	DefenseActualValue() uint16
	DamageCorrectedValue() uint16
	WeatherCorrectedValue() uint16
	IsTypeMatch() bool
	TypeCompatibilityDamageRate() float64
}
