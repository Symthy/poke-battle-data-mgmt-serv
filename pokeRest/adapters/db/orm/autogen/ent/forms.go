// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/forms"
)

// Forms is the model entity for the Forms schema.
type Forms struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FormName holds the value of the "form_name" field.
	FormName string `json:"form_name,omitempty"`
	// IsRegionalVariant holds the value of the "is_regional_variant" field.
	IsRegionalVariant bool `json:"is_regional_variant,omitempty"`
	// RegionName holds the value of the "regionName" field.
	RegionName string `json:"regionName,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FormsQuery when eager-loading is set.
	Edges FormsEdges `json:"edges"`
}

// FormsEdges holds the relations/edges for other nodes in the graph.
type FormsEdges struct {
	// FormHolder holds the value of the form_holder edge.
	FormHolder []*Pokemons `json:"form_holder,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FormHolderOrErr returns the FormHolder value or an error if the edge
// was not loaded in eager-loading.
func (e FormsEdges) FormHolderOrErr() ([]*Pokemons, error) {
	if e.loadedTypes[0] {
		return e.FormHolder, nil
	}
	return nil, &NotLoadedError{edge: "form_holder"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Forms) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case forms.FieldIsRegionalVariant:
			values[i] = new(sql.NullBool)
		case forms.FieldID:
			values[i] = new(sql.NullInt64)
		case forms.FieldFormName, forms.FieldRegionName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Forms", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Forms fields.
func (f *Forms) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case forms.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case forms.FieldFormName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field form_name", values[i])
			} else if value.Valid {
				f.FormName = value.String
			}
		case forms.FieldIsRegionalVariant:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_regional_variant", values[i])
			} else if value.Valid {
				f.IsRegionalVariant = value.Bool
			}
		case forms.FieldRegionName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field regionName", values[i])
			} else if value.Valid {
				f.RegionName = value.String
			}
		}
	}
	return nil
}

// QueryFormHolder queries the "form_holder" edge of the Forms entity.
func (f *Forms) QueryFormHolder() *PokemonsQuery {
	return (&FormsClient{config: f.config}).QueryFormHolder(f)
}

// Update returns a builder for updating this Forms.
// Note that you need to call Forms.Unwrap() before calling this method if this Forms
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Forms) Update() *FormsUpdateOne {
	return (&FormsClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Forms entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Forms) Unwrap() *Forms {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Forms is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Forms) String() string {
	var builder strings.Builder
	builder.WriteString("Forms(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", form_name=")
	builder.WriteString(f.FormName)
	builder.WriteString(", is_regional_variant=")
	builder.WriteString(fmt.Sprintf("%v", f.IsRegionalVariant))
	builder.WriteString(", regionName=")
	builder.WriteString(f.RegionName)
	builder.WriteByte(')')
	return builder.String()
}

// FormsSlice is a parsable slice of Forms.
type FormsSlice []*Forms

func (f FormsSlice) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}