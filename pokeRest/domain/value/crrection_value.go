package value

type CorrectionValue struct {
	correctionType    CorrectionType
	targetPokemonType PokemonType
	rate              float32
}

type CorrectionType string

const (
	// 技威力
	PhysicalMovePowerRate  = "physicalMovePowerRate"
	SpecialMovePowerRate   = "specialMovePowerRate"
	AttackRate             = "attackRate"
	SpecialAttackRate      = "specialAttackRate"
	AttackPowerRate        = "attackPowerRate"
	SpecialAttackPowerRate = "specialAttackPowerRate"
	DeffenceRate           = "deffenceRate"
	SpecialDeffenceRate    = "specialDeffenceRate"
	DamageRate             = "damageRate"
	WeightRate             = "weightRate"
)
