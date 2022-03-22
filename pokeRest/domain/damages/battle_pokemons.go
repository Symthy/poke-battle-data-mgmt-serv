package damages

type BattlePokemons struct {
	attackSide              BattleAttackPokemon
	deffenceSide            BattleDefencePokemon
	compatibilityDamageRate float32
}

func NewBattlePokemons(
	attackSide BattleAttackPokemon, deffenceSide BattleDefencePokemon, compatibilityDamageRate float32,
) BattlePokemons {
	return BattlePokemons{
		attackSide:              attackSide,
		deffenceSide:            deffenceSide,
		compatibilityDamageRate: compatibilityDamageRate,
	}
}

func (b BattlePokemons) IsNoDamage() bool {
	return int(b.compatibilityDamageRate) == 0
}
