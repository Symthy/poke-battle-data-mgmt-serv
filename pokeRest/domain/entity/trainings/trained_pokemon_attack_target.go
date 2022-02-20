package trainings

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type TrainedPokemonAttackTarget struct {
	id                   uint
	trainedPokemonId     uint
	moveId               int
	targetPokemonId      int
	targetPokemonNature  value.PokemonNature
	targetPokemonEffortH value.EffortValue
	targetPokemonEffortB value.EffortValue
	targetPokemonEffortD value.EffortValue
}

func NewTrainedPokemonAttackTarget(
	id uint, trainedPokemonId uint, moveId int, targetPokemonId int, targetPokemonNature string,
	targetPokemonEffortH int, targetPokemonEffortB int, targetPokemonEffortD int) TrainedPokemonAttackTarget {
	return TrainedPokemonAttackTarget{
		id:                   id,
		trainedPokemonId:     trainedPokemonId,
		moveId:               moveId,
		targetPokemonId:      targetPokemonId,
		targetPokemonNature:  value.NewPokemonNature(targetPokemonNature),
		targetPokemonEffortH: value.NewEffortValue(targetPokemonEffortH),
		targetPokemonEffortB: value.NewEffortValue(targetPokemonEffortB),
	}
}

func (t TrainedPokemonAttackTarget) Id() uint {
	return t.id
}

func (t TrainedPokemonAttackTarget) TrainedPokemonId() uint {
	return t.trainedPokemonId
}

func (t TrainedPokemonAttackTarget) MoveId() int {
	return t.moveId
}

func (t TrainedPokemonAttackTarget) TargetPokemonId() int {
	return t.targetPokemonId
}

func (t TrainedPokemonAttackTarget) TargetPokemonEffortH() value.EffortValue {
	return t.targetPokemonEffortH
}

func (t TrainedPokemonAttackTarget) TargetPokemonEffortB() value.EffortValue {
	return t.targetPokemonEffortB
}

func (t TrainedPokemonAttackTarget) TargetPokemonEffortD() value.EffortValue {
	return t.targetPokemonEffortD
}
