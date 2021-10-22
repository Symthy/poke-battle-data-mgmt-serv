package schema

import "entgo.io/ent"

// Forms holds the schema definition for the Forms entity.
type Forms struct {
	ent.Schema
}

// Fields of the Forms.
func (Forms) Fields() []ent.Field {
	return nil
}

// Edges of the Forms.
func (Forms) Edges() []ent.Edge {
	return nil
}
