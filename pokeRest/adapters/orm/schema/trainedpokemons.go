package schema

import "entgo.io/ent"

// TrainedPokemons holds the schema definition for the TrainedPokemons entity.
type TrainedPokemons struct {
	ent.Schema
}

// Fields of the TrainedPokemons.
func (TrainedPokemons) Fields() []ent.Field {
	return nil
}

// Edges of the TrainedPokemons.
func (TrainedPokemons) Edges() []ent.Edge {
	return nil
}
