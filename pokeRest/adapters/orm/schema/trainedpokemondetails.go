package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/property"
)

// TrainedPokemonDetails holds the schema definition for the TrainedPokemonDetails entity.
type TrainedPokemonDetails struct {
	ent.Schema
}

// Mixin
func (TrainedPokemonDetails) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the TrainedPokemonDetails.
func (TrainedPokemonDetails) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").Optional(),
		field.String("gender").GoType(property.Gender("")).Optional(),
		field.Int("ability_id").Positive().Min(0),
		field.Int("held_item_id").Positive().Min(0),
		field.Int("move_id1").Positive(),
		field.Int("move_id2").Positive().Optional(),
		field.Int("move_id3").Positive().Optional(),
		field.Int("move_id4").Positive().Optional(),
		field.Int("user_id").Positive().Min(0),
	}
}

// Edges of the TrainedPokemonDetails.
func (TrainedPokemonDetails) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("use_ability", Abilities.Type).
			Field("ability_id").
			Required().
			Unique(),
		edge.To("use_held_item", HeldItems.Type).
			Field("held_item_id").
			Required().
			Unique(),
		edge.To("use_move1", Moves.Type).
			Field("move_id1").
			Required().
			Unique(),
		edge.To("use_move2", Moves.Type).
			Field("move_id2").
			Unique(),
		edge.To("use_move3", Moves.Type).
			Field("move_id3").
			Unique(),
		edge.To("use_move4", Moves.Type).
			Field("move_id4").
			Unique(),
		edge.To("training_detail_user", Users.Type).
			Field("user_id").
			Required().
			Unique(),
	}
}
