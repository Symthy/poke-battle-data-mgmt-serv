package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
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

type ITrainedPokemonTarget interface {
	SetTrainedPokemonId(identifier.TrainedPokemonId)
	SetMoveId(identifier.MoveId)
	SetTargetPokemonId(identifier.PokemonId)
	SetTargetPokemonNature(value.PokemonNature)
}

type ITrainedPokemonAttackNotification interface {
	SetId(identifier.TrainedAttackTargetId)
	ITrainedPokemonTarget
	SetTargetPokemonEffortH(value.EffortValue)
	SetTargetPokemonEffortB(value.EffortValue)
	SetTargetPokemonEffortD(value.EffortValue)
}

type ITrainedPokemonDefenceNotification interface {
	SetId(identifier.TrainedDefenseTargetId)
	ITrainedPokemonTarget
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
	SetAdjustment(TrainedPokemonAdjustment)
}
