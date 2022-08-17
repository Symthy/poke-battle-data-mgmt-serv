package command

import (
	"github.com/Symthy/PokeRest/internal/domain/factory"
)

type CreatePartyCommand struct {
	factory.PartyInput
}

func NewCreatePartyCommand(
	name, battleFormat string, isPrivate bool, userId uint64, partyTagIds, trainedPokemonIds []uint64,
) CreatePartyCommand {
	return CreatePartyCommand{
		factory.NewPartyInput(0, name, battleFormat, isPrivate, userId, nil, partyTagIds, trainedPokemonIds),
	}
}
