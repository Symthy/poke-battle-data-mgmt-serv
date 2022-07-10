package damages

type IDamageClacElements interface {
	IsNoDamage() bool
	FinalAttackValue() uint16
	AttackActualValue() uint16
	DefenseActualValue() uint16
	DamageCorrectedValue() uint16
	WeatherCorrectedValue() uint16
	IsTypeMatchAttackSide() bool
	IsBurnAttackSide() bool
	AbnormalStateAttackSideCorectedValue() uint16
	TypeCompatibilityDamageRate() float64
}
