package command

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"

type UpdatePartyCommand struct {
	parties.Party
}

func NewUpdatePartyCommand(id uint, name string, battleFormat string, isPrivate bool,
	partyResultIds []uint, partyTagIds []uint, trainedPokemonIds []uint) UpdatePartyCommand {
	cmd := UpdatePartyCommand{
		parties.NewPartyForUpdated(id, name, battleFormat, isPrivate, partyResultIds, partyTagIds, trainedPokemonIds),
	}
	return cmd
}

func (cmd UpdatePartyCommand) ToDomain() parties.Party {
	return cmd.Party
}
