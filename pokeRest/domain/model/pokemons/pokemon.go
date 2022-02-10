package pokemons

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

// Todo: add fields
type Pokemon struct {
	id                 uint
	pokedexNo          int
	formNo             int
	formName           string
	name               string
	englishName        string
	generation         int
	typePrimary        value.PokemonType
	typeSecondary      value.PokemonType
	abilityIdPrimary   value.OptionalAbilityId
	abilityIdSecondary value.OptionalAbilityId
	hiddenAbilityId    value.OptionalAbilityId
	baseStatsH         int
	baseStatsA         int
	baseStatsB         int
	baseStatsC         int
	baseStatsD         int
	baseStatsS         int
	isFinalEvolution   bool
}

// Todo: builder
func NewPokemon(
	id uint,
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
	baseStatsH int,
	baseStatsA int,
	baseStatsB int,
	baseStatsC int,
	baseStatsD int,
	baseStatsS int,
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
		abilityIdPrimary:   value.NewOptionalAbilityId(abilityId1), // always exist.
		abilityIdSecondary: value.NewOptionalAbilityId(abilityId2),
		hiddenAbilityId:    value.NewOptionalAbilityId(hiddenAbilityId),
		baseStatsH:         baseStatsH,
		baseStatsA:         baseStatsA,
		baseStatsB:         baseStatsB,
		baseStatsC:         baseStatsC,
		baseStatsD:         baseStatsD,
		baseStatsS:         baseStatsS,
		isFinalEvolution:   IsFinalEvolution,
	}
}

func NewPokemonNonId(
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
	baseStatsH int,
	baseStatsA int,
	baseStatsB int,
	baseStatsC int,
	baseStatsD int,
	baseStatsS int,
	IsFinalEvolution bool) Pokemon {
	return Pokemon{
		id:                 0,
		pokedexNo:          pokedexNo,
		formNo:             formNo,
		formName:           formName,
		name:               name,
		englishName:        englishName,
		generation:         generation,
		typePrimary:        value.NewPokemonType(type1),
		typeSecondary:      value.NewPokemonType(type2),
		abilityIdPrimary:   value.NewOptionalAbilityId(abilityId1), // always exist.
		abilityIdSecondary: value.NewOptionalAbilityId(abilityId2),
		hiddenAbilityId:    value.NewOptionalAbilityId(hiddenAbilityId),
		baseStatsH:         baseStatsH,
		baseStatsA:         baseStatsA,
		baseStatsB:         baseStatsB,
		baseStatsC:         baseStatsC,
		baseStatsD:         baseStatsD,
		baseStatsS:         baseStatsS,
		isFinalEvolution:   IsFinalEvolution,
	}
}

func (p Pokemon) Id() uint {
	return p.id
}

func (p Pokemon) PokedexNo() int {
	return p.pokedexNo
}

func (p Pokemon) FormNo() int {
	return p.formNo
}

func (p Pokemon) FormName() string {
	return p.formName
}

func (p Pokemon) Name() string {
	return p.name
}

func (p Pokemon) EnglishName() string {
	return p.englishName
}

func (p Pokemon) Generation() int {
	return p.generation
}

func (p Pokemon) TypePrimary() value.PokemonType {
	return p.typePrimary
}

func (p Pokemon) TypeSecondary() value.PokemonType {
	return p.typeSecondary
}

func (p Pokemon) AbilityIdPrimary() value.OptionalAbilityId {
	return p.abilityIdPrimary
}

func (p Pokemon) AbilityIdSecondary() value.OptionalAbilityId {
	return p.abilityIdSecondary
}

func (p Pokemon) HiddenAbilityId() value.OptionalAbilityId {
	return p.hiddenAbilityId
}

func (p Pokemon) BaseStatsH() int {
	return p.baseStatsH
}

func (p Pokemon) BaseStatsA() int {
	return p.baseStatsA
}

func (p Pokemon) BaseStatsB() int {
	return p.baseStatsB
}

func (p Pokemon) BaseStatsC() int {
	return p.baseStatsC
}

func (p Pokemon) BaseStatsD() int {
	return p.baseStatsD
}

func (p Pokemon) BaseStatsS() int {
	return p.baseStatsS
}

func (p Pokemon) IsFinalEvolution() bool {
	return p.isFinalEvolution
}