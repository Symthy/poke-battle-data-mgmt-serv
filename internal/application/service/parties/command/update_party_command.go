package command

import (
	"github.com/Symthy/PokeRest/internal/domain/factory"
)

type UpdatePartyCommand struct {
	factory.PartyInput
}

func NewUpdatePartyCommand(id uint64, name, battleFormat string, isPrivate bool, userId uint64,
	partyResultIds, partyTagIds, trainedPokemonIds []uint64) UpdatePartyCommand {
	return UpdatePartyCommand{
		factory.NewPartyInput(id, name, battleFormat, isPrivate, 0, partyResultIds, partyTagIds, trainedPokemonIds),
	}
}
