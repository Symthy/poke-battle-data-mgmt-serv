// Code generated by entc, DO NOT EDIT.

package trainedpokemons

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/predicate"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/property"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// PokemonID applies equality check predicate on the "pokemon_id" field. It's identical to PokemonIDEQ.
func PokemonID(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPokemonID), v))
	})
}

// CreateUserID applies equality check predicate on the "create_user_id" field. It's identical to CreateUserIDEQ.
func CreateUserID(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateUserID), v))
	})
}

// Nature applies equality check predicate on the "nature" field. It's identical to NatureEQ.
func Nature(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNature), vc))
	})
}

// EffortValueH applies equality check predicate on the "effort_value_h" field. It's identical to EffortValueHEQ.
func EffortValueH(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueH), v))
	})
}

// EffortValueA applies equality check predicate on the "effort_value_a" field. It's identical to EffortValueAEQ.
func EffortValueA(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueA), v))
	})
}

// EffortValueB applies equality check predicate on the "effort_value_b" field. It's identical to EffortValueBEQ.
func EffortValueB(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueB), v))
	})
}

// EffortValueC applies equality check predicate on the "effort_value_c" field. It's identical to EffortValueCEQ.
func EffortValueC(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueC), v))
	})
}

// EffortValueD applies equality check predicate on the "effort_value_d" field. It's identical to EffortValueDEQ.
func EffortValueD(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueD), v))
	})
}

// EffortValueS applies equality check predicate on the "effort_value_s" field. It's identical to EffortValueSEQ.
func EffortValueS(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueS), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// PokemonIDEQ applies the EQ predicate on the "pokemon_id" field.
func PokemonIDEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPokemonID), v))
	})
}

// PokemonIDNEQ applies the NEQ predicate on the "pokemon_id" field.
func PokemonIDNEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPokemonID), v))
	})
}

// PokemonIDIn applies the In predicate on the "pokemon_id" field.
func PokemonIDIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPokemonID), v...))
	})
}

// PokemonIDNotIn applies the NotIn predicate on the "pokemon_id" field.
func PokemonIDNotIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPokemonID), v...))
	})
}

// CreateUserIDEQ applies the EQ predicate on the "create_user_id" field.
func CreateUserIDEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateUserID), v))
	})
}

// CreateUserIDNEQ applies the NEQ predicate on the "create_user_id" field.
func CreateUserIDNEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateUserID), v))
	})
}

// CreateUserIDIn applies the In predicate on the "create_user_id" field.
func CreateUserIDIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateUserID), v...))
	})
}

// CreateUserIDNotIn applies the NotIn predicate on the "create_user_id" field.
func CreateUserIDNotIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateUserID), v...))
	})
}

// CreateUserIDIsNil applies the IsNil predicate on the "create_user_id" field.
func CreateUserIDIsNil() predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateUserID)))
	})
}

// CreateUserIDNotNil applies the NotNil predicate on the "create_user_id" field.
func CreateUserIDNotNil() predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateUserID)))
	})
}

// NatureEQ applies the EQ predicate on the "nature" field.
func NatureEQ(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNature), vc))
	})
}

// NatureNEQ applies the NEQ predicate on the "nature" field.
func NatureNEQ(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNature), vc))
	})
}

// NatureIn applies the In predicate on the "nature" field.
func NatureIn(vs ...property.Nature) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNature), v...))
	})
}

// NatureNotIn applies the NotIn predicate on the "nature" field.
func NatureNotIn(vs ...property.Nature) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNature), v...))
	})
}

// NatureGT applies the GT predicate on the "nature" field.
func NatureGT(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNature), vc))
	})
}

// NatureGTE applies the GTE predicate on the "nature" field.
func NatureGTE(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNature), vc))
	})
}

// NatureLT applies the LT predicate on the "nature" field.
func NatureLT(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNature), vc))
	})
}

// NatureLTE applies the LTE predicate on the "nature" field.
func NatureLTE(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNature), vc))
	})
}

// NatureContains applies the Contains predicate on the "nature" field.
func NatureContains(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNature), vc))
	})
}

// NatureHasPrefix applies the HasPrefix predicate on the "nature" field.
func NatureHasPrefix(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNature), vc))
	})
}

