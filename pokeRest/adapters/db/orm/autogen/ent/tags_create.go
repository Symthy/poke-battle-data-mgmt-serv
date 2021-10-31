// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/party"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/tags"
)

// TagsCreate is the builder for creating a Tags entity.
type TagsCreate struct {
	config
	mutation *TagsMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TagsCreate) SetName(s string) *TagsCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetIsGenerationTag sets the "is_generation_tag" field.
func (tc *TagsCreate) SetIsGenerationTag(b bool) *TagsCreate {
	tc.mutation.SetIsGenerationTag(b)
	return tc
}

// SetNillableIsGenerationTag sets the "is_generation_tag" field if the given value is not nil.
func (tc *TagsCreate) SetNillableIsGenerationTag(b *bool) *TagsCreate {
	if b != nil {
		tc.SetIsGenerationTag(*b)
	}
	return tc
}

// SetIsSeasonTag sets the "is_season_tag" field.
func (tc *TagsCreate) SetIsSeasonTag(b bool) *TagsCreate {
	tc.mutation.SetIsSeasonTag(b)
	return tc
}

// SetNillableIsSeasonTag sets the "is_season_tag" field if the given value is not nil.
func (tc *TagsCreate) SetNillableIsSeasonTag(b *bool) *TagsCreate {
	if b != nil {
		tc.SetIsSeasonTag(*b)
	}
	return tc
}

// AddTagToPartyIDs adds the "tag_to_party" edge to the Party entity by IDs.
func (tc *TagsCreate) AddTagToPartyIDs(ids ...int) *TagsCreate {
	tc.mutation.AddTagToPartyIDs(ids...)
	return tc
}

// AddTagToParty adds the "tag_to_party" edges to the Party entity.
func (tc *TagsCreate) AddTagToParty(p ...*Party) *TagsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddTagToPartyIDs(ids...)
}

// Mutation returns the TagsMutation object of the builder.
func (tc *TagsCreate) Mutation() *TagsMutation {
	return tc.mutation
}

// Save creates the Tags in the database.
func (tc *TagsCreate) Save(ctx context.Context) (*Tags, error) {
	var (
		err  error
		node *Tags
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TagsCreate) SaveX(ctx context.Context) *Tags {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TagsCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TagsCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TagsCreate) defaults() {
	if _, ok := tc.mutation.IsGenerationTag(); !ok {
		v := tags.DefaultIsGenerationTag
		tc.mutation.SetIsGenerationTag(v)
	}
	if _, ok := tc.mutation.IsSeasonTag(); !ok {
		v := tags.DefaultIsSeasonTag
		tc.mutation.SetIsSeasonTag(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TagsCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := tags.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.IsGenerationTag(); !ok {
		return &ValidationError{Name: "is_generation_tag", err: errors.New(`ent: missing required field "is_generation_tag"`)}
	}
	if _, ok := tc.mutation.IsSeasonTag(); !ok {
		return &ValidationError{Name: "is_season_tag", err: errors.New(`ent: missing required field "is_season_tag"`)}
	}
	return nil
}

func (tc *TagsCreate) sqlSave(ctx context.Context) (*Tags, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TagsCreate) createSpec() (*Tags, *sqlgraph.CreateSpec) {
	var (
		_node = &Tags{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tags.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tags.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tags.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.IsGenerationTag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: tags.FieldIsGenerationTag,
		})
		_node.IsGenerationTag = value
	}
	if value, ok := tc.mutation.IsSeasonTag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: tags.FieldIsSeasonTag,
		})
		_node.IsSeasonTag = value
	}
	if nodes := tc.mutation.TagToPartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tags.TagToPartyTable,
			Columns: tags.TagToPartyPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TagsCreateBulk is the builder for creating many Tags entities in bulk.
type TagsCreateBulk struct {
	config
	builders []*TagsCreate
}

// Save creates the Tags entities in the database.
func (tcb *TagsCreateBulk) Save(ctx context.Context) ([]*Tags, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tags, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TagsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TagsCreateBulk) SaveX(ctx context.Context) []*Tags {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TagsCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TagsCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}