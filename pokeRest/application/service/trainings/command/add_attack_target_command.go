package command

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

type AddAdjustmentAttackTargetCommand struct {
	*factory.TrainedAttackTargetInput
}

func NewAddAdjustmentAttackTargetCommand(
	request *server.RequestCreateTrainedPokeAttackTarget) AddAdjustmentAttackTargetCommand {
	return AddAdjustmentAttackTargetCommand{
		buildTrainedAttackTargetInput(request),
	}
}

func buildTrainedAttackTargetInput(
	request *server.RequestCreateTrainedPokeAttackTarget) *factory.TrainedAttackTargetInput {
	builder := factory.NewTrainedAttackTargetBuilder()
	builder.MoveId(uint(*request.MoveId))
	builder.TrainedPokemonId(uint(*request.Target.PokemonId))
	//builder.TargetPokemonAbilityId(request.Target.)
	builder.TargetPokemonNature(*request.Target.Nature)
	builder.TargetPokemonEffortValueH(*request.Target.EffortValueH)
	builder.TargetPokemonEffortValueB(*request.Target.EffortValueB)
	builder.TargetPokemonEffortValueD(*request.Target.EffortValueD)
	return builder
}
