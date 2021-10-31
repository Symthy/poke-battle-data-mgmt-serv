// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/battleopponentparty"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/battlerecords"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/predicate"
)

// BattleOpponentPartyUpdate is the builder for updating BattleOpponentParty entities.
type BattleOpponentPartyUpdate struct {
	config
	hooks    []Hook
	mutation *BattleOpponentPartyMutation
}

// Where appends a list predicates to the BattleOpponentPartyUpdate builder.
func (bopu *BattleOpponentPartyUpdate) Where(ps ...predicate.BattleOpponentParty) *BattleOpponentPartyUpdate {
	bopu.mutation.Where(ps...)
	return bopu
}

// SetOpponentPokemonId1 sets the "opponent_pokemon_id1" field.
func (bopu *BattleOpponentPartyUpdate) SetOpponentPokemonId1(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.ResetOpponentPokemonId1()
	bopu.mutation.SetOpponentPokemonId1(i)
	return bopu
}

// AddOpponentPokemonId1 adds i to the "opponent_pokemon_id1" field.
func (bopu *BattleOpponentPartyUpdate) AddOpponentPokemonId1(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.AddOpponentPokemonId1(i)
	return bopu
}

// SetOpponentPokemonId2 sets the "opponent_pokemon_id2" field.
func (bopu *BattleOpponentPartyUpdate) SetOpponentPokemonId2(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.ResetOpponentPokemonId2()
	bopu.mutation.SetOpponentPokemonId2(i)
	return bopu
}

// AddOpponentPokemonId2 adds i to the "opponent_pokemon_id2" field.
func (bopu *BattleOpponentPartyUpdate) AddOpponentPokemonId2(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.AddOpponentPokemonId2(i)
	return bopu
}

// SetOpponentPokemonId3 sets the "opponent_pokemon_id3" field.
func (bopu *BattleOpponentPartyUpdate) SetOpponentPokemonId3(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.ResetOpponentPokemonId3()
	bopu.mutation.SetOpponentPokemonId3(i)
	return bopu
}

// AddOpponentPokemonId3 adds i to the "opponent_pokemon_id3" field.
func (bopu *BattleOpponentPartyUpdate) AddOpponentPokemonId3(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.AddOpponentPokemonId3(i)
	return bopu
}

// SetOpponentPokemonId4 sets the "opponent_pokemon_id4" field.
func (bopu *BattleOpponentPartyUpdate) SetOpponentPokemonId4(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.ResetOpponentPokemonId4()
	bopu.mutation.SetOpponentPokemonId4(i)
	return bopu
}

// SetNillableOpponentPokemonId4 sets the "opponent_pokemon_id4" field if the given value is not nil.
func (bopu *BattleOpponentPartyUpdate) SetNillableOpponentPokemonId4(i *int) *BattleOpponentPartyUpdate {
	if i != nil {
		bopu.SetOpponentPokemonId4(*i)
	}
	return bopu
}

// AddOpponentPokemonId4 adds i to the "opponent_pokemon_id4" field.
func (bopu *BattleOpponentPartyUpdate) AddOpponentPokemonId4(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.AddOpponentPokemonId4(i)
	return bopu
}

// ClearOpponentPokemonId4 clears the value of the "opponent_pokemon_id4" field.
func (bopu *BattleOpponentPartyUpdate) ClearOpponentPokemonId4() *BattleOpponentPartyUpdate {
	bopu.mutation.ClearOpponentPokemonId4()
	return bopu
}

// SetOpponentPokemonId5 sets the "opponent_pokemon_id5" field.
func (bopu *BattleOpponentPartyUpdate) SetOpponentPokemonId5(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.ResetOpponentPokemonId5()
	bopu.mutation.SetOpponentPokemonId5(i)
	return bopu
}

// SetNillableOpponentPokemonId5 sets the "opponent_pokemon_id5" field if the given value is not nil.
func (bopu *BattleOpponentPartyUpdate) SetNillableOpponentPokemonId5(i *int) *BattleOpponentPartyUpdate {
	if i != nil {
		bopu.SetOpponentPokemonId5(*i)
	}
	return bopu
}

