package moves

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/errs"
)

var _ entity.IDomain[identifier.MoveId, uint16] = (*Move)(nil)

type Move struct {
	id            identifier.MoveId
	name          string
	species       value.MoveSpecies
	pokemonType   value.PokemonType
	power         uint16  // 威力
	accuracyRate  float32 // 命中率
	pp            uint8
	isContained   bool
	canGuard      bool
	battleEffects *value.BattleEffects
}

func NewMove(id identifier.MoveId, name, species, pokemonType string,
	power uint64, accuracyRate float32, pp uint64, isContained, canGuard bool) (*Move, error) {
	if power > 300 {
		return nil, errs.ThrowErrorInvalidValue("Move", "power", string(rune(power)))
	}
	if power > 100 {
		return nil, errs.ThrowErrorInvalidValue("Move", "pp", string(rune(power)))
	}
	return &Move{
		id:           id,
		name:         name,
		species:      value.MoveSpecies(species),
		pokemonType:  value.NewPokemonType(pokemonType),
		power:        uint16(power),
		accuracyRate: accuracyRate,
		pp:           uint8(pp),
		isContained:  isContained,
		canGuard:     canGuard,
	}, nil
}

func (m Move) Id() identifier.MoveId {
	return m.id
}

func (m Move) Power() uint16 {
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
