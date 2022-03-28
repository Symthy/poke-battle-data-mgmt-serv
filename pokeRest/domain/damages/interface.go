package damages

type IPokemonBattleDataSet interface {
	AttackPokemonTypeOfFirst() string
	AttackPokemonTypeOfSecond() string
	AttackPokemonActualValueS() string
	DefencePokemonTypeOfFirst() string
	DefencePokemonTypeOfSecond() string
	DefencePokemonActualValueS() string
	MovePokemonType() string
	HasItemAttackSide() bool
	HasItemDefenceSide() bool
}