// AddOpponentPokemonId5 adds i to the "opponent_pokemon_id5" field.
func (bopu *BattleOpponentPartyUpdate) AddOpponentPokemonId5(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.AddOpponentPokemonId5(i)
	return bopu
}

// ClearOpponentPokemonId5 clears the value of the "opponent_pokemon_id5" field.
func (bopu *BattleOpponentPartyUpdate) ClearOpponentPokemonId5() *BattleOpponentPartyUpdate {
	bopu.mutation.ClearOpponentPokemonId5()
	return bopu
}

// SetOpponentPokemonId6 sets the "opponent_pokemon_id6" field.
func (bopu *BattleOpponentPartyUpdate) SetOpponentPokemonId6(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.ResetOpponentPokemonId6()
	bopu.mutation.SetOpponentPokemonId6(i)
	return bopu
}

// SetNillableOpponentPokemonId6 sets the "opponent_pokemon_id6" field if the given value is not nil.
func (bopu *BattleOpponentPartyUpdate) SetNillableOpponentPokemonId6(i *int) *BattleOpponentPartyUpdate {
	if i != nil {
		bopu.SetOpponentPokemonId6(*i)
	}
	return bopu
}

// AddOpponentPokemonId6 adds i to the "opponent_pokemon_id6" field.
func (bopu *BattleOpponentPartyUpdate) AddOpponentPokemonId6(i int) *BattleOpponentPartyUpdate {
	bopu.mutation.AddOpponentPokemonId6(i)
	return bopu
}

// ClearOpponentPokemonId6 clears the value of the "opponent_pokemon_id6" field.
func (bopu *BattleOpponentPartyUpdate) ClearOpponentPokemonId6() *BattleOpponentPartyUpdate {
	bopu.mutation.ClearOpponentPokemonId6()
	return bopu
}

// AddBattleContentIDs adds the "battle_content" edge to the BattleRecords entity by IDs.
func (bopu *BattleOpponentPartyUpdate) AddBattleContentIDs(ids ...int) *BattleOpponentPartyUpdate {
	bopu.mutation.AddBattleContentIDs(ids...)
	return bopu
}

// AddBattleContent adds the "battle_content" edges to the BattleRecords entity.
func (bopu *BattleOpponentPartyUpdate) AddBattleContent(b ...*BattleRecords) *BattleOpponentPartyUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bopu.AddBattleContentIDs(ids...)
}

// Mutation returns the BattleOpponentPartyMutation object of the builder.
func (bopu *BattleOpponentPartyUpdate) Mutation() *BattleOpponentPartyMutation {
	return bopu.mutation
}

// ClearBattleContent clears all "battle_content" edges to the BattleRecords entity.
func (bopu *BattleOpponentPartyUpdate) ClearBattleContent() *BattleOpponentPartyUpdate {
	bopu.mutation.ClearBattleContent()
	return bopu
}

// RemoveBattleContentIDs removes the "battle_content" edge to BattleRecords entities by IDs.
func (bopu *BattleOpponentPartyUpdate) RemoveBattleContentIDs(ids ...int) *BattleOpponentPartyUpdate {
	bopu.mutation.RemoveBattleContentIDs(ids...)
	return bopu
}