// NatureHasSuffix applies the HasSuffix predicate on the "nature" field.
func NatureHasSuffix(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNature), vc))
	})
}

// NatureEqualFold applies the EqualFold predicate on the "nature" field.
func NatureEqualFold(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNature), vc))
	})
}

// NatureContainsFold applies the ContainsFold predicate on the "nature" field.
func NatureContainsFold(v property.Nature) predicate.TrainedPokemons {
	vc := string(v)
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNature), vc))
	})
}

// EffortValueHEQ applies the EQ predicate on the "effort_value_h" field.
func EffortValueHEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueH), v))
	})
}

// EffortValueHNEQ applies the NEQ predicate on the "effort_value_h" field.
func EffortValueHNEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEffortValueH), v))
	})
}

// EffortValueHIn applies the In predicate on the "effort_value_h" field.
func EffortValueHIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEffortValueH), v...))
	})
}

// EffortValueHNotIn applies the NotIn predicate on the "effort_value_h" field.
func EffortValueHNotIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEffortValueH), v...))
	})
}

// EffortValueHGT applies the GT predicate on the "effort_value_h" field.
func EffortValueHGT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEffortValueH), v))
	})
}

// EffortValueHGTE applies the GTE predicate on the "effort_value_h" field.
func EffortValueHGTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEffortValueH), v))
	})
}

// EffortValueHLT applies the LT predicate on the "effort_value_h" field.
func EffortValueHLT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEffortValueH), v))
	})
}

// EffortValueHLTE applies the LTE predicate on the "effort_value_h" field.
func EffortValueHLTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEffortValueH), v))
	})
}

// EffortValueAEQ applies the EQ predicate on the "effort_value_a" field.
func EffortValueAEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueA), v))
	})
}

// EffortValueANEQ applies the NEQ predicate on the "effort_value_a" field.
func EffortValueANEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEffortValueA), v))
	})
}

// EffortValueAIn applies the In predicate on the "effort_value_a" field.
func EffortValueAIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEffortValueA), v...))
	})
}

// EffortValueANotIn applies the NotIn predicate on the "effort_value_a" field.
func EffortValueANotIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEffortValueA), v...))
	})
}

// EffortValueAGT applies the GT predicate on the "effort_value_a" field.
func EffortValueAGT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEffortValueA), v))
	})
}

// EffortValueAGTE applies the GTE predicate on the "effort_value_a" field.
func EffortValueAGTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEffortValueA), v))
	})
}

// EffortValueALT applies the LT predicate on the "effort_value_a" field.
func EffortValueALT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEffortValueA), v))
	})
}

// EffortValueALTE applies the LTE predicate on the "effort_value_a" field.
func EffortValueALTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEffortValueA), v))
	})
}

// EffortValueBEQ applies the EQ predicate on the "effort_value_b" field.
func EffortValueBEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueB), v))
	})
}

// EffortValueBNEQ applies the NEQ predicate on the "effort_value_b" field.
func EffortValueBNEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEffortValueB), v))
	})
}

// EffortValueBIn applies the In predicate on the "effort_value_b" field.
func EffortValueBIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEffortValueB), v...))
	})
}

// EffortValueBNotIn applies the NotIn predicate on the "effort_value_b" field.
func EffortValueBNotIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEffortValueB), v...))
	})
}

// EffortValueBGT applies the GT predicate on the "effort_value_b" field.
func EffortValueBGT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEffortValueB), v))
	})
}

// EffortValueBGTE applies the GTE predicate on the "effort_value_b" field.
func EffortValueBGTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEffortValueB), v))
	})
}

// EffortValueBLT applies the LT predicate on the "effort_value_b" field.
func EffortValueBLT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEffortValueB), v))
	})
}

// EffortValueBLTE applies the LTE predicate on the "effort_value_b" field.
func EffortValueBLTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEffortValueB), v))
	})
}

// EffortValueCEQ applies the EQ predicate on the "effort_value_c" field.
func EffortValueCEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueC), v))
	})
}

// EffortValueCNEQ applies the NEQ predicate on the "effort_value_c" field.
func EffortValueCNEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEffortValueC), v))
	})
}

