package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type BattlePokemonAdjustment struct {
	pokemonId    uint
	abilityId    uint
	heldItemId   uint
	nature       value.PokemonNature
	effortValueH value.EffortValue
	effortValueA value.EffortValue
	effortValueB value.EffortValue
	effortValueC value.EffortValue
	effortValueD value.EffortValue
}
