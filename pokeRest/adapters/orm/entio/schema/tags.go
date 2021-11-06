package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tags struct {
	ent.Schema
}

// Fields of the Tag.
func (Tags) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Bool("is_generation_tag").Default(false),
		field.Bool("is_season_tag").Default(false),
	}
}

// Edges of the Tag.
func (Tags) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tag_to_party", Party.Type),
	}
}
