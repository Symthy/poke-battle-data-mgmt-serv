package schema

import "entgo.io/ent"

// HeldItems holds the schema definition for the HeldItems entity.
type HeldItems struct {
	ent.Schema
}

// Fields of the HeldItems.
func (HeldItems) Fields() []ent.Field {
	return nil
}

// Edges of the HeldItems.
func (HeldItems) Edges() []ent.Edge {
	return nil
}
