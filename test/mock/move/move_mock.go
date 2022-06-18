package move

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

func MoveAcrobatics() *moves.Move {
	builder := factory.NewMoveBuilder()
	builder.Id(1)
	builder.Name("アクロバット")
	builder.MovePokemonType(value.Flying().ToString())
	builder.Power(55)
	builder.PP(15)
	builder.Accuracy(100)
	builder.Species(value.MoveSpeciesPhysical.String())
	builder.SetCanGuard(true)
	builder.SetIsContained(true)

	condition := battles.NewTriggerCondition(battles.ConditionNotHaveItem, "true")
	correction := battles.NewBattleCorrectionValue(battles.CorrectionPhysicalPower, 2, condition)
	effectsBuilder := battles.NewBattleEffectsBuilder()
	effectsBuilder.AddCorrectionValues(correction)

	builder.BattleEffects(effectsBuilder.Build())
	move, _ := builder.BuildDomain()
	return move
}

func MoveHeadSmash() *moves.Move {
	builder := factory.NewMoveBuilder()
	builder.Id(306)
	builder.Name("もろはのずつき")
	builder.MovePokemonType(value.Rock().ToString())
	builder.Power(150)
	builder.PP(5)
	builder.Accuracy(80)
	builder.Species(value.MoveSpeciesPhysical.String())
	builder.SetCanGuard(true)
	builder.SetIsContained(true)
	move, _ := builder.BuildDomain()
	return move
}
