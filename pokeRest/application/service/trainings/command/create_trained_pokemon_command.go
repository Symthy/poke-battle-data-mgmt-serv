package command

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type CreateTrainedPokemonCommand struct {
	factory.TrainedPokemonInput
}

// Todo
func NewCreateTrainedPokemonCommand(
	gender, nickname, description string, isPrivate bool, userId uint,
	adjustment factory.TrainedPokemonAdjustmentInput) CreateTrainedPokemonCommand {
	cmd := CreateTrainedPokemonCommand{
		factory.NewTrainedPokemonInput(0, gender, nickname, description, isPrivate, userId, adjustment),
	}
	return cmd
}
