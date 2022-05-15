package command

import "github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"

type DamageCalculation struct {
	attackPokemonId     uint64
	attackNature        string
	attackAbilityId     uint64
	attackHeldItemId    uint64
	attackEffortValueH  uint64
	attackEffortValueA  uint64
	attackEffortValueB  uint64
	attackEffortValueC  uint64
	attackEffortValueD  uint64
	attackEffortValueS  uint64
	moveId              uint64
	defencePokemonId    uint64
	defenceNature       string
	defenceAbilityId    uint64
	defenceHeldItemId   uint64
	defenceEffortValueH uint64
	defenceEffortValueA uint64
	defenceEffortValueB uint64
	defenceEffortValueC uint64
	defenceEffortValueD uint64
	defenceEffortValueS uint64
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
	attackEffortValueS := uint64(0) // Todo
	defenceNature := resolveNullableString(request.Target.Nature)
	defenceAbilityId := resolveNullableId(request.Target.AbilityId)
	defenceHeldItemId := resolveNullableId(request.Target.HeldItemId)
	defenceEffortValueH := resolveEffortValue(request.Target.EffortValueH)
	defenceEffortValueA := resolveEffortValue(request.Target.EffortValueA)
	defenceEffortValueB := resolveEffortValue(request.Target.EffortValueB)
	defenceEffortValueC := resolveEffortValue(request.Target.EffortValueC)
	defenceEffortValueD := resolveEffortValue(request.Target.EffortValueD)
	defenceEffortValueS := uint64(0) // Todo

	return DamageCalculation{
		attackPokemonId:     request.Self.PokemonId,
		attackNature:        attackNature,
		attackAbilityId:     attackAbilityId,
		attackHeldItemId:    attackHeldItemId,
		attackEffortValueH:  attackEffortValueH,
		attackEffortValueA:  attackEffortValueA,
		attackEffortValueB:  attackEffortValueB,
		attackEffortValueC:  attackEffortValueC,
		attackEffortValueD:  attackEffortValueD,
		attackEffortValueS:  attackEffortValueS,
		moveId:              request.MoveId,
		defencePokemonId:    request.Target.PokemonId,
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

func resolveNullableId(nullableId *uint64) uint64 {
	var id uint64 = 0
	if nullableId != nil {
		id = *nullableId
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

func resolveEffortValue(effortValue *uint64) uint64 {
	var value uint64 = 0
	if effortValue != nil {
		value = *effortValue
	}
	return value
}
