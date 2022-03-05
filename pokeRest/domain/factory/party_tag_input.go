package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type PartyTagInput struct {
	id           uint
	name         string
	isGeneration bool
	isSeason     bool
}

func NewPartyTagInput(id uint, name string, isGeneration bool, isSeason bool) PartyTagInput {
	return PartyTagInput{id, name, isGeneration, isSeason}
}

func (i PartyTagInput) BuildDomain() (*parties.PartyTag, error) {
	id, err := identifier.NewPartyTagId(i.id)
	if err != nil {
		return nil, err
	}
	domain := parties.NewPartyTag(*id, i.name, i.isGeneration, i.isSeason)
	return &domain, nil
}
