package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type IPokemonBattleDataSet interface {
	AttackPokemonTypeOfFirst() value.PokemonType
	AttackPokemonTypeOfSecond() value.PokemonType
	AttackPokemonActualValueS(*StatusCorrections) uint16
	DefensePokemonTypeOfFirst() value.PokemonType
	DefensePokemonTypeOfSecond() value.PokemonType
	DefensePokemonActualValueS(*StatusCorrections) uint16
	MovePokemonType() value.PokemonType
	HasItemAttackSide() bool
	HasItemDefenseSide() bool
	TypeCompatibilityDamageRate() float32
}
