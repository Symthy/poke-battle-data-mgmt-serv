// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/partyresultrecord"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/predicate"
)

// PartyResultRecordDelete is the builder for deleting a PartyResultRecord entity.
type PartyResultRecordDelete struct {
	config
	hooks    []Hook
	mutation *PartyResultRecordMutation
}

// Where appends a list predicates to the PartyResultRecordDelete builder.
func (prrd *PartyResultRecordDelete) Where(ps ...predicate.PartyResultRecord) *PartyResultRecordDelete {
	prrd.mutation.Where(ps...)
	return prrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (prrd *PartyResultRecordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(prrd.hooks) == 0 {
		affected, err = prrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyResultRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			prrd.mutation = mutation
			affected, err = prrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(prrd.hooks) - 1; i >= 0; i-- {
			if prrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = prrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, prrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (prrd *PartyResultRecordDelete) ExecX(ctx context.Context) int {
	n, err := prrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (prrd *PartyResultRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: partyresultrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partyresultrecord.FieldID,
			},
		},
	}
	if ps := prrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, prrd.driver, _spec)
}

// PartyResultRecordDeleteOne is the builder for deleting a single PartyResultRecord entity.
type PartyResultRecordDeleteOne struct {
	prrd *PartyResultRecordDelete
}

// Exec executes the deletion query.
func (prrdo *PartyResultRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := prrdo.prrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{partyresultrecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (prrdo *PartyResultRecordDeleteOne) ExecX(ctx context.Context) {
	prrdo.prrd.ExecX(ctx)
}
