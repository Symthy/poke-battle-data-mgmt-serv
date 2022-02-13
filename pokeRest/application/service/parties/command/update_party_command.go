package command

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"

type UpdatePartyCommand struct {
	parties.Party
}

func NewUpdatePartyCommand(id uint, name string, battleFormat string, partyTagIds []uint,
	trainedPokemonIds []uint, isPrivate bool) UpdatePartyCommand {
	cmd := UpdatePartyCommand{
		parties.NewPartyForUpdated(id, name, battleFormat, partyTagIds, trainedPokemonIds, isPrivate),
	}
	return cmd
}

func (cmd UpdatePartyCommand) ToDomain() parties.Party {
	return cmd.Party
}
