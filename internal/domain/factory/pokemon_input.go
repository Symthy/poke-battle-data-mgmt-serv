package factory

import (
	"github.com/Symthy/PokeRest/internal/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type PokemonInput struct {
	id               uint64
	pokedexNo        uint64
	formNo           uint64
	formName         string
	name             string
	englishName      string
	generation       uint64
	type1            string
	type2            string
	abilityId1       uint64
	abilityId2       uint64
	hiddenAbilityId  uint64
	baseStatsH       uint64
	baseStatsA       uint64
	baseStatsB       uint64
	baseStatsC       uint64
	baseStatsD       uint64
	baseStatsS       uint64
	isFinalEvolution bool
}

func NewPokemonInput(
	id, pokedexNo, formNo uint64, formName, name, englishName string, generation uint64, type1, type2 string,
	abilityId1, abilityId2, hiddenAbilityId, baseStatsH, baseStatsA, baseStatsB, baseStatsC, baseStatsD, baseStatsS uint64,
	isFinalEvolution bool) PokemonInput {
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

func NewPokemonBuilder() *PokemonInput {
	return &PokemonInput{}
}

func (p *PokemonInput) Id(id uint16) {
	p.id = uint64(id)
}
func (p *PokemonInput) PokedexNo(pokedexNo uint16) {
	p.pokedexNo = uint64(pokedexNo)
}
func (p *PokemonInput) FormNo(formNo uint16) {
	p.formNo = uint64(formNo)
}
func (p *PokemonInput) FormName(formName string) {
	p.formName = formName
}
func (p *PokemonInput) Name(name string) {
	p.name = name
}
func (p *PokemonInput) EnglishName(englishName string) {
	p.englishName = englishName
}
func (p *PokemonInput) Generation(generation uint16) {
	p.generation = uint64(generation)
}
func (p *PokemonInput) TypeOne(type1 string) {
	p.type1 = type1
}
func (p *PokemonInput) TypeTwo(type2 string) {
	p.type2 = type2
}
func (p *PokemonInput) AbilityIdOne(abilityId1 uint16) {
	p.abilityId1 = uint64(abilityId1)
}
func (p *PokemonInput) AbilityIdTwo(abilityId2 uint16) {
	p.abilityId2 = uint64(abilityId2)
}
func (p *PokemonInput) HiddenAbilityId(hiddenAbilityId uint16) {
	p.hiddenAbilityId = uint64(hiddenAbilityId)
}
func (p *PokemonInput) BaseStatsH(baseStatsH uint8) {
	p.baseStatsH = uint64(baseStatsH)
}
func (p *PokemonInput) BaseStatsA(baseStatsA uint8) {
	p.baseStatsA = uint64(baseStatsA)
}
func (p *PokemonInput) BaseStatsB(baseStatsB uint8) {
	p.baseStatsB = uint64(baseStatsB)
}
func (p *PokemonInput) BaseStatsC(baseStatsC uint8) {
	p.baseStatsC = uint64(baseStatsC)
}
func (p *PokemonInput) BaseStatsD(baseStatsD uint8) {
	p.baseStatsD = uint64(baseStatsD)
}
func (p *PokemonInput) BaseStatsS(baseStatsS uint8) {
	p.baseStatsS = uint64(baseStatsS)
}
func (p *PokemonInput) SetIsFinalEvolution(isFinalEvolution bool) {
	p.isFinalEvolution = isFinalEvolution
}

func (i PokemonInput) BuildDomain() (*pokemons.Pokemon, error) {
	id, err := identifier.NewPokemonId(i.id)
	if err != nil {
		return nil, err
	}
	pokedexId, err := value.NewPokedexNumber(i.pokedexNo, i.formNo)
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
	domain := pokemons.NewPokemon(*id, *pokedexId, i.formName, i.name, i.englishName,
		i.generation, typeSet, abilitySet, baseStatsH, baseStatsA, baseStatsB, baseStatsC, baseStatsD,
		baseStatsS, i.isFinalEvolution)
	return &domain, nil
}
