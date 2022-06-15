package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/effects"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type MoveInput struct {
	id            uint64
	name          string
	species       string
	moveType      string
	power         uint64
	accuracy      float32
	pp            uint64
	isContained   bool
	isCanGuard    bool
	battleEffects *effects.BattleEffects
}

func NewMoveInput(id uint64, name string, species string, moveType string, power uint64, accuracy float32, pp uint64,
	isContained bool, isCanGuard bool) MoveInput {
	return MoveInput{
		id:            id,
		name:          name,
		species:       species,
		moveType:      moveType,
		power:         power,
		accuracy:      accuracy,
		pp:            pp,
		isContained:   isContained,
		isCanGuard:    isCanGuard,
		battleEffects: nil,
	}
}

func NewMoveBuilder() *MoveInput {
	return &MoveInput{}
}

func (i MoveInput) Id(id uint64) {
	i.id = id
}
func (i MoveInput) Name(name string) {
	i.name = name
}
func (i MoveInput) Species(species string) {
	i.species = species
}
func (i MoveInput) MoveType(moveType string) {
	i.moveType = moveType
}
func (i MoveInput) Power(power uint64) {
	i.power = power
}
func (i MoveInput) Accuracy(accuracy float32) {
	i.accuracy = accuracy
}
func (i MoveInput) PP(pp uint64) {
	i.pp = pp
}
func (i MoveInput) SetIsContained(isContained bool) {
	i.isContained = isContained
}
func (i MoveInput) SetIsCanGuard(isCanGuard bool) {
	i.isCanGuard = isCanGuard
}
func (i MoveInput) BattleEffects(battleEffects *effects.BattleEffects) {
	i.battleEffects = battleEffects
}

func (i MoveInput) BuildDomain() (*moves.Move, error) {
	id, err := identifier.NewMoveId(i.id)
	if err != nil {
		return nil, err
	}
	domain, err := moves.NewMove(*id, i.name, i.species, i.moveType, i.power, i.accuracy, i.pp, i.isContained, i.isCanGuard)
	if err != nil {
		return nil, err
	}
	return domain, nil
}
