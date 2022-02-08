package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/entio/property"
)

// Pokemons holds the schema definition for the Pokemons entity.
type Pokemons struct {
	ent.Schema
}

// Fields of the Pokemons.
func (Pokemons) Fields() []ent.Field {
	return []ent.Field{
		field.Int("pokedex_no").Positive().Min(0),
		field.Int("form_no").Positive().Min(0),
		field.String("form_name").NotEmpty(),
		field.String("name").NotEmpty(),
		field.String("english_name").Match(regexp.MustCompile("^[A-Z].*$")),
		field.Enum("type1").GoType(property.Types("")),
		field.Enum("type2").GoType(property.Types("")).Optional(),
		field.Int("ability_id1").Positive(),
		field.Int("ability_id2").Positive().Optional(),
		field.Int("hidden_ability_id").Positive().Optional(),
		field.Bool("is_final_evolution").Default(false),
		field.Int("base_stats_h").Positive().Default(0),
		field.Int("base_stats_a").Positive().Default(0),
		field.Int("base_stats_b").Positive().Default(0),
		field.Int("base_stats_c").Positive().Default(0),
		field.Int("base_stats_d").Positive().Default(0),
		field.Int("base_stats_s").Positive().Default(0),
		field.Int("generation").Positive().Min(0),
	}
}

// Edges of the Pokemons.
func (Pokemons) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ability1", Abilities.Type).
			Field("ability_id1").
			Required().
			Unique(),
		edge.To("ability2", Abilities.Type).
			Field("ability_id2").
			Unique(),
		edge.To("hidden_ability", Abilities.Type).
			Field("hidden_ability_id").
			Unique(),
		edge.To("form", Forms.Type).
			Field("form_no").
			Required().
			Unique(),
		edge.From("to_trained_pokemon", TrainedPokemons.Type).
			Ref("use_pokemon"),
		edge.From("pokemon_to_move", Moves.Type).
			Ref("move_to_pokemon"),
	}
}