// EffortValueCIn applies the In predicate on the "effort_value_c" field.
func EffortValueCIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEffortValueC), v...))
	})
}

// EffortValueCNotIn applies the NotIn predicate on the "effort_value_c" field.
func EffortValueCNotIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEffortValueC), v...))
	})
}

// EffortValueCGT applies the GT predicate on the "effort_value_c" field.
func EffortValueCGT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEffortValueC), v))
	})
}

// EffortValueCGTE applies the GTE predicate on the "effort_value_c" field.
func EffortValueCGTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEffortValueC), v))
	})
}

// EffortValueCLT applies the LT predicate on the "effort_value_c" field.
func EffortValueCLT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEffortValueC), v))
	})
}

// EffortValueCLTE applies the LTE predicate on the "effort_value_c" field.
func EffortValueCLTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEffortValueC), v))
	})
}

// EffortValueDEQ applies the EQ predicate on the "effort_value_d" field.
func EffortValueDEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueD), v))
	})
}

// EffortValueDNEQ applies the NEQ predicate on the "effort_value_d" field.
func EffortValueDNEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEffortValueD), v))
	})
}

// EffortValueDIn applies the In predicate on the "effort_value_d" field.
func EffortValueDIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEffortValueD), v...))
	})
}

// EffortValueDNotIn applies the NotIn predicate on the "effort_value_d" field.
func EffortValueDNotIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEffortValueD), v...))
	})
}

// EffortValueDGT applies the GT predicate on the "effort_value_d" field.
func EffortValueDGT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEffortValueD), v))
	})
}

// EffortValueDGTE applies the GTE predicate on the "effort_value_d" field.
func EffortValueDGTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEffortValueD), v))
	})
}

// EffortValueDLT applies the LT predicate on the "effort_value_d" field.
func EffortValueDLT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEffortValueD), v))
	})
}

// EffortValueDLTE applies the LTE predicate on the "effort_value_d" field.
func EffortValueDLTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEffortValueD), v))
	})
}

// EffortValueSEQ applies the EQ predicate on the "effort_value_s" field.
func EffortValueSEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffortValueS), v))
	})
}

// EffortValueSNEQ applies the NEQ predicate on the "effort_value_s" field.
func EffortValueSNEQ(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEffortValueS), v))
	})
}

// EffortValueSIn applies the In predicate on the "effort_value_s" field.
func EffortValueSIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEffortValueS), v...))
	})
}

// EffortValueSNotIn applies the NotIn predicate on the "effort_value_s" field.
func EffortValueSNotIn(vs ...int) predicate.TrainedPokemons {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEffortValueS), v...))
	})
}

// EffortValueSGT applies the GT predicate on the "effort_value_s" field.
func EffortValueSGT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEffortValueS), v))
	})
}

// EffortValueSGTE applies the GTE predicate on the "effort_value_s" field.
func EffortValueSGTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEffortValueS), v))
	})
}

// EffortValueSLT applies the LT predicate on the "effort_value_s" field.
func EffortValueSLT(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEffortValueS), v))
	})
}

// EffortValueSLTE applies the LTE predicate on the "effort_value_s" field.
func EffortValueSLTE(v int) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEffortValueS), v))
	})
}

// HasUsePokemon applies the HasEdge predicate on the "use_pokemon" edge.
func HasUsePokemon() predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsePokemonTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UsePokemonTable, UsePokemonColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsePokemonWith applies the HasEdge predicate on the "use_pokemon" edge with a given conditions (other predicates).
func HasUsePokemonWith(preds ...predicate.Pokemons) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsePokemonInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UsePokemonTable, UsePokemonColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTrainingUser applies the HasEdge predicate on the "training_user" edge.
func HasTrainingUser() predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TrainingUserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TrainingUserTable, TrainingUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTrainingUserWith applies the HasEdge predicate on the "training_user" edge with a given conditions (other predicates).
func HasTrainingUserWith(preds ...predicate.Users) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TrainingUserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TrainingUserTable, TrainingUserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TrainedPokemons) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TrainedPokemons) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TrainedPokemons) predicate.TrainedPokemons {
	return predicate.TrainedPokemons(func(s *sql.Selector) {
		p(s.Not())
	})
}