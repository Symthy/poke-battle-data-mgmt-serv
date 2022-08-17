package dto

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/trainings"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

var _ trainings.ITrainedPokemonAttackNotification = (*TrainedPokemonAttackTargetSchemaBuilder)(nil)

type TrainedPokemonAttackTargetSchemaBuilder struct {
	*schema.TrainedPokemonAttackTarget
}

func NewTrainedPokemonAttackTargetSchemaBuilder() *TrainedPokemonAttackTargetSchemaBuilder {
	return &TrainedPokemonAttackTargetSchemaBuilder{&schema.TrainedPokemonAttackTarget{}}
}

func (b *TrainedPokemonAttackTargetSchemaBuilder) SetId(id identifier.TrainedAttackTargetId) {
	b.ID = id.Value()
}
func (b *TrainedPokemonAttackTargetSchemaBuilder) SetTargetPokemonEffortH(valueH value.EffortValue) {
	b.TargetPokemonEffortValueH = valueH.Value()
}
func (b *TrainedPokemonAttackTargetSchemaBuilder) SetTargetPokemonEffortB(valueB value.EffortValue) {
	b.TargetPokemonEffortValueB = valueB.Value()
}
func (b *TrainedPokemonAttackTargetSchemaBuilder) SetTargetPokemonEffortD(valueD value.EffortValue) {
	b.TargetPokemonEffortValueD = valueD.Value()
}

func (b *TrainedPokemonAttackTargetSchemaBuilder) SetTrainedPokemonId(id identifier.TrainedPokemonId) {
	b.TrainedPokemonId = id.Value()
}
func (b *TrainedPokemonAttackTargetSchemaBuilder) SetMoveId(id identifier.MoveId) {
	b.MoveId = id.Value()
}
func (b *TrainedPokemonAttackTargetSchemaBuilder) SetTargetPokemonId(id identifier.PokemonId) {
	b.TargetPokemonId = id.Value()
}
func (b *TrainedPokemonAttackTargetSchemaBuilder) SetTargetPokemonNature(nature value.PokemonNature) {
	b.TargetPokemonNature = enum.Nature(nature.String())
}

func (b TrainedPokemonAttackTargetSchemaBuilder) Build() *schema.TrainedPokemonAttackTarget {
	return b.TrainedPokemonAttackTarget
}
