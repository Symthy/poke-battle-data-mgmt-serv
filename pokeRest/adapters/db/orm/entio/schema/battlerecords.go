package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/entio/property"
)

// BattleRecord holds the schema definition for the BattleRecord entity.
type BattleRecords struct {
	ent.Schema
}

// Fields of the BattleRecord.
func (BattleRecords) Fields() []ent.Field {
	return []ent.Field{
		field.Int("party_id").Positive().Min(0),
		field.String("battle_format").
			GoType(property.BattleFormats("")),
		field.Int("battle_opponent_party_id").
			Positive().Min(0),
		// 自身選出
		field.Int("self_election_pokemon_id1").
			Positive().Min(0),
		field.Int("self_election_pokemon_id2").
			Positive().Min(0),
		field.Int("self_election_pokemon_id3").
			Positive().Min(0),
		field.Int("self_election_pokemon_id4").
			Positive().Min(0).Optional(),
		// 相手選出
		field.Int("opponent_election_pokemon_id1").
			Positive().Min(0),
		field.Int("opponent_election_pokemon_id2").
			Positive().Min(0),
		field.Int("opponent_election_pokemon_id3").
			Positive().Min(0),
		field.Int("opponent_election_pokemon_id4").
			Positive().Min(0),
	}
}

// Edges of the BattleRecord.
func (BattleRecords) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("use_party", Party.Type).
			Field("party_id").
			Required().
			Unique(),
		edge.To("opponent_party", BattleOpponentParty.Type).
			Field("battle_opponent_party_id").
			Required().
			Unique(),
	}
}
