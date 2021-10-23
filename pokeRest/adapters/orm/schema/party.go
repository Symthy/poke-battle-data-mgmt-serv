package schema

import "entgo.io/ent"

// Party holds the schema definition for the Party entity.
type Party struct {
	ent.Schema
}

// Fields of the Party.
func (Party) Fields() []ent.Field {
	return nil
}

// Edges of the Party.
func (Party) Edges() []ent.Edge {
	return nil
}
