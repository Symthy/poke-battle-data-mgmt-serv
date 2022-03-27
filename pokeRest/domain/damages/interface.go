package damages

type IPokemonBattleDataSet interface {
	AttackPokemonTypeOfFirst() string
	AttackPokemonTypeOfSecond() string
	AttackPokemonActualValueS() string
	DefencePokemonActualValueS() string
	MovePokemonType() string
	HasItemAttackSide() bool
	HasItemDefenceSide() bool
}
