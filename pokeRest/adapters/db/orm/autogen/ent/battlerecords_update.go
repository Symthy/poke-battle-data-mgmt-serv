// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/battleopponentparty"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/battlerecords"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/party"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/predicate"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/property"
)

// BattleRecordsUpdate is the builder for updating BattleRecords entities.
type BattleRecordsUpdate struct {
	config
	hooks    []Hook
	mutation *BattleRecordsMutation
}

// Where appends a list predicates to the BattleRecordsUpdate builder.
func (bru *BattleRecordsUpdate) Where(ps ...predicate.BattleRecords) *BattleRecordsUpdate {
	bru.mutation.Where(ps...)
	return bru
}

// SetPartyID sets the "party_id" field.
func (bru *BattleRecordsUpdate) SetPartyID(i int) *BattleRecordsUpdate {
	bru.mutation.SetPartyID(i)
	return bru
}

// SetBattleFormat sets the "battle_format" field.
func (bru *BattleRecordsUpdate) SetBattleFormat(pf property.BattleFormats) *BattleRecordsUpdate {
	bru.mutation.SetBattleFormat(pf)
	return bru
}

// SetBattleOpponentPartyID sets the "battle_opponent_party_id" field.
func (bru *BattleRecordsUpdate) SetBattleOpponentPartyID(i int) *BattleRecordsUpdate {
	bru.mutation.SetBattleOpponentPartyID(i)
	return bru
}

// SetSelfElectionPokemonId1 sets the "self_election_pokemon_id1" field.
func (bru *BattleRecordsUpdate) SetSelfElectionPokemonId1(i int) *BattleRecordsUpdate {
	bru.mutation.ResetSelfElectionPokemonId1()
	bru.mutation.SetSelfElectionPokemonId1(i)
	return bru
}

// AddSelfElectionPokemonId1 adds i to the "self_election_pokemon_id1" field.
func (bru *BattleRecordsUpdate) AddSelfElectionPokemonId1(i int) *BattleRecordsUpdate {
	bru.mutation.AddSelfElectionPokemonId1(i)
	return bru
}

// SetSelfElectionPokemonId2 sets the "self_election_pokemon_id2" field.
func (bru *BattleRecordsUpdate) SetSelfElectionPokemonId2(i int) *BattleRecordsUpdate {
	bru.mutation.ResetSelfElectionPokemonId2()
	bru.mutation.SetSelfElectionPokemonId2(i)
	return bru
}

// AddSelfElectionPokemonId2 adds i to the "self_election_pokemon_id2" field.
func (bru *BattleRecordsUpdate) AddSelfElectionPokemonId2(i int) *BattleRecordsUpdate {
	bru.mutation.AddSelfElectionPokemonId2(i)
	return bru
}

// SetSelfElectionPokemonId3 sets the "self_election_pokemon_id3" field.
func (bru *BattleRecordsUpdate) SetSelfElectionPokemonId3(i int) *BattleRecordsUpdate {
	bru.mutation.ResetSelfElectionPokemonId3()
	bru.mutation.SetSelfElectionPokemonId3(i)
	return bru
}

// AddSelfElectionPokemonId3 adds i to the "self_election_pokemon_id3" field.
func (bru *BattleRecordsUpdate) AddSelfElectionPokemonId3(i int) *BattleRecordsUpdate {
	bru.mutation.AddSelfElectionPokemonId3(i)
	return bru
}

// SetSelfElectionPokemonId4 sets the "self_election_pokemon_id4" field.
func (bru *BattleRecordsUpdate) SetSelfElectionPokemonId4(i int) *BattleRecordsUpdate {
	bru.mutation.ResetSelfElectionPokemonId4()
	bru.mutation.SetSelfElectionPokemonId4(i)
	return bru
}

// SetNillableSelfElectionPokemonId4 sets the "self_election_pokemon_id4" field if the given value is not nil.
func (bru *BattleRecordsUpdate) SetNillableSelfElectionPokemonId4(i *int) *BattleRecordsUpdate {
	if i != nil {
		bru.SetSelfElectionPokemonId4(*i)
	}
	return bru
}

