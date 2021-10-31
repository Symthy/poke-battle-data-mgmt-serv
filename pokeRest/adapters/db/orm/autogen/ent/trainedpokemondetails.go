// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/abilities"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/helditems"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/moves"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/trainedpokemondetails"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/users"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/property"
)

// TrainedPokemonDetails is the model entity for the TrainedPokemonDetails schema.
type TrainedPokemonDetails struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender property.Gender `json:"gender,omitempty"`
	// AbilityID holds the value of the "ability_id" field.
	AbilityID int `json:"ability_id,omitempty"`
	// HeldItemID holds the value of the "held_item_id" field.
	HeldItemID int `json:"held_item_id,omitempty"`
	// MoveId1 holds the value of the "move_id1" field.
	MoveId1 int `json:"move_id1,omitempty"`
	// MoveId2 holds the value of the "move_id2" field.
	MoveId2 int `json:"move_id2,omitempty"`
	// MoveId3 holds the value of the "move_id3" field.
	MoveId3 int `json:"move_id3,omitempty"`
	// MoveId4 holds the value of the "move_id4" field.
	MoveId4 int `json:"move_id4,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TrainedPokemonDetailsQuery when eager-loading is set.
	Edges TrainedPokemonDetailsEdges `json:"edges"`
}

// TrainedPokemonDetailsEdges holds the relations/edges for other nodes in the graph.
type TrainedPokemonDetailsEdges struct {
	// UseAbility holds the value of the use_ability edge.
	UseAbility *Abilities `json:"use_ability,omitempty"`
	// UseHeldItem holds the value of the use_held_item edge.
	UseHeldItem *HeldItems `json:"use_held_item,omitempty"`
	// UseMove1 holds the value of the use_move1 edge.
	UseMove1 *Moves `json:"use_move1,omitempty"`
	// UseMove2 holds the value of the use_move2 edge.
	UseMove2 *Moves `json:"use_move2,omitempty"`
	// UseMove3 holds the value of the use_move3 edge.
	UseMove3 *Moves `json:"use_move3,omitempty"`
	// UseMove4 holds the value of the use_move4 edge.
	UseMove4 *Moves `json:"use_move4,omitempty"`
	// TrainingDetailUser holds the value of the training_detail_user edge.
	TrainingDetailUser *Users `json:"training_detail_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// UseAbilityOrErr returns the UseAbility value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrainedPokemonDetailsEdges) UseAbilityOrErr() (*Abilities, error) {
	if e.loadedTypes[0] {
		if e.UseAbility == nil {
			// The edge use_ability was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: abilities.Label}
		}
		return e.UseAbility, nil
	}
	return nil, &NotLoadedError{edge: "use_ability"}
}

// UseHeldItemOrErr returns the UseHeldItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrainedPokemonDetailsEdges) UseHeldItemOrErr() (*HeldItems, error) {
	if e.loadedTypes[1] {
		if e.UseHeldItem == nil {
			// The edge use_held_item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: helditems.Label}
		}
		return e.UseHeldItem, nil
	}
	return nil, &NotLoadedError{edge: "use_held_item"}
}

// UseMove1OrErr returns the UseMove1 value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrainedPokemonDetailsEdges) UseMove1OrErr() (*Moves, error) {
	if e.loadedTypes[2] {
		if e.UseMove1 == nil {
			// The edge use_move1 was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: moves.Label}
		}
		return e.UseMove1, nil
	}
	return nil, &NotLoadedError{edge: "use_move1"}
}

// UseMove2OrErr returns the UseMove2 value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrainedPokemonDetailsEdges) UseMove2OrErr() (*Moves, error) {
	if e.loadedTypes[3] {
		if e.UseMove2 == nil {
			// The edge use_move2 was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: moves.Label}
		}
		return e.UseMove2, nil
	}
	return nil, &NotLoadedError{edge: "use_move2"}
}

