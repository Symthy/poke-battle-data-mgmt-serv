// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/battleopponentparty"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/battlerecords"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/party"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/property"
)

// BattleRecords is the model entity for the BattleRecords schema.
type BattleRecords struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PartyID holds the value of the "party_id" field.
	PartyID int `json:"party_id,omitempty"`
	// BattleFormat holds the value of the "battle_format" field.
	BattleFormat property.BattleFormats `json:"battle_format,omitempty"`
	// BattleOpponentPartyID holds the value of the "battle_opponent_party_id" field.
	BattleOpponentPartyID int `json:"battle_opponent_party_id,omitempty"`
	// SelfElectionPokemonId1 holds the value of the "self_election_pokemon_id1" field.
	SelfElectionPokemonId1 int `json:"self_election_pokemon_id1,omitempty"`
	// SelfElectionPokemonId2 holds the value of the "self_election_pokemon_id2" field.
	SelfElectionPokemonId2 int `json:"self_election_pokemon_id2,omitempty"`
	// SelfElectionPokemonId3 holds the value of the "self_election_pokemon_id3" field.
	SelfElectionPokemonId3 int `json:"self_election_pokemon_id3,omitempty"`
	// SelfElectionPokemonId4 holds the value of the "self_election_pokemon_id4" field.
	SelfElectionPokemonId4 int `json:"self_election_pokemon_id4,omitempty"`
	// OpponentElectionPokemonId1 holds the value of the "opponent_election_pokemon_id1" field.
	OpponentElectionPokemonId1 int `json:"opponent_election_pokemon_id1,omitempty"`
	// OpponentElectionPokemonId2 holds the value of the "opponent_election_pokemon_id2" field.
	OpponentElectionPokemonId2 int `json:"opponent_election_pokemon_id2,omitempty"`
	// OpponentElectionPokemonId3 holds the value of the "opponent_election_pokemon_id3" field.
	OpponentElectionPokemonId3 int `json:"opponent_election_pokemon_id3,omitempty"`
	// OpponentElectionPokemonId4 holds the value of the "opponent_election_pokemon_id4" field.
	OpponentElectionPokemonId4 int `json:"opponent_election_pokemon_id4,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BattleRecordsQuery when eager-loading is set.
	Edges BattleRecordsEdges `json:"edges"`
}

// BattleRecordsEdges holds the relations/edges for other nodes in the graph.
type BattleRecordsEdges struct {
	// UseParty holds the value of the use_party edge.
	UseParty *Party `json:"use_party,omitempty"`
	// OpponentParty holds the value of the opponent_party edge.
	OpponentParty *BattleOpponentParty `json:"opponent_party,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsePartyOrErr returns the UseParty value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BattleRecordsEdges) UsePartyOrErr() (*Party, error) {
	if e.loadedTypes[0] {
		if e.UseParty == nil {
			// The edge use_party was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: party.Label}
		}
		return e.UseParty, nil
	}
	return nil, &NotLoadedError{edge: "use_party"}
}

// OpponentPartyOrErr returns the OpponentParty value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BattleRecordsEdges) OpponentPartyOrErr() (*BattleOpponentParty, error) {
	if e.loadedTypes[1] {
		if e.OpponentParty == nil {
			// The edge opponent_party was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: battleopponentparty.Label}
		}
		return e.OpponentParty, nil
	}
	return nil, &NotLoadedError{edge: "opponent_party"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BattleRecords) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case battlerecords.FieldID, battlerecords.FieldPartyID, battlerecords.FieldBattleOpponentPartyID, battlerecords.FieldSelfElectionPokemonId1, battlerecords.FieldSelfElectionPokemonId2, battlerecords.FieldSelfElectionPokemonId3, battlerecords.FieldSelfElectionPokemonId4, battlerecords.FieldOpponentElectionPokemonId1, battlerecords.FieldOpponentElectionPokemonId2, battlerecords.FieldOpponentElectionPokemonId3, battlerecords.FieldOpponentElectionPokemonId4:
			values[i] = new(sql.NullInt64)
		case battlerecords.FieldBattleFormat:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BattleRecords", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BattleRecords fields.
