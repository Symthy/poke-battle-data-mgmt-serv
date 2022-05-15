package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type MoveInput struct {
	id          uint64
	name        string
	species     string
	moveType    string
	power       uint64
	accuracy    float32
	pp          uint64
	isContained bool
	isCanGuard  bool
}

func NewMoveInput(id uint64, name string, species string, moveType string, power uint64, accuracy float32, pp uint64,
	isContained bool, isCanGuard bool) MoveInput {
	return MoveInput{id, name, species, moveType, power, accuracy, pp, isContained, isCanGuard}
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
