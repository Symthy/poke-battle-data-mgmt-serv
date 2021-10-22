package schema

import "entgo.io/ent"

// Moves holds the schema definition for the Moves entity.
type Moves struct {
	ent.Schema
}

// Fields of the Moves.
func (Moves) Fields() []ent.Field {
	return nil
}

// Edges of the Moves.
func (Moves) Edges() []ent.Edge {
	return nil
}
