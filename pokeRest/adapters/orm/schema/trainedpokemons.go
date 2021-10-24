package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/property"
)

// TrainedPokemons holds the schema definition for the TrainedPokemons entity.
type TrainedPokemons struct {
	ent.Schema
}

// Mixin
func (TrainedPokemons) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the TrainedPokemons.
func (TrainedPokemons) Fields() []ent.Field {
	return []ent.Field{
		field.Int("pokemon_id").Positive().Min(0),
		field.Int("create_user_id").Positive().Min(0).Optional(),
		field.String("nature").GoType(property.Nature("")).NotEmpty(),
		field.Int("effort_value_h").Default(0),
		field.Int("effort_value_a").Default(0),
		field.Int("effort_value_b").Default(0),
		field.Int("effort_value_c").Default(0),
		field.Int("effort_value_d").Default(0),
		field.Int("effort_value_s").Default(0),
	}
}

// Edges of the TrainedPokemons.
func (TrainedPokemons) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("use_pokemon", Pokemons.Type).
			Field("pokemon_id").
			Required().
			Unique(),
		edge.To("training_user", Users.Type).
			Field("create_user_id").
			Unique(),
	}
}
