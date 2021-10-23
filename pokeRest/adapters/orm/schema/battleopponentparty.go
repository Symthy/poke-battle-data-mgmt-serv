package schema

import "entgo.io/ent"

// BattleOpponentParty holds the schema definition for the BattleOpponentParty entity.
type BattleOpponentParty struct {
	ent.Schema
}

// Fields of the BattleOpponentParty.
func (BattleOpponentParty) Fields() []ent.Field {
	return nil
}

// Edges of the BattleOpponentParty.
func (BattleOpponentParty) Edges() []ent.Edge {
	return nil
}
