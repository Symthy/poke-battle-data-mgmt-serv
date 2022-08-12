package dto

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ trainings.ITrainedPokemonAdjustNotification = (*TrainedPokemonAdjustmentSchemaBuilder)(nil)

type TrainedPokemonAdjustmentSchemaBuilder struct {
	*schema.TrainedPokemonAdjustment
}

func NewTrainedPokemonAdjustmentSchemaBuilder() *TrainedPokemonAdjustmentSchemaBuilder {
	return &TrainedPokemonAdjustmentSchemaBuilder{&schema.TrainedPokemonAdjustment{}}
}

func (b *TrainedPokemonAdjustmentSchemaBuilder) SetId(id identifier.TrainedAdjustmentId) {
	b.ID = id.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetPokemonId(id identifier.PokemonId) {
	b.PokemonID = id.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetNature(nature value.PokemonNature) {
	b.Nature = enum.Nature(nature.String())
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetAbilityId(id identifier.AbilityId) {
	*b.AbilityID = id.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetHeldItemId(id identifier.HeldItemId) {
	*b.HeldItemID = id.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetEffortValueH(valueH value.EffortValue) {
	b.EffortValueH = valueH.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetEffortValueA(valueA value.EffortValue) {
	b.EffortValueA = valueA.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetEffortValueB(valueB value.EffortValue) {
	b.EffortValueB = valueB.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetEffortValueC(valueC value.EffortValue) {
	b.EffortValueC = valueC.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetEffortValueD(valueD value.EffortValue) {
	b.EffortValueD = valueD.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetEffortValueS(valueS value.EffortValue) {
	b.EffortValueS = valueS.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetMoveId1(id identifier.MoveId) {
	*b.MoveID1 = id.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetMoveId2(id identifier.MoveId) {
	*b.MoveID2 = id.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetMoveId3(id identifier.MoveId) {
	*b.MoveID3 = id.Value()
}
func (b *TrainedPokemonAdjustmentSchemaBuilder) SetMoveId4(id identifier.MoveId) {
	*b.MoveID4 = id.Value()
}

func (b TrainedPokemonAdjustmentSchemaBuilder) Build() *schema.TrainedPokemonAdjustment {
	return b.TrainedPokemonAdjustment
}
