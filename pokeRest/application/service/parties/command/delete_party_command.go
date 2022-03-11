package command

import "github.com/Symthy/PokeRest/pokeRest/domain/factory"

type DeletePartyCommand struct {
	factory.PartyInput
}

func NewDeletePartyCommand(id uint, userId uint) DeletePartyCommand {
	builder := factory.NewPartyBuilder()
	builder.Id(id)
	builder.UserId(userId)
	return DeletePartyCommand{builder}
}
