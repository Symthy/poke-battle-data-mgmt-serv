package command

import "github.com/Symthy/PokeRest/pokeRest/domain/factory"

type DeleteTrainedPokemonCommand struct {
	factory.TrainedPokemonInput
}

func NewDeleteTrainedPokemonCommand(id uint, userId uint) DeleteTrainedPokemonCommand {
	builder := factory.NewTrainedPokemonBuilder()
	builder.Id(id)
	builder.UserId(userId)
	return DeleteTrainedPokemonCommand{builder}
}
