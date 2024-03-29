package trainings

import (
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type ITrainedPokemonAdjustNotification interface {
	SetId(identifier.TrainedAdjustmentId)
	SetPokemonId(identifier.PokemonId)
	SetNature(value.PokemonNature)
	SetAbilityId(identifier.AbilityId)
	SetHeldItemId(identifier.HeldItemId)
	SetEffortValueH(value.EffortValue)
	SetEffortValueA(value.EffortValue)
	SetEffortValueB(value.EffortValue)
	SetEffortValueC(value.EffortValue)
	SetEffortValueD(value.EffortValue)
	SetEffortValueS(value.EffortValue)
	SetMoveId1(identifier.MoveId)
	SetMoveId2(identifier.MoveId)
	SetMoveId3(identifier.MoveId)
	SetMoveId4(identifier.MoveId)
}

type ITrainedPokemonTargetNotification interface {
	SetTrainedPokemonId(identifier.TrainedPokemonId)
	SetMoveId(identifier.MoveId)
	SetTargetPokemonId(identifier.PokemonId)
	SetTargetPokemonNature(value.PokemonNature)
}

type ITrainedPokemonAttackNotification interface {
	SetId(identifier.TrainedAttackTargetId)
	ITrainedPokemonTargetNotification
	SetTargetPokemonEffortH(value.EffortValue)
	SetTargetPokemonEffortB(value.EffortValue)
	SetTargetPokemonEffortD(value.EffortValue)
}

type ITrainedPokemonDefenseNotification interface {
	SetId(identifier.TrainedDefenseTargetId)
	ITrainedPokemonTargetNotification
	SetTargetPokemonEffortA(value.EffortValue)
	SetTargetPokemonEffortC(value.EffortValue)
}

type ITrainedPokemonNotification interface {
	SetId(identifier.TrainedPokemonId)
	SetGender(value.Gender)
	SetNickname(string)
	SetDescription(string)
	SetIsPrivate(bool)
	SetUserId(identifier.UserId)
	// SetAdjustment(TrainedPokemonAdjustment)
}
