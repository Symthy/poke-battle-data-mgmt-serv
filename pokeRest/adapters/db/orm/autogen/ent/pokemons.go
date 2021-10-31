// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/abilities"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/forms"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/property"
)

// Pokemons is the model entity for the Pokemons schema.
type Pokemons struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PokedexNo holds the value of the "pokedex_no" field.
	PokedexNo int `json:"pokedex_no,omitempty"`
	// FormNo holds the value of the "form_no" field.
	FormNo int `json:"form_no,omitempty"`
	// FormName holds the value of the "form_name" field.
	FormName string `json:"form_name,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// EnglishName holds the value of the "english_name" field.
	EnglishName string `json:"english_name,omitempty"`
	// Type1 holds the value of the "type1" field.
	Type1 property.Types `json:"type1,omitempty"`
	// Type2 holds the value of the "type2" field.
	Type2 property.Types `json:"type2,omitempty"`
	// AbilityId1 holds the value of the "ability_id1" field.
	AbilityId1 int `json:"ability_id1,omitempty"`
	// AbilityId2 holds the value of the "ability_id2" field.
	AbilityId2 int `json:"ability_id2,omitempty"`
	// HiddenAbilityID holds the value of the "hidden_ability_id" field.
	HiddenAbilityID int `json:"hidden_ability_id,omitempty"`
	// IsFinalEvolution holds the value of the "is_final_evolution" field.
	IsFinalEvolution bool `json:"is_final_evolution,omitempty"`
	// BaseStatsH holds the value of the "base_stats_h" field.
	BaseStatsH int `json:"base_stats_h,omitempty"`
	// BaseStatsA holds the value of the "base_stats_a" field.
	BaseStatsA int `json:"base_stats_a,omitempty"`
	// BaseStatsB holds the value of the "base_stats_b" field.
	BaseStatsB int `json:"base_stats_b,omitempty"`
	// BaseStatsC holds the value of the "base_stats_c" field.
	BaseStatsC int `json:"base_stats_c,omitempty"`
	// BaseStatsD holds the value of the "base_stats_d" field.
	BaseStatsD int `json:"base_stats_d,omitempty"`
	// BaseStatsS holds the value of the "base_stats_s" field.
	BaseStatsS int `json:"base_stats_s,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PokemonsQuery when eager-loading is set.
	Edges PokemonsEdges `json:"edges"`
}