// UseMove3OrErr returns the UseMove3 value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrainedPokemonDetailsEdges) UseMove3OrErr() (*Moves, error) {
	if e.loadedTypes[4] {
		if e.UseMove3 == nil {
			// The edge use_move3 was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: moves.Label}
		}
		return e.UseMove3, nil
	}
	return nil, &NotLoadedError{edge: "use_move3"}
}

// UseMove4OrErr returns the UseMove4 value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrainedPokemonDetailsEdges) UseMove4OrErr() (*Moves, error) {
	if e.loadedTypes[5] {
		if e.UseMove4 == nil {
			// The edge use_move4 was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: moves.Label}
		}
		return e.UseMove4, nil
	}
	return nil, &NotLoadedError{edge: "use_move4"}
}

// TrainingDetailUserOrErr returns the TrainingDetailUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrainedPokemonDetailsEdges) TrainingDetailUserOrErr() (*Users, error) {
	if e.loadedTypes[6] {
		if e.TrainingDetailUser == nil {
			// The edge training_detail_user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: users.Label}
		}
		return e.TrainingDetailUser, nil
	}
	return nil, &NotLoadedError{edge: "training_detail_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TrainedPokemonDetails) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case trainedpokemondetails.FieldID, trainedpokemondetails.FieldAbilityID, trainedpokemondetails.FieldHeldItemID, trainedpokemondetails.FieldMoveId1, trainedpokemondetails.FieldMoveId2, trainedpokemondetails.FieldMoveId3, trainedpokemondetails.FieldMoveId4, trainedpokemondetails.FieldUserID:
			values[i] = new(sql.NullInt64)
		case trainedpokemondetails.FieldNickname, trainedpokemondetails.FieldGender:
			values[i] = new(sql.NullString)
		case trainedpokemondetails.FieldCreateTime, trainedpokemondetails.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TrainedPokemonDetails", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TrainedPokemonDetails fields.
func (tpd *TrainedPokemonDetails) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case trainedpokemondetails.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tpd.ID = int(value.Int64)
		case trainedpokemondetails.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				tpd.CreateTime = value.Time
			}
		case trainedpokemondetails.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				tpd.UpdateTime = value.Time
			}
		case trainedpokemondetails.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				tpd.Nickname = value.String
			}
		case trainedpokemondetails.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				tpd.Gender = property.Gender(value.String)
			}
		case trainedpokemondetails.FieldAbilityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ability_id", values[i])
			} else if value.Valid {
				tpd.AbilityID = int(value.Int64)
			}
		case trainedpokemondetails.FieldHeldItemID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field held_item_id", values[i])
			} else if value.Valid {
				tpd.HeldItemID = int(value.Int64)
			}
		case trainedpokemondetails.FieldMoveId1:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field move_id1", values[i])
			} else if value.Valid {
				tpd.MoveId1 = int(value.Int64)
			}
		case trainedpokemondetails.FieldMoveId2:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field move_id2", values[i])
			} else if value.Valid {
				tpd.MoveId2 = int(value.Int64)
			}
		case trainedpokemondetails.FieldMoveId3:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field move_id3", values[i])
			} else if value.Valid {
				tpd.MoveId3 = int(value.Int64)
			}
		case trainedpokemondetails.FieldMoveId4:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field move_id4", values[i])
			} else if value.Valid {
				tpd.MoveId4 = int(value.Int64)
			}
		case trainedpokemondetails.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				tpd.UserID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUseAbility queries the "use_ability" edge of the TrainedPokemonDetails entity.
func (tpd *TrainedPokemonDetails) QueryUseAbility() *AbilitiesQuery {
	return (&TrainedPokemonDetailsClient{config: tpd.config}).QueryUseAbility(tpd)
}

// QueryUseHeldItem queries the "use_held_item" edge of the TrainedPokemonDetails entity.
func (tpd *TrainedPokemonDetails) QueryUseHeldItem() *HeldItemsQuery {
	return (&TrainedPokemonDetailsClient{config: tpd.config}).QueryUseHeldItem(tpd)
}

