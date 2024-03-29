package command

import "github.com/Symthy/PokeRest/internal/domain/factory"

type DeleteTrainedPokemonCommand struct {
	factory.TrainedPokemonInput
}

func NewDeleteTrainedPokemonCommand(id, userId uint64) DeleteTrainedPokemonCommand {
	builder := factory.NewTrainedPokemonBuilder()
	builder.Id(id)
	builder.UserId(userId)
	return DeleteTrainedPokemonCommand{builder}
}
