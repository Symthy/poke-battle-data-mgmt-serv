package trainings

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type TrainedPokemonDefenceTarget struct {
	id                   uint
	trainedPokemonId     uint
	moveId               int
	targetPokemonId      int
	targetPokemonNature  value.PokemonNature
	targetPokemonEffortA value.EffortValue
	targetPokemonEffortC value.EffortValue
}

func NewTrainedPokemonDefenceTarget(
	id uint, trainedPokemonId uint, moveId int, targetPokemonId int, targetPokemonNature string,
	targetPokemonEffortA int, targetPokemonEffortC int) TrainedPokemonDefenceTarget {
	return TrainedPokemonDefenceTarget{
		id:                   id,
		trainedPokemonId:     trainedPokemonId,
		moveId:               moveId,
		targetPokemonId:      targetPokemonId,
		targetPokemonNature:  value.NewPokemonNature(targetPokemonNature),
		targetPokemonEffortA: value.NewEffortValue(targetPokemonEffortA),
		targetPokemonEffortC: value.NewEffortValue(targetPokemonEffortC),
	}
}

func (t TrainedPokemonDefenceTarget) Id() uint {
	return t.id
}

func (t TrainedPokemonDefenceTarget) TrainedPokemonId() uint {
	return t.trainedPokemonId
}

func (t TrainedPokemonDefenceTarget) MoveId() int {
	return t.moveId
}

func (t TrainedPokemonDefenceTarget) TargetPokemonId() int {
	return t.targetPokemonId
}

func (t TrainedPokemonDefenceTarget) TargetPokemonEffortA() value.EffortValue {
	return t.targetPokemonEffortA
}

func (t TrainedPokemonDefenceTarget) TargetPokemonEffortC() value.EffortValue {
	return t.targetPokemonEffortC
}
