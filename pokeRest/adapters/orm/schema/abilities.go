package schema

import "entgo.io/ent"

// Abilities holds the schema definition for the Abilities entity.
type Abilities struct {
	ent.Schema
}

// Fields of the Abilities.
func (Abilities) Fields() []ent.Field {
	return nil
}

// Edges of the Abilities.
func (Abilities) Edges() []ent.Edge {
	return nil
}
