package command

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type CreateTrainedPokemonCommand struct {
	factory.TrainedPokemonInput
}

func NewCreateTrainedPokemonCommand(
	request *server.RequestCreateTrainedPokemon, userId uint) CreateTrainedPokemonCommand {
	input := buildTrainedPokemonInput(request, userId)
	cmd := CreateTrainedPokemonCommand{input}
	return cmd
}

func buildTrainedPokemonInput(
	request *server.RequestCreateTrainedPokemon, userId uint) factory.TrainedPokemonInput {
	builder := factory.NewTrainedPokemonBuilder()
	if request.Gender != nil {
		builder.Gender(*request.Gender)
	}
	if request.Nickname != nil {
		builder.Nickname(*request.Nickname)
	}
	if request.Description != nil {
		builder.Description(*request.Description)
	}
	builder.SetIsPrivate(request.IsPrivate)
	builder.UserId(userId)

	adjustmentBuilder := factory.NewTrainedPokemonAdjustmentBuilder()
	adjustmentReq := request.Adjustment
	adjustmentBuilder.PokemonId(uint(adjustmentReq.PokemonId))
	adjustmentBuilder.Nature(adjustmentReq.Nature)
	adjustmentBuilder.AbilityId(uint(adjustmentReq.AbilityId))
	if adjustmentReq.HeldItemId != nil {
		adjustmentBuilder.HeldItemId(uint(*adjustmentReq.HeldItemId))
	}
	if adjustmentReq.EffortValueH != nil {
		adjustmentBuilder.EffortValueH(*adjustmentReq.EffortValueH)
	}
	if adjustmentReq.EffortValueA != nil {
		adjustmentBuilder.EffortValueA(*adjustmentReq.EffortValueA)
	}
	if adjustmentReq.EffortValueB != nil {
		adjustmentBuilder.EffortValueB(*adjustmentReq.EffortValueB)
	}
	if adjustmentReq.EffortValueC != nil {
		adjustmentBuilder.EffortValueC(*adjustmentReq.EffortValueC)
	}
	if adjustmentReq.EffortValueD != nil {
		adjustmentBuilder.EffortValueD(*adjustmentReq.EffortValueD)
	}
	if adjustmentReq.EffortValueS != nil {
		adjustmentBuilder.EffortValueS(*adjustmentReq.EffortValueS)
	}
	if len(adjustmentReq.MoveIds) >= 1 {
		adjustmentBuilder.MoveIdFirst(uint(adjustmentReq.MoveIds[0]))
	}
	if len(adjustmentReq.MoveIds) >= 2 {
		adjustmentBuilder.MoveIdFirst(uint(adjustmentReq.MoveIds[1]))
	}
	if len(adjustmentReq.MoveIds) >= 3 {
		adjustmentBuilder.MoveIdFirst(uint(adjustmentReq.MoveIds[2]))
	}
	if len(adjustmentReq.MoveIds) >= 4 {
		adjustmentBuilder.MoveIdFirst(uint(adjustmentReq.MoveIds[3]))
	}
	builder.Adjustment(adjustmentBuilder)
	return builder
}
