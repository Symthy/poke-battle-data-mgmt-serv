package model

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type Pokemon struct {
	id                 int
	pokedexNo          int
	formNo             int
	formName           string
	name               string
	englishName        string
	generation         int
	typePrimary        value.PokemonType
	typeSecondary      value.PokemonType
	abilityIdPrimary   int
	abilityIdSecondary value.OptionaAbilitylId
	hiddenAbilityId    value.OptionaAbilitylId
	IsFinalEvolution   bool
}

func NewPokemon(
	id int,
	pokedexNo int,
	formNo int,
	formName string,
	name string,
	englishName string,
	generation int,
	type1 string,
	type2 string,
	abilityId1 *int,
	abilityId2 *int,
	hiddenAbilityId *int,
	IsFinalEvolution bool) Pokemon {
	return Pokemon{
		id:                 id,
		pokedexNo:          pokedexNo,
		formNo:             formNo,
		formName:           formName,
		name:               name,
		englishName:        englishName,
		generation:         generation,
		typePrimary:        value.NewPokemonType(type1),
		typeSecondary:      value.NewPokemonType(type2),
		abilityIdPrimary:   *abilityId1, // always exist.
		abilityIdSecondary: value.NewOptionalAbilityId(abilityId2),
		hiddenAbilityId:    value.NewOptionalAbilityId(hiddenAbilityId),
		IsFinalEvolution:   IsFinalEvolution,
	}
}

func (p Pokemon) Id() int {
	return p.id
}
