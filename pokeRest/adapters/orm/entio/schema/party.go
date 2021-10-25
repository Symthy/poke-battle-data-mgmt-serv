package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/entio/property"
)

// Party holds the schema definition for the Party entity.
type Party struct {
	ent.Schema
}

// Fields of the Party.
func (Party) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Positive().Min(0),
		field.String("name").Optional(),
		field.String("battle_format").
			GoType(property.BattleFormats("")).Optional(),
		field.Int("party_result_id").Positive(),
	}
}

// Edges of the Party.
func (Party) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("party_battle_record", BattleRecords.Type).
			Ref("use_party"),
		edge.From("party_to_tag", Tags.Type).
			Ref("tag_to_party"),
		edge.To("result_record", PartyResultRecord.Type).
			Field("party_result_id"),
	}
}
