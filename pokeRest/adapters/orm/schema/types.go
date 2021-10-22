package schema

import "entgo.io/ent"

// Types holds the schema definition for the Types entity.
type Types struct {
	ent.Schema
}

// Fields of the Types.
func (Types) Fields() []ent.Field {
	return nil
}

// Edges of the Types.
func (Types) Edges() []ent.Edge {
	return nil
}
