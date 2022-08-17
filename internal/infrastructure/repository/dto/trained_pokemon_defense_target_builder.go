package dto

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/trainings"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

var _ trainings.ITrainedPokemonDefenseNotification = (*TrainedPokemonDefenseTargetSchemaBuilder)(nil)

type TrainedPokemonDefenseTargetSchemaBuilder struct {
	*schema.TrainedPokemonDefenseTarget
}

func NewTrainedPokemonDefenseTargetSchemaBuilder() *TrainedPokemonDefenseTargetSchemaBuilder {
	return &TrainedPokemonDefenseTargetSchemaBuilder{&schema.TrainedPokemonDefenseTarget{}}
}

func (b *TrainedPokemonDefenseTargetSchemaBuilder) SetId(id identifier.TrainedDefenseTargetId) {
	b.ID = id.Value()
}
func (b *TrainedPokemonDefenseTargetSchemaBuilder) SetTargetPokemonEffortA(valueA value.EffortValue) {
	b.TargetPokemonEffortValueA = valueA.Value()
}
func (b *TrainedPokemonDefenseTargetSchemaBuilder) SetTargetPokemonEffortC(valueC value.EffortValue) {
	b.TargetPokemonEffortValueC = valueC.Value()
}

func (b *TrainedPokemonDefenseTargetSchemaBuilder) SetTrainedPokemonId(id identifier.TrainedPokemonId) {
	b.TrainedPokemonId = id.Value()
}
func (b *TrainedPokemonDefenseTargetSchemaBuilder) SetMoveId(id identifier.MoveId) {
	b.MoveId = id.Value()
}
func (b *TrainedPokemonDefenseTargetSchemaBuilder) SetTargetPokemonId(id identifier.PokemonId) {
	b.TargetPokemonId = id.Value()
}
func (b *TrainedPokemonDefenseTargetSchemaBuilder) SetTargetPokemonNature(nature value.PokemonNature) {
	b.TargetPokemonNature = enum.Nature(nature.String())
}

func (b TrainedPokemonDefenseTargetSchemaBuilder) Build() *schema.TrainedPokemonDefenseTarget {
	return b.TrainedPokemonDefenseTarget
}