// AddSelfElectionPokemonId4 adds i to the "self_election_pokemon_id4" field.
func (bru *BattleRecordsUpdate) AddSelfElectionPokemonId4(i int) *BattleRecordsUpdate {
	bru.mutation.AddSelfElectionPokemonId4(i)
	return bru
}

// ClearSelfElectionPokemonId4 clears the value of the "self_election_pokemon_id4" field.
func (bru *BattleRecordsUpdate) ClearSelfElectionPokemonId4() *BattleRecordsUpdate {
	bru.mutation.ClearSelfElectionPokemonId4()
	return bru
}

// SetOpponentElectionPokemonId1 sets the "opponent_election_pokemon_id1" field.
func (bru *BattleRecordsUpdate) SetOpponentElectionPokemonId1(i int) *BattleRecordsUpdate {
	bru.mutation.ResetOpponentElectionPokemonId1()
	bru.mutation.SetOpponentElectionPokemonId1(i)
	return bru
}

// AddOpponentElectionPokemonId1 adds i to the "opponent_election_pokemon_id1" field.
func (bru *BattleRecordsUpdate) AddOpponentElectionPokemonId1(i int) *BattleRecordsUpdate {
	bru.mutation.AddOpponentElectionPokemonId1(i)
	return bru
}

// SetOpponentElectionPokemonId2 sets the "opponent_election_pokemon_id2" field.
func (bru *BattleRecordsUpdate) SetOpponentElectionPokemonId2(i int) *BattleRecordsUpdate {
	bru.mutation.ResetOpponentElectionPokemonId2()
	bru.mutation.SetOpponentElectionPokemonId2(i)
	return bru
}

// AddOpponentElectionPokemonId2 adds i to the "opponent_election_pokemon_id2" field.
func (bru *BattleRecordsUpdate) AddOpponentElectionPokemonId2(i int) *BattleRecordsUpdate {
	bru.mutation.AddOpponentElectionPokemonId2(i)
	return bru
}

// SetOpponentElectionPokemonId3 sets the "opponent_election_pokemon_id3" field.
func (bru *BattleRecordsUpdate) SetOpponentElectionPokemonId3(i int) *BattleRecordsUpdate {
	bru.mutation.ResetOpponentElectionPokemonId3()
	bru.mutation.SetOpponentElectionPokemonId3(i)
	return bru
}

// AddOpponentElectionPokemonId3 adds i to the "opponent_election_pokemon_id3" field.
func (bru *BattleRecordsUpdate) AddOpponentElectionPokemonId3(i int) *BattleRecordsUpdate {
	bru.mutation.AddOpponentElectionPokemonId3(i)
	return bru
}

// SetOpponentElectionPokemonId4 sets the "opponent_election_pokemon_id4" field.
func (bru *BattleRecordsUpdate) SetOpponentElectionPokemonId4(i int) *BattleRecordsUpdate {
	bru.mutation.ResetOpponentElectionPokemonId4()
	bru.mutation.SetOpponentElectionPokemonId4(i)
	return bru
}

// AddOpponentElectionPokemonId4 adds i to the "opponent_election_pokemon_id4" field.
func (bru *BattleRecordsUpdate) AddOpponentElectionPokemonId4(i int) *BattleRecordsUpdate {
	bru.mutation.AddOpponentElectionPokemonId4(i)
	return bru
}

// SetUsePartyID sets the "use_party" edge to the Party entity by ID.
func (bru *BattleRecordsUpdate) SetUsePartyID(id int) *BattleRecordsUpdate {
	bru.mutation.SetUsePartyID(id)
	return bru
}

// SetUseParty sets the "use_party" edge to the Party entity.
func (bru *BattleRecordsUpdate) SetUseParty(p *Party) *BattleRecordsUpdate {
	return bru.SetUsePartyID(p.ID)
}

// SetOpponentPartyID sets the "opponent_party" edge to the BattleOpponentParty entity by ID.
func (bru *BattleRecordsUpdate) SetOpponentPartyID(id int) *BattleRecordsUpdate {
	bru.mutation.SetOpponentPartyID(id)
	return bru
}

