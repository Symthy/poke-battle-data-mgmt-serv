package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
)

type CreateTrainedPokemonCommand struct {
	trainings.TrainedPokemon
}

// Todo
func NewCreateTrainedPokemonCommand() CreateTrainedPokemonCommand {
	cmd := CreateTrainedPokemonCommand{
		// trainings.NewTrainedPokemonOfUnregistered(),
	}
	return cmd
}

func (c CreateTrainedPokemonCommand) ToDomain() trainings.TrainedPokemon {
	return c.TrainedPokemon
}
