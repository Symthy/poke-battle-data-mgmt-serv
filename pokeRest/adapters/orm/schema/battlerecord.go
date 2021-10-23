package schema

import "entgo.io/ent"

// BattleRecord holds the schema definition for the BattleRecord entity.
type BattleRecord struct {
	ent.Schema
}

// Fields of the BattleRecord.
func (BattleRecord) Fields() []ent.Field {
	return nil
}

// Edges of the BattleRecord.
func (BattleRecord) Edges() []ent.Edge {
	return nil
}
