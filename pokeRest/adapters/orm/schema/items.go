package schema

import "entgo.io/ent"

// Items holds the schema definition for the Items entity.
type Items struct {
	ent.Schema
}

// Fields of the Items.
func (Items) Fields() []ent.Field {
	return nil
}

// Edges of the Items.
func (Items) Edges() []ent.Edge {
	return nil
}