// QueryUseMove1 queries the "use_move1" edge of the TrainedPokemonDetails entity.
func (tpd *TrainedPokemonDetails) QueryUseMove1() *MovesQuery {
	return (&TrainedPokemonDetailsClient{config: tpd.config}).QueryUseMove1(tpd)
}

// QueryUseMove2 queries the "use_move2" edge of the TrainedPokemonDetails entity.
func (tpd *TrainedPokemonDetails) QueryUseMove2() *MovesQuery {
	return (&TrainedPokemonDetailsClient{config: tpd.config}).QueryUseMove2(tpd)
}

// QueryUseMove3 queries the "use_move3" edge of the TrainedPokemonDetails entity.
func (tpd *TrainedPokemonDetails) QueryUseMove3() *MovesQuery {
	return (&TrainedPokemonDetailsClient{config: tpd.config}).QueryUseMove3(tpd)
}

// QueryUseMove4 queries the "use_move4" edge of the TrainedPokemonDetails entity.
func (tpd *TrainedPokemonDetails) QueryUseMove4() *MovesQuery {
	return (&TrainedPokemonDetailsClient{config: tpd.config}).QueryUseMove4(tpd)
}

// QueryTrainingDetailUser queries the "training_detail_user" edge of the TrainedPokemonDetails entity.
func (tpd *TrainedPokemonDetails) QueryTrainingDetailUser() *UsersQuery {
	return (&TrainedPokemonDetailsClient{config: tpd.config}).QueryTrainingDetailUser(tpd)
}

// Update returns a builder for updating this TrainedPokemonDetails.
// Note that you need to call TrainedPokemonDetails.Unwrap() before calling this method if this TrainedPokemonDetails
// was returned from a transaction, and the transaction was committed or rolled back.
func (tpd *TrainedPokemonDetails) Update() *TrainedPokemonDetailsUpdateOne {
	return (&TrainedPokemonDetailsClient{config: tpd.config}).UpdateOne(tpd)
}

// Unwrap unwraps the TrainedPokemonDetails entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tpd *TrainedPokemonDetails) Unwrap() *TrainedPokemonDetails {
	tx, ok := tpd.config.driver.(*txDriver)
	if !ok {
		panic("ent: TrainedPokemonDetails is not a transactional entity")
	}
	tpd.config.driver = tx.drv
	return tpd
}

// String implements the fmt.Stringer.
func (tpd *TrainedPokemonDetails) String() string {
	var builder strings.Builder
	builder.WriteString("TrainedPokemonDetails(")
	builder.WriteString(fmt.Sprintf("id=%v", tpd.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(tpd.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(tpd.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", nickname=")
	builder.WriteString(tpd.Nickname)
	builder.WriteString(", gender=")
	builder.WriteString(fmt.Sprintf("%v", tpd.Gender))
	builder.WriteString(", ability_id=")
	builder.WriteString(fmt.Sprintf("%v", tpd.AbilityID))
	builder.WriteString(", held_item_id=")
	builder.WriteString(fmt.Sprintf("%v", tpd.HeldItemID))
	builder.WriteString(", move_id1=")
	builder.WriteString(fmt.Sprintf("%v", tpd.MoveId1))
	builder.WriteString(", move_id2=")
	builder.WriteString(fmt.Sprintf("%v", tpd.MoveId2))
	builder.WriteString(", move_id3=")
	builder.WriteString(fmt.Sprintf("%v", tpd.MoveId3))
	builder.WriteString(", move_id4=")
	builder.WriteString(fmt.Sprintf("%v", tpd.MoveId4))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", tpd.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// TrainedPokemonDetailsSlice is a parsable slice of TrainedPokemonDetails.
type TrainedPokemonDetailsSlice []*TrainedPokemonDetails

func (tpd TrainedPokemonDetailsSlice) config(cfg config) {
	for _i := range tpd {
		tpd[_i].config = cfg
	}
}