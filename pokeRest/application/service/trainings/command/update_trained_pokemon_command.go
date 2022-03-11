package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type UpdateTrainedPokemonCommand struct {
	factory.TrainedPokemonInput
}

// Todo
func NewUpdateTrainedPokemonCommand(id uint, gender, nickname, description string, isPrivate bool,
	userId uint, adjustment factory.TrainedPokemonAdjustmentInput) UpdateTrainedPokemonCommand {
	cmd := UpdateTrainedPokemonCommand{
		factory.NewTrainedPokemonInput(id, gender, nickname, description, isPrivate, userId, adjustment),
	}
	return cmd
}