// RemoveBattleContent removes "battle_content" edges to BattleRecords entities.
func (bopu *BattleOpponentPartyUpdate) RemoveBattleContent(b ...*BattleRecords) *BattleOpponentPartyUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bopu.RemoveBattleContentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bopu *BattleOpponentPartyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bopu.hooks) == 0 {
		if err = bopu.check(); err != nil {
			return 0, err
		}
		affected, err = bopu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BattleOpponentPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bopu.check(); err != nil {
				return 0, err
			}
			bopu.mutation = mutation
			affected, err = bopu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bopu.hooks) - 1; i >= 0; i-- {
			if bopu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bopu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bopu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bopu *BattleOpponentPartyUpdate) SaveX(ctx context.Context) int {
	affected, err := bopu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bopu *BattleOpponentPartyUpdate) Exec(ctx context.Context) error {
	_, err := bopu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bopu *BattleOpponentPartyUpdate) ExecX(ctx context.Context) {
	if err := bopu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bopu *BattleOpponentPartyUpdate) check() error {
	if v, ok := bopu.mutation.OpponentPokemonId1(); ok {
		if err := battleopponentparty.OpponentPokemonId1Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id1", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id1\": %w", err)}
		}
	}
	if v, ok := bopu.mutation.OpponentPokemonId2(); ok {
		if err := battleopponentparty.OpponentPokemonId2Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id2", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id2\": %w", err)}
		}
	}
	if v, ok := bopu.mutation.OpponentPokemonId3(); ok {
		if err := battleopponentparty.OpponentPokemonId3Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id3", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id3\": %w", err)}
		}
	}
	if v, ok := bopu.mutation.OpponentPokemonId4(); ok {
		if err := battleopponentparty.OpponentPokemonId4Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id4", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id4\": %w", err)}
		}
	}
	if v, ok := bopu.mutation.OpponentPokemonId5(); ok {
		if err := battleopponentparty.OpponentPokemonId5Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id5", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id5\": %w", err)}
		}
	}
	if v, ok := bopu.mutation.OpponentPokemonId6(); ok {
		if err := battleopponentparty.OpponentPokemonId6Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id6", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id6\": %w", err)}
		}
	}
	return nil
}

