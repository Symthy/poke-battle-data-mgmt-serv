package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Forms holds the schema definition for the Forms entity.
type Forms struct {
	ent.Schema
}

// Fields of the Forms.
func (Forms) Fields() []ent.Field {
	return []ent.Field{
		field.String("form_name").NotEmpty(),
		field.Bool("is_regional_variant").Default(false),
		field.String("regionName").NotEmpty(),
	}
}

// Edges of the Forms.
func (Forms) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("form_holder", Pokemons.Type).
			Ref("form"),
	}
}