// PokemonsEdges holds the relations/edges for other nodes in the graph.
type PokemonsEdges struct {
	// Ability1 holds the value of the ability1 edge.
	Ability1 *Abilities `json:"ability1,omitempty"`
	// Ability2 holds the value of the ability2 edge.
	Ability2 *Abilities `json:"ability2,omitempty"`
	// HiddenAbility holds the value of the hidden_ability edge.
	HiddenAbility *Abilities `json:"hidden_ability,omitempty"`
	// Form holds the value of the form edge.
	Form *Forms `json:"form,omitempty"`
	// ToTrainedPokemon holds the value of the to_trained_pokemon edge.
	ToTrainedPokemon []*TrainedPokemons `json:"to_trained_pokemon,omitempty"`
	// PokemonToMove holds the value of the pokemon_to_move edge.
	PokemonToMove []*Moves `json:"pokemon_to_move,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// Ability1OrErr returns the Ability1 value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PokemonsEdges) Ability1OrErr() (*Abilities, error) {
	if e.loadedTypes[0] {
		if e.Ability1 == nil {
			// The edge ability1 was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: abilities.Label}
		}
		return e.Ability1, nil
	}
	return nil, &NotLoadedError{edge: "ability1"}
}

// Ability2OrErr returns the Ability2 value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PokemonsEdges) Ability2OrErr() (*Abilities, error) {
	if e.loadedTypes[1] {
		if e.Ability2 == nil {
			// The edge ability2 was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: abilities.Label}
		}
		return e.Ability2, nil
	}
	return nil, &NotLoadedError{edge: "ability2"}
}

// HiddenAbilityOrErr returns the HiddenAbility value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PokemonsEdges) HiddenAbilityOrErr() (*Abilities, error) {
	if e.loadedTypes[2] {
		if e.HiddenAbility == nil {
			// The edge hidden_ability was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: abilities.Label}
		}
		return e.HiddenAbility, nil
	}
	return nil, &NotLoadedError{edge: "hidden_ability"}
}

// FormOrErr returns the Form value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PokemonsEdges) FormOrErr() (*Forms, error) {
	if e.loadedTypes[3] {
		if e.Form == nil {
			// The edge form was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: forms.Label}
		}
		return e.Form, nil
	}
	return nil, &NotLoadedError{edge: "form"}
}

// ToTrainedPokemonOrErr returns the ToTrainedPokemon value or an error if the edge
// was not loaded in eager-loading.
func (e PokemonsEdges) ToTrainedPokemonOrErr() ([]*TrainedPokemons, error) {
	if e.loadedTypes[4] {
		return e.ToTrainedPokemon, nil
	}
	return nil, &NotLoadedError{edge: "to_trained_pokemon"}
}

// PokemonToMoveOrErr returns the PokemonToMove value or an error if the edge
// was not loaded in eager-loading.
func (e PokemonsEdges) PokemonToMoveOrErr() ([]*Moves, error) {
	if e.loadedTypes[5] {
		return e.PokemonToMove, nil
	}
	return nil, &NotLoadedError{edge: "pokemon_to_move"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pokemons) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pokemons.FieldIsFinalEvolution:
			values[i] = new(sql.NullBool)
		case pokemons.FieldID, pokemons.FieldPokedexNo, pokemons.FieldFormNo, pokemons.FieldAbilityId1, pokemons.FieldAbilityId2, pokemons.FieldHiddenAbilityID, pokemons.FieldBaseStatsH, pokemons.FieldBaseStatsA, pokemons.FieldBaseStatsB, pokemons.FieldBaseStatsC, pokemons.FieldBaseStatsD, pokemons.FieldBaseStatsS:
			values[i] = new(sql.NullInt64)
		case pokemons.FieldFormName, pokemons.FieldName, pokemons.FieldEnglishName, pokemons.FieldType1, pokemons.FieldType2:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pokemons", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pokemons fields.
func (po *Pokemons) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pokemons.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case pokemons.FieldPokedexNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pokedex_no", values[i])
			} else if value.Valid {
				po.PokedexNo = int(value.Int64)
			}
		case pokemons.FieldFormNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field form_no", values[i])
			} else if value.Valid {
				po.FormNo = int(value.Int64)
			}
		case pokemons.FieldFormName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field form_name", values[i])
			} else if value.Valid {
				po.FormName = value.String
			}
		case pokemons.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				po.Name = value.String
			}
		case pokemons.FieldEnglishName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field english_name", values[i])
			} else if value.Valid {
				po.EnglishName = value.String
			}
		case pokemons.FieldType1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type1", values[i])
			} else if value.Valid {
				po.Type1 = property.Types(value.String)
			}
		case pokemons.FieldType2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type2", values[i])
			} else if value.Valid {
				po.Type2 = property.Types(value.String)
			}
		case pokemons.FieldAbilityId1:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ability_id1", values[i])
			} else if value.Valid {
				po.AbilityId1 = int(value.Int64)
			}
		case pokemons.FieldAbilityId2:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ability_id2", values[i])
			} else if value.Valid {
				po.AbilityId2 = int(value.Int64)
			}
		case pokemons.FieldHiddenAbilityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hidden_ability_id", values[i])
			} else if value.Valid {
				po.HiddenAbilityID = int(value.Int64)
			}
		case pokemons.FieldIsFinalEvolution:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_final_evolution", values[i])
			} else if value.Valid {
				po.IsFinalEvolution = value.Bool
			}
		case pokemons.FieldBaseStatsH:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field base_stats_h", values[i])
			} else if value.Valid {
				po.BaseStatsH = int(value.Int64)
			}
		case pokemons.FieldBaseStatsA:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field base_stats_a", values[i])
			} else if value.Valid {
				po.BaseStatsA = int(value.Int64)
			}
		case pokemons.FieldBaseStatsB:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field base_stats_b", values[i])
			} else if value.Valid {
				po.BaseStatsB = int(value.Int64)
			}
		case pokemons.FieldBaseStatsC:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field base_stats_c", values[i])
			} else if value.Valid {
				po.BaseStatsC = int(value.Int64)
			}
		case pokemons.FieldBaseStatsD:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field base_stats_d", values[i])
			} else if value.Valid {
				po.BaseStatsD = int(value.Int64)
			}
		case pokemons.FieldBaseStatsS:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field base_stats_s", values[i])
			} else if value.Valid {
				po.BaseStatsS = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAbility1 queries the "ability1" edge of the Pokemons entity.
func (po *Pokemons) QueryAbility1() *AbilitiesQuery {
	return (&PokemonsClient{config: po.config}).QueryAbility1(po)
}

// QueryAbility2 queries the "ability2" edge of the Pokemons entity.
func (po *Pokemons) QueryAbility2() *AbilitiesQuery {
	return (&PokemonsClient{config: po.config}).QueryAbility2(po)
}

// QueryHiddenAbility queries the "hidden_ability" edge of the Pokemons entity.
func (po *Pokemons) QueryHiddenAbility() *AbilitiesQuery {
	return (&PokemonsClient{config: po.config}).QueryHiddenAbility(po)
}

// QueryForm queries the "form" edge of the Pokemons entity.
func (po *Pokemons) QueryForm() *FormsQuery {
	return (&PokemonsClient{config: po.config}).QueryForm(po)
}

// QueryToTrainedPokemon queries the "to_trained_pokemon" edge of the Pokemons entity.
func (po *Pokemons) QueryToTrainedPokemon() *TrainedPokemonsQuery {
	return (&PokemonsClient{config: po.config}).QueryToTrainedPokemon(po)
}

// QueryPokemonToMove queries the "pokemon_to_move" edge of the Pokemons entity.
func (po *Pokemons) QueryPokemonToMove() *MovesQuery {
	return (&PokemonsClient{config: po.config}).QueryPokemonToMove(po)
}

// Update returns a builder for updating this Pokemons.
// Note that you need to call Pokemons.Unwrap() before calling this method if this Pokemons
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Pokemons) Update() *PokemonsUpdateOne {
	return (&PokemonsClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Pokemons entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Pokemons) Unwrap() *Pokemons {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pokemons is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Pokemons) String() string {
	var builder strings.Builder
	builder.WriteString("Pokemons(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", pokedex_no=")
	builder.WriteString(fmt.Sprintf("%v", po.PokedexNo))
	builder.WriteString(", form_no=")
	builder.WriteString(fmt.Sprintf("%v", po.FormNo))
	builder.WriteString(", form_name=")
	builder.WriteString(po.FormName)
	builder.WriteString(", name=")
	builder.WriteString(po.Name)
	builder.WriteString(", english_name=")
	builder.WriteString(po.EnglishName)
	builder.WriteString(", type1=")
	builder.WriteString(fmt.Sprintf("%v", po.Type1))
	builder.WriteString(", type2=")
	builder.WriteString(fmt.Sprintf("%v", po.Type2))
	builder.WriteString(", ability_id1=")
	builder.WriteString(fmt.Sprintf("%v", po.AbilityId1))
	builder.WriteString(", ability_id2=")
	builder.WriteString(fmt.Sprintf("%v", po.AbilityId2))
	builder.WriteString(", hidden_ability_id=")
	builder.WriteString(fmt.Sprintf("%v", po.HiddenAbilityID))
	builder.WriteString(", is_final_evolution=")
	builder.WriteString(fmt.Sprintf("%v", po.IsFinalEvolution))
	builder.WriteString(", base_stats_h=")
	builder.WriteString(fmt.Sprintf("%v", po.BaseStatsH))
	builder.WriteString(", base_stats_a=")
	builder.WriteString(fmt.Sprintf("%v", po.BaseStatsA))
	builder.WriteString(", base_stats_b=")
	builder.WriteString(fmt.Sprintf("%v", po.BaseStatsB))
	builder.WriteString(", base_stats_c=")
	builder.WriteString(fmt.Sprintf("%v", po.BaseStatsC))
	builder.WriteString(", base_stats_d=")
	builder.WriteString(fmt.Sprintf("%v", po.BaseStatsD))
	builder.WriteString(", base_stats_s=")
	builder.WriteString(fmt.Sprintf("%v", po.BaseStatsS))
	builder.WriteByte(')')
	return builder.String()
}

// PokemonsSlice is a parsable slice of Pokemons.
type PokemonsSlice []*Pokemons

func (po PokemonsSlice) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
