package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type IPokemonBattleDataSet interface {
	AttackPokemonTypeOfFirst() value.PokemonType
	AttackPokemonTypeOfSecond() value.PokemonType
	AttackPokemonActualValueS() int
	DefensePokemonTypeOfFirst() value.PokemonType
	DefensePokemonTypeOfSecond() value.PokemonType
	DefensePokemonActualValueS() int
	MovePokemonType() value.PokemonType
	HasItemAttackSide() bool
	HasItemDefenseSide() bool
}
