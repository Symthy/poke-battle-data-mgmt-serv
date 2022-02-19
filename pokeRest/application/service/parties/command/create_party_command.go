package command

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"

type CreatePartyCommand struct {
	parties.Party
}

func NewCreatePartyCommand(name string, battleFormat string, isPrivate bool, userId uint, partyTagIds []uint, trainedPokemonIds []uint) CreatePartyCommand {
	cmd := CreatePartyCommand{
		parties.NewPartyOfUnregistered(name, battleFormat, isPrivate, userId, partyTagIds, trainedPokemonIds),
	}
	return cmd
}

func (cmd CreatePartyCommand) ToDomain() parties.Party {
	return cmd.Party
}
