package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type UpdatePartyCommand struct {
	factory.PartyInput
}

func NewUpdatePartyCommand(id uint, name string, battleFormat string, isPrivate bool, userId uint,
	partyResultIds []uint, partyTagIds []uint, trainedPokemonIds []uint) UpdatePartyCommand {
	return UpdatePartyCommand{
		factory.NewPartyInput(id, name, battleFormat, isPrivate, 0, partyResultIds, partyTagIds, trainedPokemonIds),
	}
}
