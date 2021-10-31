// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/battleopponentparty"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/predicate"
)

// BattleOpponentPartyDelete is the builder for deleting a BattleOpponentParty entity.
type BattleOpponentPartyDelete struct {
	config
	hooks    []Hook
	mutation *BattleOpponentPartyMutation
}

// Where appends a list predicates to the BattleOpponentPartyDelete builder.
func (bopd *BattleOpponentPartyDelete) Where(ps ...predicate.BattleOpponentParty) *BattleOpponentPartyDelete {
	bopd.mutation.Where(ps...)
	return bopd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bopd *BattleOpponentPartyDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bopd.hooks) == 0 {
		affected, err = bopd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BattleOpponentPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bopd.mutation = mutation
			affected, err = bopd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bopd.hooks) - 1; i >= 0; i-- {
			if bopd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bopd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bopd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bopd *BattleOpponentPartyDelete) ExecX(ctx context.Context) int {
	n, err := bopd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bopd *BattleOpponentPartyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: battleopponentparty.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: battleopponentparty.FieldID,
			},
		},
	}
	if ps := bopd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bopd.driver, _spec)
}

// BattleOpponentPartyDeleteOne is the builder for deleting a single BattleOpponentParty entity.
type BattleOpponentPartyDeleteOne struct {
	bopd *BattleOpponentPartyDelete
}

// Exec executes the deletion query.
func (bopdo *BattleOpponentPartyDeleteOne) Exec(ctx context.Context) error {
	n, err := bopdo.bopd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{battleopponentparty.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bopdo *BattleOpponentPartyDeleteOne) ExecX(ctx context.Context) {
	bopdo.bopd.ExecX(ctx)
}
