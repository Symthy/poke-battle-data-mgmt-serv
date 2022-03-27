package damages

var _ IPokemonBattleDataSet = (*PokemonBattleDataSet)(nil)

type PokemonBattleDataSet struct {
	AttackSidePokemon
	AttackSideBattleEffects
	DefenceSidePokemon
	DefenceSideBattleEffects
	AttackMove
	typeCompatibilityDamageRate float32
}

func NewPokemonBattleDataSet(
	attackSide AttackSidePokemon, attackEffects AttackSideBattleEffects,
	deffenceSide DefenceSidePokemon, defenceEffects DefenceSideBattleEffects,
	attackMove AttackMove, compatibilityDamageRate float32,
) PokemonBattleDataSet {
	return PokemonBattleDataSet{
		AttackSidePokemon:           attackSide,
		AttackSideBattleEffects:     attackEffects,
		DefenceSidePokemon:          deffenceSide,
		DefenceSideBattleEffects:    defenceEffects,
		AttackMove:                  attackMove,
		typeCompatibilityDamageRate: compatibilityDamageRate,
	}
}

func (b PokemonBattleDataSet) IsNoDamage() bool {
	return int(b.typeCompatibilityDamageRate) == 0
}

func (p PokemonBattleDataSet) ResolvePowerValue() int {
	applier := p.AttackSideBattleEffects.SupplyPowerCorrectionApplier(p.AttackMove.getSpecies())
	return int(applier(float32(p.getPower())))
}
