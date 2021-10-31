package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/entio/schema/common"
)

// HeldItems holds the schema definition for the HeldItems entity.
type HeldItems struct {
	ent.Schema
}

// Mixin
func (HeldItems) Mixin() []ent.Mixin {
	return []ent.Mixin{
		common.CorrectionValueMixin{},
	}
}

// Fields of the HeldItems.
func (HeldItems) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
	}
}

// Edges of the HeldItems.
func (HeldItems) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("to_trained_pokemon_item", TrainedPokemonDetails.Type).
			Ref("use_held_item"),
	}
}
