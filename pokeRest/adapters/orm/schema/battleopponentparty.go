package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BattleOpponentParty holds the schema definition for the BattleOpponentParty entity.
type BattleOpponentParty struct {
	ent.Schema
}

// Fields of the BattleOpponentParty.
func (BattleOpponentParty) Fields() []ent.Field {
	return []ent.Field{
		field.Int("opponent_pokemon_id1").
			Positive().Min(0),
		field.Int("opponent_pokemon_id2").
			Positive().Min(0),
		field.Int("opponent_pokemon_id3").
			Positive().Min(0),
		field.Int("opponent_pokemon_id4").
			Positive().Min(0).Optional(),
		field.Int("opponent_pokemon_id5").
			Positive().Min(0).Optional(),
		field.Int("opponent_pokemon_id6").
			Positive().Min(0).Optional(),
	}
}

// Edges of the BattleOpponentParty.
func (BattleOpponentParty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("battle_content", BattleRecords.Type).
			Ref("opponent_party"),
	}
}