func (bopu *BattleOpponentPartyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   battleopponentparty.Table,
			Columns: battleopponentparty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: battleopponentparty.FieldID,
			},
		},
	}
	if ps := bopu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bopu.mutation.OpponentPokemonId1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId1,
		})
	}
	if value, ok := bopu.mutation.AddedOpponentPokemonId1(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId1,
		})
	}
	if value, ok := bopu.mutation.OpponentPokemonId2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId2,
		})
	}
	if value, ok := bopu.mutation.AddedOpponentPokemonId2(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId2,
		})
	}
	if value, ok := bopu.mutation.OpponentPokemonId3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId3,
		})
	}
	if value, ok := bopu.mutation.AddedOpponentPokemonId3(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId3,
		})
	}
	if value, ok := bopu.mutation.OpponentPokemonId4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId4,
		})
	}
	if value, ok := bopu.mutation.AddedOpponentPokemonId4(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId4,
		})
	}
	if bopu.mutation.OpponentPokemonId4Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: battleopponentparty.FieldOpponentPokemonId4,
		})
	}
	if value, ok := bopu.mutation.OpponentPokemonId5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId5,
		})
	}
	if value, ok := bopu.mutation.AddedOpponentPokemonId5(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId5,
		})
	}
	if bopu.mutation.OpponentPokemonId5Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: battleopponentparty.FieldOpponentPokemonId5,
		})
	}
	if value, ok := bopu.mutation.OpponentPokemonId6(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId6,
		})
	}
	if value, ok := bopu.mutation.AddedOpponentPokemonId6(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId6,
		})
	}
	if bopu.mutation.OpponentPokemonId6Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: battleopponentparty.FieldOpponentPokemonId6,
		})
	}
	if bopu.mutation.BattleContentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   battleopponentparty.BattleContentTable,
			Columns: []string{battleopponentparty.BattleContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battlerecords.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bopu.mutation.RemovedBattleContentIDs(); len(nodes) > 0 && !bopu.mutation.BattleContentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   battleopponentparty.BattleContentTable,
			Columns: []string{battleopponentparty.BattleContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battlerecords.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bopu.mutation.BattleContentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   battleopponentparty.BattleContentTable,
			Columns: []string{battleopponentparty.BattleContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battlerecords.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bopu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{battleopponentparty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BattleOpponentPartyUpdateOne is the builder for updating a single BattleOpponentParty entity.
type BattleOpponentPartyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BattleOpponentPartyMutation
}

// SetOpponentPokemonId1 sets the "opponent_pokemon_id1" field.
func (bopuo *BattleOpponentPartyUpdateOne) SetOpponentPokemonId1(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.ResetOpponentPokemonId1()
	bopuo.mutation.SetOpponentPokemonId1(i)
	return bopuo
}

// AddOpponentPokemonId1 adds i to the "opponent_pokemon_id1" field.
func (bopuo *BattleOpponentPartyUpdateOne) AddOpponentPokemonId1(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.AddOpponentPokemonId1(i)
	return bopuo
}

// SetOpponentPokemonId2 sets the "opponent_pokemon_id2" field.
func (bopuo *BattleOpponentPartyUpdateOne) SetOpponentPokemonId2(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.ResetOpponentPokemonId2()
	bopuo.mutation.SetOpponentPokemonId2(i)
	return bopuo
}

// AddOpponentPokemonId2 adds i to the "opponent_pokemon_id2" field.
func (bopuo *BattleOpponentPartyUpdateOne) AddOpponentPokemonId2(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.AddOpponentPokemonId2(i)
	return bopuo
}

// SetOpponentPokemonId3 sets the "opponent_pokemon_id3" field.
func (bopuo *BattleOpponentPartyUpdateOne) SetOpponentPokemonId3(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.ResetOpponentPokemonId3()
	bopuo.mutation.SetOpponentPokemonId3(i)
	return bopuo
}

// AddOpponentPokemonId3 adds i to the "opponent_pokemon_id3" field.
func (bopuo *BattleOpponentPartyUpdateOne) AddOpponentPokemonId3(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.AddOpponentPokemonId3(i)
	return bopuo
}

// SetOpponentPokemonId4 sets the "opponent_pokemon_id4" field.
func (bopuo *BattleOpponentPartyUpdateOne) SetOpponentPokemonId4(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.ResetOpponentPokemonId4()
	bopuo.mutation.SetOpponentPokemonId4(i)
	return bopuo
}

// SetNillableOpponentPokemonId4 sets the "opponent_pokemon_id4" field if the given value is not nil.
func (bopuo *BattleOpponentPartyUpdateOne) SetNillableOpponentPokemonId4(i *int) *BattleOpponentPartyUpdateOne {
	if i != nil {
		bopuo.SetOpponentPokemonId4(*i)
	}
	return bopuo
}

// AddOpponentPokemonId4 adds i to the "opponent_pokemon_id4" field.
func (bopuo *BattleOpponentPartyUpdateOne) AddOpponentPokemonId4(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.AddOpponentPokemonId4(i)
	return bopuo
}

// ClearOpponentPokemonId4 clears the value of the "opponent_pokemon_id4" field.
func (bopuo *BattleOpponentPartyUpdateOne) ClearOpponentPokemonId4() *BattleOpponentPartyUpdateOne {
	bopuo.mutation.ClearOpponentPokemonId4()
	return bopuo
}

// SetOpponentPokemonId5 sets the "opponent_pokemon_id5" field.
func (bopuo *BattleOpponentPartyUpdateOne) SetOpponentPokemonId5(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.ResetOpponentPokemonId5()
	bopuo.mutation.SetOpponentPokemonId5(i)
	return bopuo
}

// SetNillableOpponentPokemonId5 sets the "opponent_pokemon_id5" field if the given value is not nil.
func (bopuo *BattleOpponentPartyUpdateOne) SetNillableOpponentPokemonId5(i *int) *BattleOpponentPartyUpdateOne {
	if i != nil {
		bopuo.SetOpponentPokemonId5(*i)
	}
	return bopuo
}

// AddOpponentPokemonId5 adds i to the "opponent_pokemon_id5" field.
func (bopuo *BattleOpponentPartyUpdateOne) AddOpponentPokemonId5(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.AddOpponentPokemonId5(i)
	return bopuo
}

// ClearOpponentPokemonId5 clears the value of the "opponent_pokemon_id5" field.
func (bopuo *BattleOpponentPartyUpdateOne) ClearOpponentPokemonId5() *BattleOpponentPartyUpdateOne {
	bopuo.mutation.ClearOpponentPokemonId5()
	return bopuo
}

// SetOpponentPokemonId6 sets the "opponent_pokemon_id6" field.
func (bopuo *BattleOpponentPartyUpdateOne) SetOpponentPokemonId6(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.ResetOpponentPokemonId6()
	bopuo.mutation.SetOpponentPokemonId6(i)
	return bopuo
}

// SetNillableOpponentPokemonId6 sets the "opponent_pokemon_id6" field if the given value is not nil.
func (bopuo *BattleOpponentPartyUpdateOne) SetNillableOpponentPokemonId6(i *int) *BattleOpponentPartyUpdateOne {
	if i != nil {
		bopuo.SetOpponentPokemonId6(*i)
	}
	return bopuo
}

// AddOpponentPokemonId6 adds i to the "opponent_pokemon_id6" field.
func (bopuo *BattleOpponentPartyUpdateOne) AddOpponentPokemonId6(i int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.AddOpponentPokemonId6(i)
	return bopuo
}

// ClearOpponentPokemonId6 clears the value of the "opponent_pokemon_id6" field.
func (bopuo *BattleOpponentPartyUpdateOne) ClearOpponentPokemonId6() *BattleOpponentPartyUpdateOne {
	bopuo.mutation.ClearOpponentPokemonId6()
	return bopuo
}

// AddBattleContentIDs adds the "battle_content" edge to the BattleRecords entity by IDs.
func (bopuo *BattleOpponentPartyUpdateOne) AddBattleContentIDs(ids ...int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.AddBattleContentIDs(ids...)
	return bopuo
}

// AddBattleContent adds the "battle_content" edges to the BattleRecords entity.
func (bopuo *BattleOpponentPartyUpdateOne) AddBattleContent(b ...*BattleRecords) *BattleOpponentPartyUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bopuo.AddBattleContentIDs(ids...)
}

// Mutation returns the BattleOpponentPartyMutation object of the builder.
func (bopuo *BattleOpponentPartyUpdateOne) Mutation() *BattleOpponentPartyMutation {
	return bopuo.mutation
}

// ClearBattleContent clears all "battle_content" edges to the BattleRecords entity.
func (bopuo *BattleOpponentPartyUpdateOne) ClearBattleContent() *BattleOpponentPartyUpdateOne {
	bopuo.mutation.ClearBattleContent()
	return bopuo
}

// RemoveBattleContentIDs removes the "battle_content" edge to BattleRecords entities by IDs.
func (bopuo *BattleOpponentPartyUpdateOne) RemoveBattleContentIDs(ids ...int) *BattleOpponentPartyUpdateOne {
	bopuo.mutation.RemoveBattleContentIDs(ids...)
	return bopuo
}

// RemoveBattleContent removes "battle_content" edges to BattleRecords entities.
func (bopuo *BattleOpponentPartyUpdateOne) RemoveBattleContent(b ...*BattleRecords) *BattleOpponentPartyUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bopuo.RemoveBattleContentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bopuo *BattleOpponentPartyUpdateOne) Select(field string, fields ...string) *BattleOpponentPartyUpdateOne {
	bopuo.fields = append([]string{field}, fields...)
	return bopuo
}

// Save executes the query and returns the updated BattleOpponentParty entity.
func (bopuo *BattleOpponentPartyUpdateOne) Save(ctx context.Context) (*BattleOpponentParty, error) {
	var (
		err  error
		node *BattleOpponentParty
	)
	if len(bopuo.hooks) == 0 {
		if err = bopuo.check(); err != nil {
			return nil, err
		}
		node, err = bopuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BattleOpponentPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bopuo.check(); err != nil {
				return nil, err
			}
			bopuo.mutation = mutation
			node, err = bopuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bopuo.hooks) - 1; i >= 0; i-- {
			if bopuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bopuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bopuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bopuo *BattleOpponentPartyUpdateOne) SaveX(ctx context.Context) *BattleOpponentParty {
	node, err := bopuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bopuo *BattleOpponentPartyUpdateOne) Exec(ctx context.Context) error {
	_, err := bopuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bopuo *BattleOpponentPartyUpdateOne) ExecX(ctx context.Context) {
	if err := bopuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bopuo *BattleOpponentPartyUpdateOne) check() error {
	if v, ok := bopuo.mutation.OpponentPokemonId1(); ok {
		if err := battleopponentparty.OpponentPokemonId1Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id1", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id1\": %w", err)}
		}
	}
	if v, ok := bopuo.mutation.OpponentPokemonId2(); ok {
		if err := battleopponentparty.OpponentPokemonId2Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id2", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id2\": %w", err)}
		}
	}
	if v, ok := bopuo.mutation.OpponentPokemonId3(); ok {
		if err := battleopponentparty.OpponentPokemonId3Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id3", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id3\": %w", err)}
		}
	}
	if v, ok := bopuo.mutation.OpponentPokemonId4(); ok {
		if err := battleopponentparty.OpponentPokemonId4Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id4", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id4\": %w", err)}
		}
	}
	if v, ok := bopuo.mutation.OpponentPokemonId5(); ok {
		if err := battleopponentparty.OpponentPokemonId5Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id5", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id5\": %w", err)}
		}
	}
	if v, ok := bopuo.mutation.OpponentPokemonId6(); ok {
		if err := battleopponentparty.OpponentPokemonId6Validator(v); err != nil {
			return &ValidationError{Name: "opponent_pokemon_id6", err: fmt.Errorf("ent: validator failed for field \"opponent_pokemon_id6\": %w", err)}
		}
	}
	return nil
}

