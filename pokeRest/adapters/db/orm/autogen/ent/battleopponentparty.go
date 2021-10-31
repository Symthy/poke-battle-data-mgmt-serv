// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/battleopponentparty"
)

// BattleOpponentParty is the model entity for the BattleOpponentParty schema.
type BattleOpponentParty struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OpponentPokemonId1 holds the value of the "opponent_pokemon_id1" field.
	OpponentPokemonId1 int `json:"opponent_pokemon_id1,omitempty"`
	// OpponentPokemonId2 holds the value of the "opponent_pokemon_id2" field.
	OpponentPokemonId2 int `json:"opponent_pokemon_id2,omitempty"`
	// OpponentPokemonId3 holds the value of the "opponent_pokemon_id3" field.
	OpponentPokemonId3 int `json:"opponent_pokemon_id3,omitempty"`
	// OpponentPokemonId4 holds the value of the "opponent_pokemon_id4" field.
	OpponentPokemonId4 int `json:"opponent_pokemon_id4,omitempty"`
	// OpponentPokemonId5 holds the value of the "opponent_pokemon_id5" field.
	OpponentPokemonId5 int `json:"opponent_pokemon_id5,omitempty"`
	// OpponentPokemonId6 holds the value of the "opponent_pokemon_id6" field.
	OpponentPokemonId6 int `json:"opponent_pokemon_id6,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BattleOpponentPartyQuery when eager-loading is set.
	Edges BattleOpponentPartyEdges `json:"edges"`
}

// BattleOpponentPartyEdges holds the relations/edges for other nodes in the graph.
type BattleOpponentPartyEdges struct {
	// BattleContent holds the value of the battle_content edge.
	BattleContent []*BattleRecords `json:"battle_content,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BattleContentOrErr returns the BattleContent value or an error if the edge
// was not loaded in eager-loading.
func (e BattleOpponentPartyEdges) BattleContentOrErr() ([]*BattleRecords, error) {
	if e.loadedTypes[0] {
		return e.BattleContent, nil
	}
	return nil, &NotLoadedError{edge: "battle_content"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BattleOpponentParty) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case battleopponentparty.FieldID, battleopponentparty.FieldOpponentPokemonId1, battleopponentparty.FieldOpponentPokemonId2, battleopponentparty.FieldOpponentPokemonId3, battleopponentparty.FieldOpponentPokemonId4, battleopponentparty.FieldOpponentPokemonId5, battleopponentparty.FieldOpponentPokemonId6:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BattleOpponentParty", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BattleOpponentParty fields.
func (bop *BattleOpponentParty) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case battleopponentparty.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bop.ID = int(value.Int64)
		case battleopponentparty.FieldOpponentPokemonId1:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field opponent_pokemon_id1", values[i])
			} else if value.Valid {
				bop.OpponentPokemonId1 = int(value.Int64)
			}
		case battleopponentparty.FieldOpponentPokemonId2:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field opponent_pokemon_id2", values[i])
			} else if value.Valid {
				bop.OpponentPokemonId2 = int(value.Int64)
			}
		case battleopponentparty.FieldOpponentPokemonId3:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field opponent_pokemon_id3", values[i])
			} else if value.Valid {
				bop.OpponentPokemonId3 = int(value.Int64)
			}
		case battleopponentparty.FieldOpponentPokemonId4:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field opponent_pokemon_id4", values[i])
			} else if value.Valid {
				bop.OpponentPokemonId4 = int(value.Int64)
			}
		case battleopponentparty.FieldOpponentPokemonId5:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field opponent_pokemon_id5", values[i])
			} else if value.Valid {
				bop.OpponentPokemonId5 = int(value.Int64)
			}
		case battleopponentparty.FieldOpponentPokemonId6:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field opponent_pokemon_id6", values[i])
			} else if value.Valid {
				bop.OpponentPokemonId6 = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryBattleContent queries the "battle_content" edge of the BattleOpponentParty entity.
func (bop *BattleOpponentParty) QueryBattleContent() *BattleRecordsQuery {
	return (&BattleOpponentPartyClient{config: bop.config}).QueryBattleContent(bop)
}

// Update returns a builder for updating this BattleOpponentParty.
// Note that you need to call BattleOpponentParty.Unwrap() before calling this method if this BattleOpponentParty
// was returned from a transaction, and the transaction was committed or rolled back.
func (bop *BattleOpponentParty) Update() *BattleOpponentPartyUpdateOne {
	return (&BattleOpponentPartyClient{config: bop.config}).UpdateOne(bop)
}

// Unwrap unwraps the BattleOpponentParty entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bop *BattleOpponentParty) Unwrap() *BattleOpponentParty {
	tx, ok := bop.config.driver.(*txDriver)
	if !ok {
		panic("ent: BattleOpponentParty is not a transactional entity")
	}
	bop.config.driver = tx.drv
	return bop
}

// String implements the fmt.Stringer.
func (bop *BattleOpponentParty) String() string {
	var builder strings.Builder
	builder.WriteString("BattleOpponentParty(")
	builder.WriteString(fmt.Sprintf("id=%v", bop.ID))
	builder.WriteString(", opponent_pokemon_id1=")
	builder.WriteString(fmt.Sprintf("%v", bop.OpponentPokemonId1))
	builder.WriteString(", opponent_pokemon_id2=")
	builder.WriteString(fmt.Sprintf("%v", bop.OpponentPokemonId2))
	builder.WriteString(", opponent_pokemon_id3=")
	builder.WriteString(fmt.Sprintf("%v", bop.OpponentPokemonId3))
	builder.WriteString(", opponent_pokemon_id4=")
	builder.WriteString(fmt.Sprintf("%v", bop.OpponentPokemonId4))
	builder.WriteString(", opponent_pokemon_id5=")
	builder.WriteString(fmt.Sprintf("%v", bop.OpponentPokemonId5))
	builder.WriteString(", opponent_pokemon_id6=")
	builder.WriteString(fmt.Sprintf("%v", bop.OpponentPokemonId6))
	builder.WriteByte(')')
	return builder.String()
}

// BattleOpponentParties is a parsable slice of BattleOpponentParty.
type BattleOpponentParties []*BattleOpponentParty

func (bop BattleOpponentParties) config(cfg config) {
	for _i := range bop {
		bop[_i].config = cfg
	}
}