// SetOpponentParty sets the "opponent_party" edge to the BattleOpponentParty entity.
func (bru *BattleRecordsUpdate) SetOpponentParty(b *BattleOpponentParty) *BattleRecordsUpdate {
	return bru.SetOpponentPartyID(b.ID)
}

// Mutation returns the BattleRecordsMutation object of the builder.
func (bru *BattleRecordsUpdate) Mutation() *BattleRecordsMutation {
	return bru.mutation
}

// ClearUseParty clears the "use_party" edge to the Party entity.
func (bru *BattleRecordsUpdate) ClearUseParty() *BattleRecordsUpdate {
	bru.mutation.ClearUseParty()
	return bru
}

// ClearOpponentParty clears the "opponent_party" edge to the BattleOpponentParty entity.
func (bru *BattleRecordsUpdate) ClearOpponentParty() *BattleRecordsUpdate {
	bru.mutation.ClearOpponentParty()
	return bru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bru *BattleRecordsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bru.hooks) == 0 {
		if err = bru.check(); err != nil {
			return 0, err
		}
		affected, err = bru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BattleRecordsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bru.check(); err != nil {
				return 0, err
			}
			bru.mutation = mutation
			affected, err = bru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bru.hooks) - 1; i >= 0; i-- {
			if bru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bru *BattleRecordsUpdate) SaveX(ctx context.Context) int {
	affected, err := bru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bru *BattleRecordsUpdate) Exec(ctx context.Context) error {
	_, err := bru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bru *BattleRecordsUpdate) ExecX(ctx context.Context) {
	if err := bru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bru *BattleRecordsUpdate) check() error {
	if v, ok := bru.mutation.PartyID(); ok {
		if err := battlerecords.PartyIDValidator(v); err != nil {
			return &ValidationError{Name: "party_id", err: fmt.Errorf("ent: validator failed for field \"party_id\": %w", err)}
		}
	}
	if v, ok := bru.mutation.BattleOpponentPartyID(); ok {
		if err := battlerecords.BattleOpponentPartyIDValidator(v); err != nil {
			return &ValidationError{Name: "battle_opponent_party_id", err: fmt.Errorf("ent: validator failed for field \"battle_opponent_party_id\": %w", err)}
		}
	}
	if v, ok := bru.mutation.SelfElectionPokemonId1(); ok {
		if err := battlerecords.SelfElectionPokemonId1Validator(v); err != nil {
			return &ValidationError{Name: "self_election_pokemon_id1", err: fmt.Errorf("ent: validator failed for field \"self_election_pokemon_id1\": %w", err)}
		}
	}
	if v, ok := bru.mutation.SelfElectionPokemonId2(); ok {
		if err := battlerecords.SelfElectionPokemonId2Validator(v); err != nil {
			return &ValidationError{Name: "self_election_pokemon_id2", err: fmt.Errorf("ent: validator failed for field \"self_election_pokemon_id2\": %w", err)}
		}
	}
	if v, ok := bru.mutation.SelfElectionPokemonId3(); ok {
		if err := battlerecords.SelfElectionPokemonId3Validator(v); err != nil {
			return &ValidationError{Name: "self_election_pokemon_id3", err: fmt.Errorf("ent: validator failed for field \"self_election_pokemon_id3\": %w", err)}
		}
	}
	if v, ok := bru.mutation.SelfElectionPokemonId4(); ok {
		if err := battlerecords.SelfElectionPokemonId4Validator(v); err != nil {
			return &ValidationError{Name: "self_election_pokemon_id4", err: fmt.Errorf("ent: validator failed for field \"self_election_pokemon_id4\": %w", err)}
		}
	}
	if v, ok := bru.mutation.OpponentElectionPokemonId1(); ok {
		if err := battlerecords.OpponentElectionPokemonId1Validator(v); err != nil {
			return &ValidationError{Name: "opponent_election_pokemon_id1", err: fmt.Errorf("ent: validator failed for field \"opponent_election_pokemon_id1\": %w", err)}
		}
	}
	if v, ok := bru.mutation.OpponentElectionPokemonId2(); ok {
		if err := battlerecords.OpponentElectionPokemonId2Validator(v); err != nil {
			return &ValidationError{Name: "opponent_election_pokemon_id2", err: fmt.Errorf("ent: validator failed for field \"opponent_election_pokemon_id2\": %w", err)}
		}
	}
	if v, ok := bru.mutation.OpponentElectionPokemonId3(); ok {
		if err := battlerecords.OpponentElectionPokemonId3Validator(v); err != nil {
			return &ValidationError{Name: "opponent_election_pokemon_id3", err: fmt.Errorf("ent: validator failed for field \"opponent_election_pokemon_id3\": %w", err)}
		}
	}
	if v, ok := bru.mutation.OpponentElectionPokemonId4(); ok {
		if err := battlerecords.OpponentElectionPokemonId4Validator(v); err != nil {
			return &ValidationError{Name: "opponent_election_pokemon_id4", err: fmt.Errorf("ent: validator failed for field \"opponent_election_pokemon_id4\": %w", err)}
		}
	}
	if _, ok := bru.mutation.UsePartyID(); bru.mutation.UsePartyCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"use_party\"")
	}
	if _, ok := bru.mutation.OpponentPartyID(); bru.mutation.OpponentPartyCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"opponent_party\"")
	}
	return nil
}

