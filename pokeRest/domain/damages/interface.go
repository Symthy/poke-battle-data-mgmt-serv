package damages

type IDamageClacElements interface {
	IsNoDamage() bool
	FinalAttackValue() uint16
	AttackActualValue() uint16
	DefenseActualValue() uint16
	DamageCorrectedValue() uint16
	ApplyWeatherCorrection(damage uint16) uint16
	IsTypeMatch() bool
	TypeCompatibilityDamageRate() float64
}
