package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PartyResultRecord holds the schema definition for the PartyResultRecord entity.
type PartyResultRecord struct {
	ent.Schema
}

// Fields of the PartyResultRecord.
func (PartyResultRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("party_id").Positive(),
		field.Int("generation").Positive(),
		field.Int("seasen").Positive(),
		field.Int("max_rate").Positive().Optional(),
		field.Int("max_ranking").Positive().Optional(),
	}
}

// Edges of the PartyResultRecord.
func (PartyResultRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("result_to_party", Party.Type).
			Field("party_id").
			Required().
			Unique(),
	}
}
