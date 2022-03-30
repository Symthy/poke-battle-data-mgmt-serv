package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

var _ IPokemonBattleDataSet = (*PokemonBattleDataSet)(nil)

type PokemonBattleDataSet struct {
	AttackSidePokemon
	AttackSideBattleEffects
	DefenceSidePokemon
	DefenceSideBattleEffects
	AttackMove
	correctedAttackPokemon      value.PokemonActualValues
	correctedDefencePokemon     value.PokemonActualValues
	weather                     WeatherType
	fieldState                  FieldStateType
	typeCompatibilityDamageRate float32
}

func NewPokemonBattleDataSet(
	attackSide AttackSidePokemon, attackEffects AttackSideBattleEffects,
	deffenceSide DefenceSidePokemon, defenceEffects DefenceSideBattleEffects,
	attackMove AttackMove, compatibilityDamageRate float32,
) PokemonBattleDataSet {
	data := PokemonBattleDataSet{
		AttackSidePokemon:           attackSide,
		AttackSideBattleEffects:     attackEffects,
		DefenceSidePokemon:          deffenceSide,
		DefenceSideBattleEffects:    defenceEffects,
		AttackMove:                  attackMove,
		typeCompatibilityDamageRate: compatibilityDamageRate,
	}
	statusAppliers := attackEffects.SupplyAllStatusCorrectionApplier(data)
	data.correctedAttackPokemon = data.attackPokemonActualValues.ApplyCorrection(statusAppliers)
	data.correctedDefencePokemon = data.defencePokemonActualValues.ApplyCorrection(statusAppliers)
	return data
}

func (b PokemonBattleDataSet) IsNoDamage() bool {
	return int(b.typeCompatibilityDamageRate) == 0
}

func (p PokemonBattleDataSet) AttackPokemonTypeOfFirst() string
func (p PokemonBattleDataSet) AttackPokemonTypeOfSecond() string
func (p PokemonBattleDataSet) AttackPokemonActualValueS() string
func (p PokemonBattleDataSet) DefencePokemonTypeOfFirst() string
func (p PokemonBattleDataSet) DefencePokemonTypeOfSecond() string
func (p PokemonBattleDataSet) DefencePokemonActualValueS() string
func (p PokemonBattleDataSet) MovePokemonType() string
func (p PokemonBattleDataSet) HasItemAttackSide() bool
func (p PokemonBattleDataSet) HasItemDefenceSide() bool

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

func (p PokemonBattleDataSet) ResolveDefenceValue() int {
	if p.species.IsPhysical() {
		return p.correctedDefencePokemon.B()
	}
	if p.species.IsSpecial() {
		return p.correctedDefencePokemon.D()
	}
	return 0
}
