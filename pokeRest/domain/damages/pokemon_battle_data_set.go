package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

var _ battles.IPokemonBattleDataSet = (*PokemonBattleDataSet)(nil)

// Todo: refactor
type PokemonBattleDataSet struct {
	attackSide                  *AttackSidePokemon
	attackAbnormalState         *AbnormalState
	defenseSide                 *DefenseSidePokemon
	defenseAbnormalState        *AbnormalState
	attackMove                  *AttackMove
	weather                     WeatherState
	fieldState                  FieldState
	typeCompatibilityDamageRate float32
	attackEffects               *AttackSideBattleEffects
	defenseEffects              *DefenseSideBattleEffects
}

func NewPokemonBattleDataSet(
	attackSide *AttackSidePokemon, defenseSide *DefenseSidePokemon, attackMove *AttackMove,
	typeCompatibilityDamageRate float32,
) *PokemonBattleDataSet {
	data := &PokemonBattleDataSet{
		attackSide:                  attackSide,
		defenseSide:                 defenseSide,
		attackMove:                  attackMove,
		typeCompatibilityDamageRate: typeCompatibilityDamageRate,
	}
	return data
}

func (b PokemonBattleDataSet) IsNoDamage() bool {
	return int(b.typeCompatibilityDamageRate) == 0
}
func (p PokemonBattleDataSet) MoveSpecies() value.MoveSpecies {
	return p.attackMove.species
}
func (p PokemonBattleDataSet) AttackPokemonTypeOfFirst() value.PokemonType {
	return p.attackSide.attackPokemonType.FirstType()
}
func (p PokemonBattleDataSet) AttackPokemonTypeOfSecond() value.PokemonType {
	return p.attackSide.attackPokemonType.SecondType()
}
func (p PokemonBattleDataSet) AttackPokemonActualValues() *value.PokemonActualValues {
	return p.attackSide.attackPokemonActualValues
}
func (p PokemonBattleDataSet) AttackCorrectedActualValue() uint16 {
	if p.attackMove.species.IsPhysical() {
		return p.attackEffects.statusCorrections.ApplyA(p)
	}
	if p.attackMove.species.IsSpecial() {
		return p.attackEffects.statusCorrections.ApplyC(p)
	}
	return 0
}
func (p PokemonBattleDataSet) AttackPokemonCorrectedActualValueS() uint16 {
	return p.attackEffects.statusCorrections.ApplyS(p)
}
func (p PokemonBattleDataSet) DefensePokemonTypeOfFirst() value.PokemonType {
	return p.defenseSide.defensePokemonTypes.FirstType()
}
func (p PokemonBattleDataSet) DefensePokemonTypeOfSecond() value.PokemonType {
	return p.defenseSide.defensePokemonTypes.SecondType()
}
func (p PokemonBattleDataSet) DefensePokemonActualValues() *value.PokemonActualValues {
	return p.defenseSide.defensePokemonActualValues
}
func (p PokemonBattleDataSet) DefenseCorrectedActualValue() uint16 {
	if p.attackMove.species.IsPhysical() {
		return p.defenseEffects.statusCorrections.ApplyB(p)
	}
	if p.attackMove.species.IsSpecial() {
		return p.defenseEffects.statusCorrections.ApplyD(p)
	}
	return 0
}
func (p PokemonBattleDataSet) DefensePokemonCorrectedActualValueS() uint16 {
	val := p.defenseEffects.statusCorrections.ApplyS(p)
	return p.attackAbnormalState.ApplyCorrection(val)
}
func (p PokemonBattleDataSet) PowerCorrectedValue() uint16 {
	return p.attackEffects.powerCorrections.Apply(4096, p)
}
func (p PokemonBattleDataSet) MovePokemonType() value.PokemonType {
	return p.attackMove.pokemonType
}
func (p PokemonBattleDataSet) MovePowerValue() uint16 {
	return p.attackEffects.moveCorrections.Apply(p.attackMove.power, p)
}
func (p PokemonBattleDataSet) HasItemAttackSide() bool {
	return p.attackSide.attackPokemonHasItem
}
func (p PokemonBattleDataSet) HasItemDefenseSide() bool {
	return p.defenseSide.defensePokemonHasItems
}
func (p PokemonBattleDataSet) IsTypeMatch() bool {
	return p.attackSide.attackPokemonType.FirstType().Equals(p.attackMove.pokemonType) ||
		p.attackSide.attackPokemonType.SecondType().Equals(p.attackMove.pokemonType)
}
func (p PokemonBattleDataSet) TypeCompatibilityDamageRate() float32 {
	return p.typeCompatibilityDamageRate
}
func (p PokemonBattleDataSet) DamageCorrectedValue() uint16 {
	return p.attackEffects.damageCorrections.Apply(p)
}
func (p PokemonBattleDataSet) ApplyWeatherCorrection(damage uint16) uint16 {
	return p.weather.ApplyCorrection(damage, p)
}
func (p PokemonBattleDataSet) FieldCorrectedValue() uint16 {
	return p.fieldState.ApplyCorrection(p)
}
