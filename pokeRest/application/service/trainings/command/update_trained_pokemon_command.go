package command

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"

type UpdateTrainedPokemonCommand struct {
	trainings.TrainedPokemon
}

// Todo
func NewUpdateTrainedPokemonCommand() UpdateTrainedPokemonCommand {
	cmd := UpdateTrainedPokemonCommand{
		// trainings.NewTrainedPokemonForUpdate(),
	}
	return cmd
}

func (c UpdateTrainedPokemonCommand) ToDomain() trainings.TrainedPokemon {
	return c.TrainedPokemon
}
