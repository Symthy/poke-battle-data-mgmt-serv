package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

var _ battles.IPokemonBattleDataSet = (*PokemonBattleDataSet)(nil)

// Todo: refactor
type PokemonBattleDataSet struct {
	*AttackSidePokemon
	*DefenseSidePokemon
	AttackMove
	weather                     WeatherState
	fieldState                  FieldStateType
	typeCompatibilityDamageRate float32
}

func NewPokemonBattleDataSet(
	attackSide *AttackSidePokemon, defenseSide *DefenseSidePokemon, attackMove AttackMove,
	typeCompatibilityDamageRate float32,
) *PokemonBattleDataSet {
	data := &PokemonBattleDataSet{
		AttackSidePokemon:           attackSide,
		DefenseSidePokemon:          defenseSide,
		AttackMove:                  attackMove,
		typeCompatibilityDamageRate: typeCompatibilityDamageRate,
	}
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
func (p PokemonBattleDataSet) AttackPokemonActualValueS(statusCorrections *battles.StatusCorrections) uint16 {
	valueS := p.AttackSidePokemon.attackPokemonActualValues.S()
	if statusCorrections == nil {
		return valueS
	}
	return statusCorrections.Apply(valueS, battles.CorrectionStatusS, p, battles.BattleAttackSide)
}
func (p PokemonBattleDataSet) DefensePokemonTypeOfFirst() value.PokemonType {
	return p.DefenseSidePokemon.defensePokemonTypes.FirstType()
}
func (p PokemonBattleDataSet) DefensePokemonTypeOfSecond() value.PokemonType {
	return p.DefenseSidePokemon.defensePokemonTypes.SecondType()
}
func (p PokemonBattleDataSet) DefensePokemonActualValueS(statusCorrections *battles.StatusCorrections) uint16 {
	valueS := p.DefenseSidePokemon.defensePokemonActualValues.S()
	if statusCorrections == nil {
		return valueS
	}
	return statusCorrections.Apply(valueS, battles.CorrectionStatusS, p, battles.BattleAttackSide)
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
