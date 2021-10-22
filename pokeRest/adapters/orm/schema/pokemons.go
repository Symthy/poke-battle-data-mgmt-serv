package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/property"
)

// Pokemons holds the schema definition for the Pokemons entity.
type Pokemons struct {
	ent.Schema
}

// Fields of the Pokemons.
func (Pokemons) Fields() []ent.Field {
	return []ent.Field{
		field.Int("pokedex_no").Positive(),
		field.Int("form_no").Positive(),
		field.String("form_name").NotEmpty(),
		field.String("name").NotEmpty(),
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
	}
}

// Edges of the Pokemons.
func (Pokemons) Edges() []ent.Edge {
	return nil
}
