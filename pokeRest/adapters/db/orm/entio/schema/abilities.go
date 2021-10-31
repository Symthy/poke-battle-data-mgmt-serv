package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/entio/schema/common"
)

// Abilities holds the schema definition for the Abilities entity.
type Abilities struct {
	ent.Schema
}

// Mixin
func (Abilities) Mixin() []ent.Mixin {
	return []ent.Mixin{
		common.CorrectionValueMixin{},
	}
}

// Fields of the Abilities.
func (Abilities) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
	}
}

// Edges of the Abilities.
func (Abilities) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ability_holder1", Pokemons.Type).
			Ref("ability1"),
		edge.From("ability_holder2", Pokemons.Type).
			Ref("ability2"),
		edge.From("hidden_ability_holder", Pokemons.Type).
			Ref("hidden_ability"),
		edge.From("to_trained_pokemon_ability", TrainedPokemonDetails.Type).
			Ref("use_ability"),
	}
}
