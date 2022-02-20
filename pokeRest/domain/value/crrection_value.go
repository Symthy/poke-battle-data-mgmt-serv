package value

type CorrectionValue struct {
	correctionType    CorrectionType
	targetPokemonType *PokemonType
	rate              float32
}

func isInvalidValue(rate float32) bool {
	return rate == 1 || rate <= 0
}

func PhysicalMoveCorrectionValue(targetPokemonType *PokemonType, rate float32) *CorrectionValue {
	if isInvalidValue(rate) {
		return nil
	}
	return &CorrectionValue{
		correctionType:    PhysicalMoveCorrection,
		targetPokemonType: targetPokemonType,
		rate:              rate,
	}
}

func SpecialMoveCorrectionValue(targetPokemonType *PokemonType, rate float32) *CorrectionValue {
	if isInvalidValue(rate) {
		return nil
	}
	return &CorrectionValue{
		correctionType:    SpecialMoveCorrection,
		targetPokemonType: targetPokemonType,
		rate:              rate,
	}
}

func AttackCorrectionValue(targetPokemonType *PokemonType, rate float32) *CorrectionValue {
	if isInvalidValue(rate) {
		return nil
	}
	return &CorrectionValue{
		correctionType:    AttackCorrection,
		targetPokemonType: targetPokemonType,
		rate:              rate,
	}
}

func SpecialAttackCorrectionValue(targetPokemonType *PokemonType, rate float32) *CorrectionValue {
	if isInvalidValue(rate) {
		return nil
	}
	return &CorrectionValue{
		correctionType:    SpecialAttackCorrection,
		targetPokemonType: targetPokemonType,
		rate:              rate,
	}
}

func AttackPowerCorrectionValue(targetPokemonType *PokemonType, rate float32) *CorrectionValue {
	if isInvalidValue(rate) {
		return nil
	}
	return &CorrectionValue{
		correctionType:    AttackPowerCorrection,
		targetPokemonType: targetPokemonType,
		rate:              rate,
	}
}

func SpecialAttackPowerCorrectionValue(targetPokemonType *PokemonType, rate float32) *CorrectionValue {
	if isInvalidValue(rate) {
		return nil
	}
	return &CorrectionValue{
		correctionType:    SpecialAttackPowerCorrection,
		targetPokemonType: targetPokemonType,
		rate:              rate,
	}
}

func DeffenceCorrectionValue(targetPokemonType *PokemonType, rate float32) *CorrectionValue {
	if isInvalidValue(rate) {
		return nil
	}
	return &CorrectionValue{
		correctionType:    DeffenceCorrection,
		targetPokemonType: targetPokemonType,
		rate:              rate,
	}
}

func SpecialDeffenceCorrectionValue(targetPokemonType *PokemonType, rate float32) *CorrectionValue {
	if isInvalidValue(rate) {
		return nil
	}
	return &CorrectionValue{
		correctionType:    SpecialDeffenceCorrection,
		targetPokemonType: targetPokemonType,
		rate:              rate,
	}
}

func DamageCorrectionValue(targetPokemonType PokemonType, rate float32) *CorrectionValue {
	if isInvalidValue(rate) {
		return nil
	}
	return &CorrectionValue{
		correctionType:    DamageCorrection,
		targetPokemonType: &targetPokemonType,
		rate:              rate,
	}
}

func WeightCorrectionValue(targetPokemonType *PokemonType, rate float32) *CorrectionValue {
	if isInvalidValue(rate) {
		return nil
	}
	return &CorrectionValue{
		correctionType:    WeightCorrection,
		targetPokemonType: targetPokemonType,
		rate:              rate,
	}
}

type CorrectionType string

const (
	// 技威力
	PhysicalMoveCorrection = "physicalMove"
	SpecialMoveCorrection  = "specialMove"
	// 攻撃力
	AttackCorrection        = "attack"
	SpecialAttackCorrection = "specialAttack"
	// 攻撃威力
	AttackPowerCorrection        = "attackPower"
	SpecialAttackPowerCorrection = "specialAttackPower"
	// 防御力
	DeffenceCorrection        = "deffence"
	SpecialDeffenceCorrection = "specialDeffence"
	// その他
	DamageCorrection = "damage"
	WeightCorrection = "weight"
)
