package command

import (
	"github.com/Symthy/PokeRest/internal/domain/factory"
)

type UpdateTrainedPokemonCommand struct {
	factory.TrainedPokemonInput
}

// Todo
func NewUpdateTrainedPokemonCommand(id uint64, gender, nickname, description string, isPrivate bool,
	userId uint64, adjustment factory.TrainedPokemonAdjustmentInput) UpdateTrainedPokemonCommand {
	cmd := UpdateTrainedPokemonCommand{
		factory.NewTrainedPokemonInput(id, gender, nickname, description, isPrivate, userId, adjustment),
	}
	return cmd
}
