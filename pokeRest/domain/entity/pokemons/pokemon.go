package pokemons

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.PokemonId] = (*Pokemon)(nil)

// Todo: add fields
type Pokemon struct {
	id               identifier.PokemonId
	pokedexNo        int
	formNo           int
	formName         string
	name             string
	englishName      string
	generation       int
	typeSet          value.PokemonTypeSet
	abilitySet       value.PokemonAbilityIdSet
	baseStatsH       value.PokemonBaseStats
	baseStatsA       value.PokemonBaseStats
	baseStatsB       value.PokemonBaseStats
	baseStatsC       value.PokemonBaseStats
	baseStatsD       value.PokemonBaseStats
	baseStatsS       value.PokemonBaseStats
	isFinalEvolution bool
}

// Todo: factory
func NewPokemon(
	id identifier.PokemonId,
	pokedexNo int,
	formNo int,
	formName string,
	name string,
	englishName string,
	generation int,
	typeSet value.PokemonTypeSet,
	abilitySet value.PokemonAbilityIdSet,
	baseStatsH value.PokemonBaseStats,
	baseStatsA value.PokemonBaseStats,
	baseStatsB value.PokemonBaseStats,
	baseStatsC value.PokemonBaseStats,
	baseStatsD value.PokemonBaseStats,
	baseStatsS value.PokemonBaseStats,
	IsFinalEvolution bool) Pokemon {
	return Pokemon{
		id:               id,
		pokedexNo:        pokedexNo,
		formNo:           formNo,
		formName:         formName,
		name:             name,
		englishName:      englishName,
		generation:       generation,
		typeSet:          typeSet,
		abilitySet:       abilitySet,
		baseStatsH:       baseStatsH,
		baseStatsA:       baseStatsA,
		baseStatsB:       baseStatsB,
		baseStatsC:       baseStatsC,
		baseStatsD:       baseStatsD,
		baseStatsS:       baseStatsS,
		isFinalEvolution: IsFinalEvolution,
	}
}

func (p Pokemon) Id() identifier.PokemonId {
	return p.id
}

func (p Pokemon) PokedexNo() int {
	return p.pokedexNo
}

func (p Pokemon) FormNo() int {
	return p.formNo
}

func (p Pokemon) Notify(note IPokemonNotification) {
	note.SetId(p.id)
	note.SetPokedexNo(p.pokedexNo)
	note.SetFormNo(p.formNo)
	note.SetFormName(p.formName)
	note.SetName(p.name)
	note.SetEnglishName(p.englishName)
	note.SetGeneration(p.generation)
	note.SetTypeSet(p.typeSet)
	note.SetAbilitySet(p.abilitySet)
	note.SetBaseStatsH(p.baseStatsH)
	note.SetBaseStatsA(p.baseStatsA)
	note.SetBaseStatsB(p.baseStatsB)
	note.SetBaseStatsC(p.baseStatsC)
	note.SetBaseStatsD(p.baseStatsD)
	note.SetBaseStatsS(p.baseStatsS)
	note.SetIsFinalEvolution(p.isFinalEvolution)
}