func (br *BattleRecords) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case battlerecords.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			br.ID = int(value.Int64)
		case battlerecords.FieldPartyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field party_id", values[i])
			} else if value.Valid {
				br.PartyID = int(value.Int64)
			}
		case battlerecords.FieldBattleFormat:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field battle_format", values[i])
			} else if value.Valid {
				br.BattleFormat = property.BattleFormats(value.String)
			}
		case battlerecords.FieldBattleOpponentPartyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field battle_opponent_party_id", values[i])
			} else if value.Valid {
				br.BattleOpponentPartyID = int(value.Int64)
			}
		case battlerecords.FieldSelfElectionPokemonId1:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field self_election_pokemon_id1", values[i])
			} else if value.Valid {
				br.SelfElectionPokemonId1 = int(value.Int64)
			}
		case battlerecords.FieldSelfElectionPokemonId2:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field self_election_pokemon_id2", values[i])
			} else if value.Valid {
				br.SelfElectionPokemonId2 = int(value.Int64)
			}
		case battlerecords.FieldSelfElectionPokemonId3:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field self_election_pokemon_id3", values[i])
			} else if value.Valid {
				br.SelfElectionPokemonId3 = int(value.Int64)
			}
		case battlerecords.FieldSelfElectionPokemonId4:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field self_election_pokemon_id4", values[i])
			} else if value.Valid {
				br.SelfElectionPokemonId4 = int(value.Int64)
			}
		case battlerecords.FieldOpponentElectionPokemonId1:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field opponent_election_pokemon_id1", values[i])
			} else if value.Valid {
				br.OpponentElectionPokemonId1 = int(value.Int64)
			}
		case battlerecords.FieldOpponentElectionPokemonId2:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field opponent_election_pokemon_id2", values[i])
			} else if value.Valid {
				br.OpponentElectionPokemonId2 = int(value.Int64)
			}
		case battlerecords.FieldOpponentElectionPokemonId3:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field opponent_election_pokemon_id3", values[i])
			} else if value.Valid {
				br.OpponentElectionPokemonId3 = int(value.Int64)
			}
		case battlerecords.FieldOpponentElectionPokemonId4:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field opponent_election_pokemon_id4", values[i])
			} else if value.Valid {
				br.OpponentElectionPokemonId4 = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUseParty queries the "use_party" edge of the BattleRecords entity.
func (br *BattleRecords) QueryUseParty() *PartyQuery {
	return (&BattleRecordsClient{config: br.config}).QueryUseParty(br)
}

// QueryOpponentParty queries the "opponent_party" edge of the BattleRecords entity.
func (br *BattleRecords) QueryOpponentParty() *BattleOpponentPartyQuery {
	return (&BattleRecordsClient{config: br.config}).QueryOpponentParty(br)
}

// Update returns a builder for updating this BattleRecords.
// Note that you need to call BattleRecords.Unwrap() before calling this method if this BattleRecords
// was returned from a transaction, and the transaction was committed or rolled back.
func (br *BattleRecords) Update() *BattleRecordsUpdateOne {
	return (&BattleRecordsClient{config: br.config}).UpdateOne(br)
}

// Unwrap unwraps the BattleRecords entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (br *BattleRecords) Unwrap() *BattleRecords {
	tx, ok := br.config.driver.(*txDriver)
	if !ok {
		panic("ent: BattleRecords is not a transactional entity")
	}
	br.config.driver = tx.drv
	return br
}

// String implements the fmt.Stringer.
func (br *BattleRecords) String() string {
	var builder strings.Builder
	builder.WriteString("BattleRecords(")
	builder.WriteString(fmt.Sprintf("id=%v", br.ID))
	builder.WriteString(", party_id=")
	builder.WriteString(fmt.Sprintf("%v", br.PartyID))
	builder.WriteString(", battle_format=")
	builder.WriteString(fmt.Sprintf("%v", br.BattleFormat))
	builder.WriteString(", battle_opponent_party_id=")
	builder.WriteString(fmt.Sprintf("%v", br.BattleOpponentPartyID))
	builder.WriteString(", self_election_pokemon_id1=")
	builder.WriteString(fmt.Sprintf("%v", br.SelfElectionPokemonId1))
	builder.WriteString(", self_election_pokemon_id2=")
	builder.WriteString(fmt.Sprintf("%v", br.SelfElectionPokemonId2))
	builder.WriteString(", self_election_pokemon_id3=")
	builder.WriteString(fmt.Sprintf("%v", br.SelfElectionPokemonId3))
	builder.WriteString(", self_election_pokemon_id4=")
	builder.WriteString(fmt.Sprintf("%v", br.SelfElectionPokemonId4))
	builder.WriteString(", opponent_election_pokemon_id1=")
	builder.WriteString(fmt.Sprintf("%v", br.OpponentElectionPokemonId1))
	builder.WriteString(", opponent_election_pokemon_id2=")
	builder.WriteString(fmt.Sprintf("%v", br.OpponentElectionPokemonId2))
	builder.WriteString(", opponent_election_pokemon_id3=")
	builder.WriteString(fmt.Sprintf("%v", br.OpponentElectionPokemonId3))
	builder.WriteString(", opponent_election_pokemon_id4=")
	builder.WriteString(fmt.Sprintf("%v", br.OpponentElectionPokemonId4))
	builder.WriteByte(')')
	return builder.String()
}

// BattleRecordsSlice is a parsable slice of BattleRecords.
type BattleRecordsSlice []*BattleRecords

func (br BattleRecordsSlice) config(cfg config) {
	for _i := range br {
		br[_i].config = cfg
	}
}