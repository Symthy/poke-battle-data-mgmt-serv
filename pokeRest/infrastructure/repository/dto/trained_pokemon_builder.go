package dto

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ trainings.ITrainedPokemonNotification = (*TrainedPokemonSchemaBuilder)(nil)

type TrainedPokemonSchemaBuilder struct {
	*schema.TrainedPokemon
}

func NewTrainedPokemonSchemaBuilder() *TrainedPokemonSchemaBuilder {
	return &TrainedPokemonSchemaBuilder{&schema.TrainedPokemon{}}
}

func (b *TrainedPokemonSchemaBuilder) SetId(id identifier.TrainedPokemonId) {
	b.ID = id.Value()
}
func (b *TrainedPokemonSchemaBuilder) SetGender(gender value.Gender) {
	b.Gender = enum.Gender(gender.String())
}
func (b *TrainedPokemonSchemaBuilder) SetNickname(nickname string) {
	b.Nickname = &nickname
}
func (b *TrainedPokemonSchemaBuilder) SetDescription(description string) {
	b.Description = &description
}
func (b *TrainedPokemonSchemaBuilder) SetIsPrivate(isPrivate bool) {
	b.IsPrivate = isPrivate
}
func (b *TrainedPokemonSchemaBuilder) SetUserId(id identifier.UserId) {
	*b.CreateUserId = id.Value()
}

func (b TrainedPokemonSchemaBuilder) Build() *schema.TrainedPokemon {
	return b.TrainedPokemon
}
