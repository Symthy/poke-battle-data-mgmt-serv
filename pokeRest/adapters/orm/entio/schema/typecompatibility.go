package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/entio/property"
)

// TypeCompatibility holds the schema definition for the TypeCompatibility entity.
type TypeCompatibility struct {
	ent.Schema
}

// Fields of the TypeCompatibility.
func (TypeCompatibility) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("attack_type").GoType(property.Types("")),
		field.Enum("defence_type").GoType(property.Types("")),
		field.Int("compatibility").Default(1), // 相性
	}
}

// Edges of the TypeCompatibility.
func (TypeCompatibility) Edges() []ent.Edge {
	return nil
}
