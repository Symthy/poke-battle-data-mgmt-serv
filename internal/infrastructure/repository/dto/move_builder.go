package dto

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/moves"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

var _ moves.IMoveNotification = (*MoveSchemaBuilder)(nil)

type MoveSchemaBuilder struct {
	*schema.Move
}

func NewMoveSchemaBuilder() *MoveSchemaBuilder {
	return &MoveSchemaBuilder{&schema.Move{}}
}

func (b *MoveSchemaBuilder) SetId(id identifier.MoveId) {
	b.ID = id.Value()
}
func (b *MoveSchemaBuilder) SetName(name string) {
	b.Name = name
}
func (b *MoveSchemaBuilder) SetSpecies(species value.MoveSpecies) {
	b.Species = enum.MoveSpecies(species.String())
}
func (b *MoveSchemaBuilder) SetPower(power uint16) {
	b.Power = power
}
func (b *MoveSchemaBuilder) SetAccuracyRate(accuracy float32) {
	b.Accuracy = accuracy
}
func (b *MoveSchemaBuilder) SetPP(pp uint8) {
	b.PP = pp
}
func (b *MoveSchemaBuilder) SetIsContained(isContained bool) {
	b.IsContained = &isContained
}
func (b *MoveSchemaBuilder) SetCanGuard(canGuard bool) {
	b.CanGuard = &canGuard
}

func (b *MoveSchemaBuilder) Build() *schema.Move {
	return b.Move
}
