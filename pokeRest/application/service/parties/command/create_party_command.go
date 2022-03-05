package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type CreatePartyCommand struct {
	factory.PartyInput
}

func NewCreatePartyCommand(name string, battleFormat string, isPrivate bool, userId uint, partyTagIds []uint, trainedPokemonIds []uint) CreatePartyCommand {
	return CreatePartyCommand{
		factory.NewPartyInput(0, name, battleFormat, isPrivate, userId, nil, partyTagIds, trainedPokemonIds),
	}
}