func (bopuo *BattleOpponentPartyUpdateOne) sqlSave(ctx context.Context) (_node *BattleOpponentParty, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   battleopponentparty.Table,
			Columns: battleopponentparty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: battleopponentparty.FieldID,
			},
		},
	}
	id, ok := bopuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing BattleOpponentParty.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := bopuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, battleopponentparty.FieldID)
		for _, f := range fields {
			if !battleopponentparty.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != battleopponentparty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bopuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bopuo.mutation.OpponentPokemonId1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId1,
		})
	}
	if value, ok := bopuo.mutation.AddedOpponentPokemonId1(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId1,
		})
	}
	if value, ok := bopuo.mutation.OpponentPokemonId2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId2,
		})
	}
	if value, ok := bopuo.mutation.AddedOpponentPokemonId2(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId2,
		})
	}
	if value, ok := bopuo.mutation.OpponentPokemonId3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId3,
		})
	}
	if value, ok := bopuo.mutation.AddedOpponentPokemonId3(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId3,
		})
	}
	if value, ok := bopuo.mutation.OpponentPokemonId4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId4,
		})
	}
	if value, ok := bopuo.mutation.AddedOpponentPokemonId4(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId4,
		})
	}
	if bopuo.mutation.OpponentPokemonId4Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: battleopponentparty.FieldOpponentPokemonId4,
		})
	}
	if value, ok := bopuo.mutation.OpponentPokemonId5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId5,
		})
	}
	if value, ok := bopuo.mutation.AddedOpponentPokemonId5(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId5,
		})
	}
	if bopuo.mutation.OpponentPokemonId5Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: battleopponentparty.FieldOpponentPokemonId5,
		})
	}
	if value, ok := bopuo.mutation.OpponentPokemonId6(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId6,
		})
	}
	if value, ok := bopuo.mutation.AddedOpponentPokemonId6(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battleopponentparty.FieldOpponentPokemonId6,
		})
	}
	if bopuo.mutation.OpponentPokemonId6Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: battleopponentparty.FieldOpponentPokemonId6,
		})
	}
	if bopuo.mutation.BattleContentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   battleopponentparty.BattleContentTable,
			Columns: []string{battleopponentparty.BattleContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battlerecords.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bopuo.mutation.RemovedBattleContentIDs(); len(nodes) > 0 && !bopuo.mutation.BattleContentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   battleopponentparty.BattleContentTable,
			Columns: []string{battleopponentparty.BattleContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battlerecords.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bopuo.mutation.BattleContentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   battleopponentparty.BattleContentTable,
			Columns: []string{battleopponentparty.BattleContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battlerecords.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BattleOpponentParty{config: bopuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bopuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{battleopponentparty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
