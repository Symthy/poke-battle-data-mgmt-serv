package battles

type OverrideTarget string

const (
	// ダメージ
	OverrideFixedDamage OverrideTarget = "FixedDamage"
	// タイプ一致補正
	OverrideTypeMatchCorrection OverrideTarget = "TypeMatch"
	// タイプ
	OverridePokemonType OverrideTarget = "PokemonType"
)
