// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/tags"
)

// Tags is the model entity for the Tags schema.
type Tags struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsGenerationTag holds the value of the "is_generation_tag" field.
	IsGenerationTag bool `json:"is_generation_tag,omitempty"`
	// IsSeasonTag holds the value of the "is_season_tag" field.
	IsSeasonTag bool `json:"is_season_tag,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TagsQuery when eager-loading is set.
	Edges TagsEdges `json:"edges"`
}

// TagsEdges holds the relations/edges for other nodes in the graph.
type TagsEdges struct {
	// TagToParty holds the value of the tag_to_party edge.
	TagToParty []*Party `json:"tag_to_party,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TagToPartyOrErr returns the TagToParty value or an error if the edge
// was not loaded in eager-loading.
func (e TagsEdges) TagToPartyOrErr() ([]*Party, error) {
	if e.loadedTypes[0] {
		return e.TagToParty, nil
	}
	return nil, &NotLoadedError{edge: "tag_to_party"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tags) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tags.FieldIsGenerationTag, tags.FieldIsSeasonTag:
			values[i] = new(sql.NullBool)
		case tags.FieldID:
			values[i] = new(sql.NullInt64)
		case tags.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tags", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tags fields.
func (t *Tags) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tags.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tags.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tags.FieldIsGenerationTag:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_generation_tag", values[i])
			} else if value.Valid {
				t.IsGenerationTag = value.Bool
			}
		case tags.FieldIsSeasonTag:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_season_tag", values[i])
			} else if value.Valid {
				t.IsSeasonTag = value.Bool
			}
		}
	}
	return nil
}

// QueryTagToParty queries the "tag_to_party" edge of the Tags entity.
func (t *Tags) QueryTagToParty() *PartyQuery {
	return (&TagsClient{config: t.config}).QueryTagToParty(t)
}

// Update returns a builder for updating this Tags.
// Note that you need to call Tags.Unwrap() before calling this method if this Tags
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tags) Update() *TagsUpdateOne {
	return (&TagsClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Tags entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tags) Unwrap() *Tags {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tags is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tags) String() string {
	var builder strings.Builder
	builder.WriteString("Tags(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", is_generation_tag=")
	builder.WriteString(fmt.Sprintf("%v", t.IsGenerationTag))
	builder.WriteString(", is_season_tag=")
	builder.WriteString(fmt.Sprintf("%v", t.IsSeasonTag))
	builder.WriteByte(')')
	return builder.String()
}

// TagsSlice is a parsable slice of Tags.
type TagsSlice []*Tags

func (t TagsSlice) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}