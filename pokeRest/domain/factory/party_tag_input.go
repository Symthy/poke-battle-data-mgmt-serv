package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type PartyTagInput struct {
	id           uint64
	name         string
	isGeneration bool
	isSeason     bool
}

func NewPartyTagInput(id uint64, name string, isGeneration, isSeason bool) PartyTagInput {
	return PartyTagInput{id, name, isGeneration, isSeason}
}

func NewPartyTagBuilder() *PartyTagInput {
	return &PartyTagInput{}
}

func (i *PartyTagInput) Id(id uint64) {
	i.id = id
}
func (i *PartyTagInput) Name(name string) {
	i.name = name
}
func (i *PartyTagInput) SetIsGeneration(isGeneration bool) {
	i.isGeneration = isGeneration
}
func (i *PartyTagInput) SetIsSeason(isSeason bool) {
	i.isSeason = isSeason
}

func (i PartyTagInput) BuildDomain() (*parties.PartyTag, error) {
	id, err := identifier.NewPartyTagId(i.id)
	if err != nil {
		return nil, err
	}
	domain := parties.NewPartyTag(*id, i.name, i.isGeneration, i.isSeason)
	return &domain, nil
}
