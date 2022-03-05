package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type MoveInput struct {
	id          uint
	name        string
	species     string
	power       int
	accuracy    float32
	pp          int
	isContained bool
	isCanGuard  bool
}

func NewMoveInput(id uint, name string, species string, power int, accuracy float32, pp int,
	isContained bool, isCanGuard bool) MoveInput {
	return MoveInput{id, name, species, power, accuracy, pp, isContained, isCanGuard}
}

func (i MoveInput) BuildDomain() (*moves.Move, error) {
	id, err := identifier.NewMoveId(i.id)
	if err != nil {
		return nil, err
	}
	domain := moves.NewMove(*id, i.name, i.species, i.power, i.accuracy, i.pp, i.isContained, i.isCanGuard)
	return &domain, nil
}
