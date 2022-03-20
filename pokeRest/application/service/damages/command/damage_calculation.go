package command

import "github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"

type DamageCalculation struct {
	attackPokemonId     uint
	attackNature        string
	attackAbilityId     uint
	attackHeldItemId    uint
	attackEffortValueH  int
	attackEffortValueA  int
	attackEffortValueB  int
	attackEffortValueC  int
	attackEffortValueD  int
	attackEffortValueS  int
	moveId              uint
	defencePokemonId    uint
	defenceNature       string
	defenceAbilityId    uint
	defenceHeldItemId   uint
	defenceEffortValueH int
	defenceEffortValueA int
	defenceEffortValueB int
	defenceEffortValueC int
	defenceEffortValueD int
	defenceEffortValueS int
}

func (c DamageCalculation) NewDamageCalculation(request server.RequestCalclateSingleDamage) DamageCalculation {
	attackNature := resolveNullableString(request.Self.Nature)
	attackAbilityId := resolveNullableId(request.Self.AbilityId)
	attackHeldItemId := resolveNullableId(request.Self.HeldItemId)
	attackEffortValueH := resolveEffortValue(request.Self.EffortValueH)
	attackEffortValueA := resolveEffortValue(request.Self.EffortValueA)
	attackEffortValueB := resolveEffortValue(request.Self.EffortValueB)
	attackEffortValueC := resolveEffortValue(request.Self.EffortValueC)
	attackEffortValueD := resolveEffortValue(request.Self.EffortValueD)
	attackEffortValueS := 0 // Todo
	defenceNature := resolveNullableString(request.Target.Nature)
	defenceAbilityId := resolveNullableId(request.Target.AbilityId)
	defenceHeldItemId := resolveNullableId(request.Target.HeldItemId)
	defenceEffortValueH := resolveEffortValue(request.Target.EffortValueH)
	defenceEffortValueA := resolveEffortValue(request.Target.EffortValueA)
	defenceEffortValueB := resolveEffortValue(request.Target.EffortValueB)
	defenceEffortValueC := resolveEffortValue(request.Target.EffortValueC)
	defenceEffortValueD := resolveEffortValue(request.Target.EffortValueD)
	defenceEffortValueS := 0 // Todo

	return DamageCalculation{
		attackPokemonId:     uint(request.Self.PokemonId),
		attackNature:        attackNature,
		attackAbilityId:     attackAbilityId,
		attackHeldItemId:    attackHeldItemId,
		attackEffortValueH:  attackEffortValueH,
		attackEffortValueA:  attackEffortValueA,
		attackEffortValueB:  attackEffortValueB,
		attackEffortValueC:  attackEffortValueC,
		attackEffortValueD:  attackEffortValueD,
		attackEffortValueS:  attackEffortValueS,
		moveId:              uint(request.MoveId),
		defencePokemonId:    uint(request.Target.PokemonId),
		defenceNature:       defenceNature,
		defenceAbilityId:    defenceAbilityId,
		defenceHeldItemId:   defenceHeldItemId,
		defenceEffortValueH: defenceEffortValueH,
		defenceEffortValueA: defenceEffortValueA,
		defenceEffortValueB: defenceEffortValueB,
		defenceEffortValueC: defenceEffortValueC,
		defenceEffortValueD: defenceEffortValueD,
		defenceEffortValueS: defenceEffortValueS,
	}

}

func resolveNullableId(nullableId *int) uint {
	var id uint = 0
	if nullableId != nil {
		id = uint(*nullableId)
	}
	return id
}

func resolveNullableString(value *string) string {
	result := ""
	if value != nil {
		result = *value
	}
	return result
}

func resolveEffortValue(effortValue *int) int {
	var value int = 0
	if effortValue != nil {
		value = *effortValue
	}
	return value
}
