package moves

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.MoveId] = (*Move)(nil)

type Move struct {
	id           identifier.MoveId
	name         string
	species      value.MoveSpecies
	power        int     // 威力
	accuracyRate float32 // 命中率
	pp           int
	isContained  bool
	isCanGuard   bool
}

func NewMove(id identifier.MoveId, name string, species string, power int, accuracyRate float32,
	pp int, isContained bool, isCanGuard bool) Move {
	return Move{
		id:           id,
		name:         name,
		species:      value.MoveSpecies(species),
		power:        power,
		accuracyRate: accuracyRate,
		pp:           pp,
		isContained:  isContained,
		isCanGuard:   isCanGuard,
	}
}

func (m Move) Id() identifier.MoveId {
	return m.id
}

func (m Move) Notify(note IMoveNotification) {
	note.SetId(m.id)
	note.SetName(m.name)
	note.SetSpecies(m.species)
	note.SetPower(m.power)
	note.SetAccuracyRate(m.accuracyRate)
	note.SetPP(m.pp)
	note.SetIsContained(m.isContained)
	note.SetIsCanGuard(m.isCanGuard)
}
