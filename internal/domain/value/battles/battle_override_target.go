package battles

type OverrideTarget string

func (o OverrideTarget) ToString() string {
	return string(o)
}

const (
	// ダメージ
	OverrideFixedDamage OverrideTarget = "FixedDamage"
	// タイプ一致補正
	OverrideTypeMatchCorrection OverrideTarget = "TypeMatch"
	// タイプ
	OverridePokemonType OverrideTarget = "PokemonType"
)
