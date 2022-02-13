package command

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"

type CreatePartyCommand struct {
	parties.Party
}

func NewCreatePartyCommand(name string, battleFormat string, partyTagIds []uint,
	trainedPokemonIds []uint, isPrivate bool, userId uint) CreatePartyCommand {
	cmd := CreatePartyCommand{
		parties.NewPartyOfUnregistered(name, battleFormat, partyTagIds, trainedPokemonIds, isPrivate, userId),
	}
	return cmd
}

func (cmd CreatePartyCommand) ToDomain() parties.Party {
	return cmd.Party
}
