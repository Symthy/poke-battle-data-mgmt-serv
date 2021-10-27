package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/entio/property"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("display_name").Optional(),
		field.String("email").NotEmpty(),
		field.String("role").GoType(property.Role("")),
		field.String("profile").Optional(),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_trained_pokemon", TrainedPokemons.Type).
			Ref("training_user"),
		edge.From("user_trained_pokemon_detail", TrainedPokemonDetails.Type).
			Ref("training_detail_user"),
	}
}