func (bru *BattleRecordsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   battlerecords.Table,
			Columns: battlerecords.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: battlerecords.FieldID,
			},
		},
	}
	if ps := bru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bru.mutation.BattleFormat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: battlerecords.FieldBattleFormat,
		})
	}
	if value, ok := bru.mutation.SelfElectionPokemonId1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId1,
		})
	}
	if value, ok := bru.mutation.AddedSelfElectionPokemonId1(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId1,
		})
	}
	if value, ok := bru.mutation.SelfElectionPokemonId2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId2,
		})
	}
	if value, ok := bru.mutation.AddedSelfElectionPokemonId2(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId2,
		})
	}
	if value, ok := bru.mutation.SelfElectionPokemonId3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId3,
		})
	}
	if value, ok := bru.mutation.AddedSelfElectionPokemonId3(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId3,
		})
	}
	if value, ok := bru.mutation.SelfElectionPokemonId4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId4,
		})
	}
	if value, ok := bru.mutation.AddedSelfElectionPokemonId4(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId4,
		})
	}
	if bru.mutation.SelfElectionPokemonId4Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: battlerecords.FieldSelfElectionPokemonId4,
		})
	}
	if value, ok := bru.mutation.OpponentElectionPokemonId1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId1,
		})
	}
	if value, ok := bru.mutation.AddedOpponentElectionPokemonId1(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId1,
		})
	}
	if value, ok := bru.mutation.OpponentElectionPokemonId2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId2,
		})
	}
	if value, ok := bru.mutation.AddedOpponentElectionPokemonId2(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId2,
		})
	}
	if value, ok := bru.mutation.OpponentElectionPokemonId3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId3,
		})
	}
	if value, ok := bru.mutation.AddedOpponentElectionPokemonId3(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId3,
		})
	}
	if value, ok := bru.mutation.OpponentElectionPokemonId4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId4,
		})
	}
	if value, ok := bru.mutation.AddedOpponentElectionPokemonId4(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId4,
		})
	}
	if bru.mutation.UsePartyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   battlerecords.UsePartyTable,
			Columns: []string{battlerecords.UsePartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bru.mutation.UsePartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   battlerecords.UsePartyTable,
			Columns: []string{battlerecords.UsePartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bru.mutation.OpponentPartyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   battlerecords.OpponentPartyTable,
			Columns: []string{battlerecords.OpponentPartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battleopponentparty.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bru.mutation.OpponentPartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   battlerecords.OpponentPartyTable,
			Columns: []string{battlerecords.OpponentPartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battleopponentparty.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{battlerecords.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BattleRecordsUpdateOne is the builder for updating a single BattleRecords entity.
type BattleRecordsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BattleRecordsMutation
}

// SetPartyID sets the "party_id" field.
func (bruo *BattleRecordsUpdateOne) SetPartyID(i int) *BattleRecordsUpdateOne {
	bruo.mutation.SetPartyID(i)
	return bruo
}

// SetBattleFormat sets the "battle_format" field.
func (bruo *BattleRecordsUpdateOne) SetBattleFormat(pf property.BattleFormats) *BattleRecordsUpdateOne {
	bruo.mutation.SetBattleFormat(pf)
	return bruo
}

// SetBattleOpponentPartyID sets the "battle_opponent_party_id" field.
func (bruo *BattleRecordsUpdateOne) SetBattleOpponentPartyID(i int) *BattleRecordsUpdateOne {
	bruo.mutation.SetBattleOpponentPartyID(i)
	return bruo
}

// SetSelfElectionPokemonId1 sets the "self_election_pokemon_id1" field.
func (bruo *BattleRecordsUpdateOne) SetSelfElectionPokemonId1(i int) *BattleRecordsUpdateOne {
	bruo.mutation.ResetSelfElectionPokemonId1()
	bruo.mutation.SetSelfElectionPokemonId1(i)
	return bruo
}

// AddSelfElectionPokemonId1 adds i to the "self_election_pokemon_id1" field.
func (bruo *BattleRecordsUpdateOne) AddSelfElectionPokemonId1(i int) *BattleRecordsUpdateOne {
	bruo.mutation.AddSelfElectionPokemonId1(i)
	return bruo
}

// SetSelfElectionPokemonId2 sets the "self_election_pokemon_id2" field.
func (bruo *BattleRecordsUpdateOne) SetSelfElectionPokemonId2(i int) *BattleRecordsUpdateOne {
	bruo.mutation.ResetSelfElectionPokemonId2()
	bruo.mutation.SetSelfElectionPokemonId2(i)
	return bruo
}

// AddSelfElectionPokemonId2 adds i to the "self_election_pokemon_id2" field.
func (bruo *BattleRecordsUpdateOne) AddSelfElectionPokemonId2(i int) *BattleRecordsUpdateOne {
	bruo.mutation.AddSelfElectionPokemonId2(i)
	return bruo
}

// SetSelfElectionPokemonId3 sets the "self_election_pokemon_id3" field.
func (bruo *BattleRecordsUpdateOne) SetSelfElectionPokemonId3(i int) *BattleRecordsUpdateOne {
	bruo.mutation.ResetSelfElectionPokemonId3()
	bruo.mutation.SetSelfElectionPokemonId3(i)
	return bruo
}

// AddSelfElectionPokemonId3 adds i to the "self_election_pokemon_id3" field.
func (bruo *BattleRecordsUpdateOne) AddSelfElectionPokemonId3(i int) *BattleRecordsUpdateOne {
	bruo.mutation.AddSelfElectionPokemonId3(i)
	return bruo
}

// SetSelfElectionPokemonId4 sets the "self_election_pokemon_id4" field.
func (bruo *BattleRecordsUpdateOne) SetSelfElectionPokemonId4(i int) *BattleRecordsUpdateOne {
	bruo.mutation.ResetSelfElectionPokemonId4()
	bruo.mutation.SetSelfElectionPokemonId4(i)
	return bruo
}

// SetNillableSelfElectionPokemonId4 sets the "self_election_pokemon_id4" field if the given value is not nil.
func (bruo *BattleRecordsUpdateOne) SetNillableSelfElectionPokemonId4(i *int) *BattleRecordsUpdateOne {
	if i != nil {
		bruo.SetSelfElectionPokemonId4(*i)
	}
	return bruo
}

// AddSelfElectionPokemonId4 adds i to the "self_election_pokemon_id4" field.
func (bruo *BattleRecordsUpdateOne) AddSelfElectionPokemonId4(i int) *BattleRecordsUpdateOne {
	bruo.mutation.AddSelfElectionPokemonId4(i)
	return bruo
}

// ClearSelfElectionPokemonId4 clears the value of the "self_election_pokemon_id4" field.
func (bruo *BattleRecordsUpdateOne) ClearSelfElectionPokemonId4() *BattleRecordsUpdateOne {
	bruo.mutation.ClearSelfElectionPokemonId4()
	return bruo
}

// SetOpponentElectionPokemonId1 sets the "opponent_election_pokemon_id1" field.
func (bruo *BattleRecordsUpdateOne) SetOpponentElectionPokemonId1(i int) *BattleRecordsUpdateOne {
	bruo.mutation.ResetOpponentElectionPokemonId1()
	bruo.mutation.SetOpponentElectionPokemonId1(i)
	return bruo
}

// AddOpponentElectionPokemonId1 adds i to the "opponent_election_pokemon_id1" field.
func (bruo *BattleRecordsUpdateOne) AddOpponentElectionPokemonId1(i int) *BattleRecordsUpdateOne {
	bruo.mutation.AddOpponentElectionPokemonId1(i)
	return bruo
}

// SetOpponentElectionPokemonId2 sets the "opponent_election_pokemon_id2" field.
func (bruo *BattleRecordsUpdateOne) SetOpponentElectionPokemonId2(i int) *BattleRecordsUpdateOne {
	bruo.mutation.ResetOpponentElectionPokemonId2()
	bruo.mutation.SetOpponentElectionPokemonId2(i)
	return bruo
}

// AddOpponentElectionPokemonId2 adds i to the "opponent_election_pokemon_id2" field.
func (bruo *BattleRecordsUpdateOne) AddOpponentElectionPokemonId2(i int) *BattleRecordsUpdateOne {
	bruo.mutation.AddOpponentElectionPokemonId2(i)
	return bruo
}

// SetOpponentElectionPokemonId3 sets the "opponent_election_pokemon_id3" field.
func (bruo *BattleRecordsUpdateOne) SetOpponentElectionPokemonId3(i int) *BattleRecordsUpdateOne {
	bruo.mutation.ResetOpponentElectionPokemonId3()
	bruo.mutation.SetOpponentElectionPokemonId3(i)
	return bruo
}

// AddOpponentElectionPokemonId3 adds i to the "opponent_election_pokemon_id3" field.
func (bruo *BattleRecordsUpdateOne) AddOpponentElectionPokemonId3(i int) *BattleRecordsUpdateOne {
	bruo.mutation.AddOpponentElectionPokemonId3(i)
	return bruo
}

// SetOpponentElectionPokemonId4 sets the "opponent_election_pokemon_id4" field.
func (bruo *BattleRecordsUpdateOne) SetOpponentElectionPokemonId4(i int) *BattleRecordsUpdateOne {
	bruo.mutation.ResetOpponentElectionPokemonId4()
	bruo.mutation.SetOpponentElectionPokemonId4(i)
	return bruo
}

// AddOpponentElectionPokemonId4 adds i to the "opponent_election_pokemon_id4" field.
func (bruo *BattleRecordsUpdateOne) AddOpponentElectionPokemonId4(i int) *BattleRecordsUpdateOne {
	bruo.mutation.AddOpponentElectionPokemonId4(i)
	return bruo
}

// SetUsePartyID sets the "use_party" edge to the Party entity by ID.
func (bruo *BattleRecordsUpdateOne) SetUsePartyID(id int) *BattleRecordsUpdateOne {
	bruo.mutation.SetUsePartyID(id)
	return bruo
}

// SetUseParty sets the "use_party" edge to the Party entity.
func (bruo *BattleRecordsUpdateOne) SetUseParty(p *Party) *BattleRecordsUpdateOne {
	return bruo.SetUsePartyID(p.ID)
}

// SetOpponentPartyID sets the "opponent_party" edge to the BattleOpponentParty entity by ID.
func (bruo *BattleRecordsUpdateOne) SetOpponentPartyID(id int) *BattleRecordsUpdateOne {
	bruo.mutation.SetOpponentPartyID(id)
	return bruo
}

// SetOpponentParty sets the "opponent_party" edge to the BattleOpponentParty entity.
func (bruo *BattleRecordsUpdateOne) SetOpponentParty(b *BattleOpponentParty) *BattleRecordsUpdateOne {
	return bruo.SetOpponentPartyID(b.ID)
}

// Mutation returns the BattleRecordsMutation object of the builder.
func (bruo *BattleRecordsUpdateOne) Mutation() *BattleRecordsMutation {
	return bruo.mutation
}

// ClearUseParty clears the "use_party" edge to the Party entity.
func (bruo *BattleRecordsUpdateOne) ClearUseParty() *BattleRecordsUpdateOne {
	bruo.mutation.ClearUseParty()
	return bruo
}

// ClearOpponentParty clears the "opponent_party" edge to the BattleOpponentParty entity.
func (bruo *BattleRecordsUpdateOne) ClearOpponentParty() *BattleRecordsUpdateOne {
	bruo.mutation.ClearOpponentParty()
	return bruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bruo *BattleRecordsUpdateOne) Select(field string, fields ...string) *BattleRecordsUpdateOne {
	bruo.fields = append([]string{field}, fields...)
	return bruo
}

// Save executes the query and returns the updated BattleRecords entity.
func (bruo *BattleRecordsUpdateOne) Save(ctx context.Context) (*BattleRecords, error) {
	var (
		err  error
		node *BattleRecords
	)
	if len(bruo.hooks) == 0 {
		if err = bruo.check(); err != nil {
			return nil, err
		}
		node, err = bruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BattleRecordsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bruo.check(); err != nil {
				return nil, err
			}
			bruo.mutation = mutation
			node, err = bruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bruo.hooks) - 1; i >= 0; i-- {
			if bruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bruo *BattleRecordsUpdateOne) SaveX(ctx context.Context) *BattleRecords {
	node, err := bruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bruo *BattleRecordsUpdateOne) Exec(ctx context.Context) error {
	_, err := bruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bruo *BattleRecordsUpdateOne) ExecX(ctx context.Context) {
	if err := bruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bruo *BattleRecordsUpdateOne) check() error {
	if v, ok := bruo.mutation.PartyID(); ok {
		if err := battlerecords.PartyIDValidator(v); err != nil {
			return &ValidationError{Name: "party_id", err: fmt.Errorf("ent: validator failed for field \"party_id\": %w", err)}
		}
	}
	if v, ok := bruo.mutation.BattleOpponentPartyID(); ok {
		if err := battlerecords.BattleOpponentPartyIDValidator(v); err != nil {
			return &ValidationError{Name: "battle_opponent_party_id", err: fmt.Errorf("ent: validator failed for field \"battle_opponent_party_id\": %w", err)}
		}
	}
	if v, ok := bruo.mutation.SelfElectionPokemonId1(); ok {
		if err := battlerecords.SelfElectionPokemonId1Validator(v); err != nil {
			return &ValidationError{Name: "self_election_pokemon_id1", err: fmt.Errorf("ent: validator failed for field \"self_election_pokemon_id1\": %w", err)}
		}
	}
	if v, ok := bruo.mutation.SelfElectionPokemonId2(); ok {
		if err := battlerecords.SelfElectionPokemonId2Validator(v); err != nil {
			return &ValidationError{Name: "self_election_pokemon_id2", err: fmt.Errorf("ent: validator failed for field \"self_election_pokemon_id2\": %w", err)}
		}
	}
	if v, ok := bruo.mutation.SelfElectionPokemonId3(); ok {
		if err := battlerecords.SelfElectionPokemonId3Validator(v); err != nil {
			return &ValidationError{Name: "self_election_pokemon_id3", err: fmt.Errorf("ent: validator failed for field \"self_election_pokemon_id3\": %w", err)}
		}
	}
	if v, ok := bruo.mutation.SelfElectionPokemonId4(); ok {
		if err := battlerecords.SelfElectionPokemonId4Validator(v); err != nil {
			return &ValidationError{Name: "self_election_pokemon_id4", err: fmt.Errorf("ent: validator failed for field \"self_election_pokemon_id4\": %w", err)}
		}
	}
	if v, ok := bruo.mutation.OpponentElectionPokemonId1(); ok {
		if err := battlerecords.OpponentElectionPokemonId1Validator(v); err != nil {
			return &ValidationError{Name: "opponent_election_pokemon_id1", err: fmt.Errorf("ent: validator failed for field \"opponent_election_pokemon_id1\": %w", err)}
		}
	}
	if v, ok := bruo.mutation.OpponentElectionPokemonId2(); ok {
		if err := battlerecords.OpponentElectionPokemonId2Validator(v); err != nil {
			return &ValidationError{Name: "opponent_election_pokemon_id2", err: fmt.Errorf("ent: validator failed for field \"opponent_election_pokemon_id2\": %w", err)}
		}
	}
	if v, ok := bruo.mutation.OpponentElectionPokemonId3(); ok {
		if err := battlerecords.OpponentElectionPokemonId3Validator(v); err != nil {
			return &ValidationError{Name: "opponent_election_pokemon_id3", err: fmt.Errorf("ent: validator failed for field \"opponent_election_pokemon_id3\": %w", err)}
		}
	}
	if v, ok := bruo.mutation.OpponentElectionPokemonId4(); ok {
		if err := battlerecords.OpponentElectionPokemonId4Validator(v); err != nil {
			return &ValidationError{Name: "opponent_election_pokemon_id4", err: fmt.Errorf("ent: validator failed for field \"opponent_election_pokemon_id4\": %w", err)}
		}
	}
	if _, ok := bruo.mutation.UsePartyID(); bruo.mutation.UsePartyCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"use_party\"")
	}
	if _, ok := bruo.mutation.OpponentPartyID(); bruo.mutation.OpponentPartyCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"opponent_party\"")
	}
	return nil
}

func (bruo *BattleRecordsUpdateOne) sqlSave(ctx context.Context) (_node *BattleRecords, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   battlerecords.Table,
			Columns: battlerecords.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: battlerecords.FieldID,
			},
		},
	}
	id, ok := bruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing BattleRecords.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := bruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, battlerecords.FieldID)
		for _, f := range fields {
			if !battlerecords.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != battlerecords.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bruo.mutation.BattleFormat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: battlerecords.FieldBattleFormat,
		})
	}
	if value, ok := bruo.mutation.SelfElectionPokemonId1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId1,
		})
	}
	if value, ok := bruo.mutation.AddedSelfElectionPokemonId1(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId1,
		})
	}
	if value, ok := bruo.mutation.SelfElectionPokemonId2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId2,
		})
	}
	if value, ok := bruo.mutation.AddedSelfElectionPokemonId2(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId2,
		})
	}
	if value, ok := bruo.mutation.SelfElectionPokemonId3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId3,
		})
	}
	if value, ok := bruo.mutation.AddedSelfElectionPokemonId3(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId3,
		})
	}
	if value, ok := bruo.mutation.SelfElectionPokemonId4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId4,
		})
	}
	if value, ok := bruo.mutation.AddedSelfElectionPokemonId4(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldSelfElectionPokemonId4,
		})
	}
	if bruo.mutation.SelfElectionPokemonId4Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: battlerecords.FieldSelfElectionPokemonId4,
		})
	}
	if value, ok := bruo.mutation.OpponentElectionPokemonId1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId1,
		})
	}
	if value, ok := bruo.mutation.AddedOpponentElectionPokemonId1(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId1,
		})
	}
	if value, ok := bruo.mutation.OpponentElectionPokemonId2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId2,
		})
	}
	if value, ok := bruo.mutation.AddedOpponentElectionPokemonId2(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId2,
		})
	}
	if value, ok := bruo.mutation.OpponentElectionPokemonId3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId3,
		})
	}
	if value, ok := bruo.mutation.AddedOpponentElectionPokemonId3(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId3,
		})
	}
	if value, ok := bruo.mutation.OpponentElectionPokemonId4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId4,
		})
	}
	if value, ok := bruo.mutation.AddedOpponentElectionPokemonId4(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: battlerecords.FieldOpponentElectionPokemonId4,
		})
	}
	if bruo.mutation.UsePartyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   battlerecords.UsePartyTable,
			Columns: []string{battlerecords.UsePartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bruo.mutation.UsePartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   battlerecords.UsePartyTable,
			Columns: []string{battlerecords.UsePartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bruo.mutation.OpponentPartyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   battlerecords.OpponentPartyTable,
			Columns: []string{battlerecords.OpponentPartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battleopponentparty.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bruo.mutation.OpponentPartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   battlerecords.OpponentPartyTable,
			Columns: []string{battlerecords.OpponentPartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: battleopponentparty.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BattleRecords{config: bruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{battlerecords.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
