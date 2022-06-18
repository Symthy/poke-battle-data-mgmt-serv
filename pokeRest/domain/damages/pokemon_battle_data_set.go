package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

var _ IPokemonBattleDataSet = (*PokemonBattleDataSet)(nil)

// Todo: refactor
type PokemonBattleDataSet struct {
	*AttackSidePokemon
	*AttackSideBattleEffects
	*DefenseSidePokemon
	*DefenseSideBattleEffects
	AttackMove
	correctedAttackPokemonActualValues  *value.PokemonActualValues
	correctedDefensePokemonActualValues *value.PokemonActualValues
	weather                             WeatherState
	fieldState                          FieldStateType
	typeCompatibilityDamageRate         float32
}

func NewPokemonBattleDataSet(
	attackSide *AttackSidePokemon, attackEffects *AttackSideBattleEffects,
	defenseSide *DefenseSidePokemon, defenseEffects *DefenseSideBattleEffects,
	attackMove AttackMove, typeCompatibilityDamageRate float32,
) *PokemonBattleDataSet {
	data := &PokemonBattleDataSet{
		AttackSidePokemon:           attackSide,
		AttackSideBattleEffects:     attackEffects,
		DefenseSidePokemon:          defenseSide,
		DefenseSideBattleEffects:    defenseEffects,
		AttackMove:                  attackMove,
		typeCompatibilityDamageRate: typeCompatibilityDamageRate,
	}
	allStatusApplier := attackEffects.SupplyAllStatusCorrectionApplier(data)
	data.correctedAttackPokemonActualValues = allStatusApplier(data.attackPokemonActualValues)
	data.correctedDefensePokemonActualValues = allStatusApplier(data.defensePokemonActualValues)
	return data
}

func (b PokemonBattleDataSet) IsNoDamage() bool {
	return int(b.typeCompatibilityDamageRate) == 0
}

func (p PokemonBattleDataSet) AttackPokemonTypeOfFirst() value.PokemonType {
	return p.AttackSidePokemon.attackPokemonType.FirstType()
}
func (p PokemonBattleDataSet) AttackPokemonTypeOfSecond() value.PokemonType {
	return p.AttackSidePokemon.attackPokemonType.SecondType()
}
func (p PokemonBattleDataSet) AttackPokemonActualValueS() uint16 {
	return p.correctedAttackPokemonActualValues.S()
}
func (p PokemonBattleDataSet) DefensePokemonTypeOfFirst() value.PokemonType {
	return p.DefenseSidePokemon.defensePokemonTypes.FirstType()
}
func (p PokemonBattleDataSet) DefensePokemonTypeOfSecond() value.PokemonType {
	return p.DefenseSidePokemon.defensePokemonTypes.SecondType()
}
func (p PokemonBattleDataSet) DefensePokemonActualValueS() uint16 {
	return p.correctedDefensePokemonActualValues.S()
}
func (p PokemonBattleDataSet) MovePokemonType() value.PokemonType {
	return p.AttackMove.pokemonType
}
func (p PokemonBattleDataSet) HasItemAttackSide() bool {
	return p.AttackSidePokemon.attackPokemonHasItem
}
func (p PokemonBattleDataSet) HasItemDefenseSide() bool {
	return p.DefenseSidePokemon.defensePokemonHasItems
}
func (p PokemonBattleDataSet) TypeCompatibilityDamageRate() float32 {
	return p.typeCompatibilityDamageRate
}

// Todo: move
func (p PokemonBattleDataSet) ResolvePowerValue() uint16 {
	attackPower := p.resolveAttackPowerValue()
	movePower := p.resolveMovePowerValue()

	return fmath.RoundUpIfDecimalGreaterFive[uint16](float64(movePower*attackPower) / 4096)
}

func (p PokemonBattleDataSet) resolveMovePowerValue() uint16 {
	applier := p.AttackSideBattleEffects.SupplyMovePowerCorrectionApplier(p.species, p)
	return applier(p.AttackMove.power)
}

func (p PokemonBattleDataSet) resolveAttackPowerValue() uint16 {
	applier := p.AttackSideBattleEffects.SupplyPowerCorrectionApplier(p.species, p)
	return applier(4096)
}

func (p PokemonBattleDataSet) ResolveAttackValue() uint16 {
	if p.species.IsPhysical() {
		return p.correctedAttackPokemonActualValues.A()
	}
	if p.species.IsSpecial() {
		return p.correctedAttackPokemonActualValues.C()
	}
	return 0
}

func (p PokemonBattleDataSet) ResolveDefenseValue() uint16 {
	if p.species.IsPhysical() {
		return p.correctedDefensePokemonActualValues.B()
	}
	if p.species.IsSpecial() {
		return p.correctedDefensePokemonActualValues.D()
	}
	return 0
}
