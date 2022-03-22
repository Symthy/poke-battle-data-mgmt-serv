package moves

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.MoveId] = (*Move)(nil)

type Move struct {
	id            identifier.MoveId
	name          string
	species       value.MoveSpecies
	pokemonType   value.PokemonType
	power         int     // 威力
	accuracyRate  float32 // 命中率
	pp            int
	isContained   bool
	canGuard      bool
	battleEffects value.BattleEffects
}

func NewMove(id identifier.MoveId, name string, species string, pokemonType string,
	power int, accuracyRate float32, pp int, isContained bool, canGuard bool) Move {
	return Move{
		id:           id,
		name:         name,
		species:      value.MoveSpecies(species),
		pokemonType:  value.NewPokemonType(pokemonType),
		power:        power,
		accuracyRate: accuracyRate,
		pp:           pp,
		isContained:  isContained,
		canGuard:     canGuard,
	}
}

func (m Move) Id() identifier.MoveId {
	return m.id
}

func (m Move) Power() int {
	return m.power
}

func (m Move) NotifyBattleEffects(effects *value.BattleEffects) {
	effects.Merge(m.battleEffects)
}

func (m Move) NotifyForDamageCalc(note IMoveNotificationForDamageCalc) {
	note.SetMoveType(m.pokemonType)
	note.SetMoveSpecies(m.species)
	note.SetMovePower(m.power)
	note.SetIsMoveContained(m.isContained)
	note.SetCanMoveGuard(m.canGuard)
}

func (m Move) Notify(note IMoveNotification) {
	note.SetId(m.id)
	note.SetName(m.name)
	note.SetSpecies(m.species)
	note.SetPower(m.power)
	note.SetAccuracyRate(m.accuracyRate)
	note.SetPP(m.pp)
	note.SetIsContained(m.isContained)
	note.SetCanGuard(m.canGuard)
}
