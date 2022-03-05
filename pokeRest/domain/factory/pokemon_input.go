package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type PokemonInput struct {
	id               uint
	pokedexNo        int
	formNo           int
	formName         string
	name             string
	englishName      string
	generation       int
	type1            string
	type2            string
	abilityId1       uint
	abilityId2       uint
	hiddenAbilityId  uint
	baseStatsH       int
	baseStatsA       int
	baseStatsB       int
	baseStatsC       int
	baseStatsD       int
	baseStatsS       int
	isFinalEvolution bool
}

func NewPokemonInput(
	id uint, pokedexNo int, formNo int, formName string, name string, englishName string,
	generation int, type1 string, type2 string, abilityId1 uint, abilityId2 uint, hiddenAbilityId uint,
	baseStatsH int, baseStatsA int, baseStatsB int, baseStatsC int, baseStatsD int, baseStatsS int,
	isFinalEvolution bool,
) PokemonInput {
	return PokemonInput{
		id:               id,
		pokedexNo:        pokedexNo,
		formNo:           formNo,
		formName:         formName,
		name:             name,
		englishName:      englishName,
		generation:       generation,
		type1:            type1,
		type2:            type2,
		abilityId1:       abilityId1,
		abilityId2:       abilityId2,
		hiddenAbilityId:  hiddenAbilityId,
		baseStatsH:       baseStatsH,
		baseStatsA:       baseStatsA,
		baseStatsB:       baseStatsB,
		baseStatsC:       baseStatsC,
		baseStatsD:       baseStatsD,
		baseStatsS:       baseStatsS,
		isFinalEvolution: isFinalEvolution,
	}
}

func (i PokemonInput) BuildDomain() (*pokemons.Pokemon, error) {
	id, err := identifier.NewPokemonId(i.id)
	if err != nil {
		return nil, err
	}
	typeSet := value.NewPokemonTypeSet(value.NewPokemonType(i.type1), value.NewPokemonType(i.type2))
	abilityId1, err := identifier.NewAbilityId(i.abilityId1)
	if err != nil {
		return nil, err
	}
	abilityId2, err := identifier.NewAbilityId(i.abilityId2)
	if err != nil {
		return nil, err
	}
	hiddenAbilityId, err := identifier.NewAbilityId(i.hiddenAbilityId)
	if err != nil {
		return nil, err
	}
	abilitySet := value.NewPokemonAbilityIdSet(*abilityId1, *abilityId2, *hiddenAbilityId)
	baseStatsH := value.NewPokemonBaseStats(i.baseStatsH)
	baseStatsA := value.NewPokemonBaseStats(i.baseStatsA)
	baseStatsB := value.NewPokemonBaseStats(i.baseStatsB)
	baseStatsC := value.NewPokemonBaseStats(i.baseStatsC)
	baseStatsD := value.NewPokemonBaseStats(i.baseStatsD)
	baseStatsS := value.NewPokemonBaseStats(i.baseStatsS)
	domain := pokemons.NewPokemon(*id, i.pokedexNo, i.formNo, i.formName, i.name, i.englishName,
		i.generation, typeSet, abilitySet, baseStatsH, baseStatsA, baseStatsB, baseStatsC, baseStatsD,
		baseStatsS, i.isFinalEvolution)
	return &domain, nil
}
