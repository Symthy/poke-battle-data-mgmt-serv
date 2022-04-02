package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

var _ IPokemonBattleDataSet = (*PokemonBattleDataSet)(nil)

type PokemonBattleDataSet struct {
	AttackSidePokemon
	AttackSideBattleEffects
	DefenseSidePokemon
	DefenseSideBattleEffects
	AttackMove
	correctedAttackPokemon      value.PokemonActualValues
	correctedDefensePokemon     value.PokemonActualValues
	weather                     WeatherState
	fieldState                  FieldStateType
	typeCompatibilityDamageRate float32
}

func NewPokemonBattleDataSet(
	attackSide AttackSidePokemon, attackEffects AttackSideBattleEffects,
	defenseSide DefenseSidePokemon, defenseEffects DefenseSideBattleEffects,
	attackMove AttackMove, compatibilityDamageRate float32,
) PokemonBattleDataSet {
	data := PokemonBattleDataSet{
		AttackSidePokemon:           attackSide,
		AttackSideBattleEffects:     attackEffects,
		DefenseSidePokemon:          defenseSide,
		DefenseSideBattleEffects:    defenseEffects,
		AttackMove:                  attackMove,
		typeCompatibilityDamageRate: compatibilityDamageRate,
	}
	statusAppliers := attackEffects.SupplyAllStatusCorrectionApplier(data)
	data.correctedAttackPokemon = data.attackPokemonActualValues.ApplyCorrection(statusAppliers)
	data.correctedDefensePokemon = data.defensePokemonActualValues.ApplyCorrection(statusAppliers)
	return data
}

func (b PokemonBattleDataSet) IsNoDamage() bool {
	return int(b.typeCompatibilityDamageRate) == 0
}

func (p PokemonBattleDataSet) AttackPokemonTypeOfFirst() value.PokemonType {
	return p.AttackSidePokemon.attackPokemonType.FirstType()
}
func (p PokemonBattleDataSet) AttackPokemonTypeOfSecond() value.PokemonType
func (p PokemonBattleDataSet) AttackPokemonActualValueS() int
func (p PokemonBattleDataSet) DefensePokemonTypeOfFirst() value.PokemonType
func (p PokemonBattleDataSet) DefensePokemonTypeOfSecond() value.PokemonType
func (p PokemonBattleDataSet) DefensePokemonActualValueS() int
func (p PokemonBattleDataSet) MovePokemonType() value.PokemonType
func (p PokemonBattleDataSet) HasItemAttackSide() bool
func (p PokemonBattleDataSet) HasItemDefenseSide() bool

// Todo: move
func (p PokemonBattleDataSet) ResolvePowerValue() int {
	attackPower := p.resolveAttackPowerValue()
	movePower := p.resolveMovePowerValue()

	return fmath.RoundUpIfDecimalGreaterFive(float64(movePower*attackPower) / 4096)
}

func (p PokemonBattleDataSet) resolveMovePowerValue() int {
	applier := p.AttackSideBattleEffects.SupplyMovePowerCorrectionApplier(p.species, p)
	return applier(p.AttackMove.power)
}

func (p PokemonBattleDataSet) resolveAttackPowerValue() int {
	applier := p.AttackSideBattleEffects.SupplyPowerCorrectionApplier(p.species, p)
	return applier(4096)
}

func (p PokemonBattleDataSet) ResolveAttackValue() int {
	if p.species.IsPhysical() {
		return p.correctedAttackPokemon.A()
	}
	if p.species.IsSpecial() {
		return p.correctedAttackPokemon.C()
	}
	return 0
}

func (p PokemonBattleDataSet) ResolveDefenseValue() int {
	if p.species.IsPhysical() {
		return p.correctedDefensePokemon.B()
	}
	if p.species.IsSpecial() {
		return p.correctedDefensePokemon.D()
	}
	return 0
}
