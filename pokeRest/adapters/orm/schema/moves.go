package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/property"
)

// Moves holds the schema definition for the Moves entity.
type Moves struct {
	ent.Schema
}

// Fields of the Moves.
func (Moves) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Enum("type").GoType(property.Types("")),
		field.String("species").GoType(property.MoveSpecies("")),
		field.Int("power").Positive().Min(0),
		field.Int("accuracy").Positive().Min(0),
		field.Int("pp").Positive().Min(0),
		field.Bool("is_contact").Default(false),
		field.Bool("is_can_guard").Default(false),
	}
}

// Edges of the Moves.
func (Moves) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("to_trained_pokemon_move1", TrainedPokemonDetails.Type).
			Ref("use_move1"),
		edge.From("to_trained_pokemon_move2", TrainedPokemonDetails.Type).
			Ref("use_move2"),
		edge.From("to_trained_pokemon_move3", TrainedPokemonDetails.Type).
			Ref("use_move3"),
		edge.From("to_trained_pokemon_move4", TrainedPokemonDetails.Type).
			Ref("use_move4"),
		edge.To("move_to_pokemon", Pokemons.Type),
	}
}
