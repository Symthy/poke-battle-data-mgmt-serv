package pokemons

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.PokemonId, uint16] = (*Pokemon)(nil)

// Todo: add fields
type Pokemon struct {
	id               identifier.PokemonId
	pokedexId        value.PokedexId
	formName         string
	name             string
	englishName      string
	generation       uint16
	typeSet          *value.PokemonTypeSet
	abilitySet       *value.PokemonAbilityIdSet
	baseStatsH       value.PokemonBaseStats
	baseStatsA       value.PokemonBaseStats
	baseStatsB       value.PokemonBaseStats
	baseStatsC       value.PokemonBaseStats
	baseStatsD       value.PokemonBaseStats
	baseStatsS       value.PokemonBaseStats
	isFinalEvolution bool
}

func NewPokemon(
	id identifier.PokemonId,
	pokedexId value.PokedexId,
	formName string,
	name string,
	englishName string,
	generation uint64,
	typeSet *value.PokemonTypeSet,
	abilitySet *value.PokemonAbilityIdSet,
	baseStatsH value.PokemonBaseStats,
	baseStatsA value.PokemonBaseStats,
	baseStatsB value.PokemonBaseStats,
	baseStatsC value.PokemonBaseStats,
	baseStatsD value.PokemonBaseStats,
	baseStatsS value.PokemonBaseStats,
	IsFinalEvolution bool) Pokemon {
	return Pokemon{
		id:               id,
		pokedexId:        pokedexId,
		formName:         formName,
		name:             name,
		englishName:      englishName,
		generation:       uint16(generation),
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

func (p Pokemon) TypeSet() value.PokemonTypeSet {
	return p.typeSet
}

func (p Pokemon) ResolveActualValues(effortValues *value.EffortValues) *value.PokemonActualValues {
	actualValues := value.NewPokemonActualValues(
		calculateActualValueH(p.baseStatsH, 31, effortValues.H()),
		calculateActualValueABCDS(p.baseStatsA, 31, effortValues.A()),
		calculateActualValueABCDS(p.baseStatsB, 31, effortValues.B()),
		calculateActualValueABCDS(p.baseStatsC, 31, effortValues.C()),
		calculateActualValueABCDS(p.baseStatsD, 31, effortValues.D()),
		calculateActualValueABCDS(p.baseStatsS, 31, effortValues.S()),
	)
	return actualValues
}

func calculateActualValueH(baseStats value.PokemonBaseStats, individualValue uint8, effortValue value.EffortValue) uint16 {
	level := uint16(50)
	b := uint16(baseStats.Value() * 2)
	i := uint16(individualValue)
	e := uint16(effortValue.Value() / 4)
	l := float64(level) / 100.0
	a := float64(b+i+e) * l
	return uint16(a) + level + 10
}

func calculateActualValueABCDS(baseStats value.PokemonBaseStats, individualValue uint8, effortValue value.EffortValue) uint16 {
	level := uint16(50)
	b := uint16(baseStats.Value() * 2)
	i := uint16(individualValue)
	e := uint16(effortValue.Value() / 4)
	l := float64(level) / 100.0
	a := float64(b+i+e) * l
	return uint16(a) + 5
}

func (p Pokemon) Notify(note IPokemonNotification) {
	note.SetId(p.id)
	note.SetPokedexId(p.pokedexId)
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
