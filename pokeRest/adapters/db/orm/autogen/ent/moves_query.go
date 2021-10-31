// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/moves"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/predicate"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/trainedpokemondetails"
)

// MovesQuery is the builder for querying Moves entities.
type MovesQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Moves
	// eager-loading edges.
	withToTrainedPokemonMove1 *TrainedPokemonDetailsQuery
	withToTrainedPokemonMove2 *TrainedPokemonDetailsQuery
	withToTrainedPokemonMove3 *TrainedPokemonDetailsQuery
	withToTrainedPokemonMove4 *TrainedPokemonDetailsQuery
	withMoveToPokemon         *PokemonsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MovesQuery builder.
func (mq *MovesQuery) Where(ps ...predicate.Moves) *MovesQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MovesQuery) Limit(limit int) *MovesQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MovesQuery) Offset(offset int) *MovesQuery {
	mq.offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MovesQuery) Unique(unique bool) *MovesQuery {
	mq.unique = &unique
	return mq
}

// Order adds an order step to the query.
func (mq *MovesQuery) Order(o ...OrderFunc) *MovesQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryToTrainedPokemonMove1 chains the current query on the "to_trained_pokemon_move1" edge.
func (mq *MovesQuery) QueryToTrainedPokemonMove1() *TrainedPokemonDetailsQuery {
	query := &TrainedPokemonDetailsQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moves.Table, moves.FieldID, selector),
			sqlgraph.To(trainedpokemondetails.Table, trainedpokemondetails.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, moves.ToTrainedPokemonMove1Table, moves.ToTrainedPokemonMove1Column),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryToTrainedPokemonMove2 chains the current query on the "to_trained_pokemon_move2" edge.
func (mq *MovesQuery) QueryToTrainedPokemonMove2() *TrainedPokemonDetailsQuery {
	query := &TrainedPokemonDetailsQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moves.Table, moves.FieldID, selector),
			sqlgraph.To(trainedpokemondetails.Table, trainedpokemondetails.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, moves.ToTrainedPokemonMove2Table, moves.ToTrainedPokemonMove2Column),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryToTrainedPokemonMove3 chains the current query on the "to_trained_pokemon_move3" edge.
func (mq *MovesQuery) QueryToTrainedPokemonMove3() *TrainedPokemonDetailsQuery {
	query := &TrainedPokemonDetailsQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moves.Table, moves.FieldID, selector),
			sqlgraph.To(trainedpokemondetails.Table, trainedpokemondetails.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, moves.ToTrainedPokemonMove3Table, moves.ToTrainedPokemonMove3Column),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryToTrainedPokemonMove4 chains the current query on the "to_trained_pokemon_move4" edge.
func (mq *MovesQuery) QueryToTrainedPokemonMove4() *TrainedPokemonDetailsQuery {
	query := &TrainedPokemonDetailsQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moves.Table, moves.FieldID, selector),
			sqlgraph.To(trainedpokemondetails.Table, trainedpokemondetails.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, moves.ToTrainedPokemonMove4Table, moves.ToTrainedPokemonMove4Column),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMoveToPokemon chains the current query on the "move_to_pokemon" edge.
func (mq *MovesQuery) QueryMoveToPokemon() *PokemonsQuery {
	query := &PokemonsQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moves.Table, moves.FieldID, selector),
			sqlgraph.To(pokemons.Table, pokemons.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, moves.MoveToPokemonTable, moves.MoveToPokemonPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Moves entity from the query.
// Returns a *NotFoundError when no Moves was found.
func (mq *MovesQuery) First(ctx context.Context) (*Moves, error) {
	nodes, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{moves.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MovesQuery) FirstX(ctx context.Context) *Moves {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Moves ID from the query.
// Returns a *NotFoundError when no Moves ID was found.
func (mq *MovesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{moves.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MovesQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Moves entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Moves entity is not found.
// Returns a *NotFoundError when no Moves entities are found.
func (mq *MovesQuery) Only(ctx context.Context) (*Moves, error) {
	nodes, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{moves.Label}
	default:
		return nil, &NotSingularError{moves.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MovesQuery) OnlyX(ctx context.Context) *Moves {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Moves ID in the query.
// Returns a *NotSingularError when exactly one Moves ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mq *MovesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{moves.Label}
	default:
		err = &NotSingularError{moves.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MovesQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MovesSlice.
func (mq *MovesQuery) All(ctx context.Context) ([]*Moves, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MovesQuery) AllX(ctx context.Context) []*Moves {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Moves IDs.
func (mq *MovesQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(moves.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MovesQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MovesQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MovesQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MovesQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MovesQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MovesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MovesQuery) Clone() *MovesQuery {
	if mq == nil {
		return nil
	}
	return &MovesQuery{
		config:                    mq.config,
		limit:                     mq.limit,
		offset:                    mq.offset,
		order:                     append([]OrderFunc{}, mq.order...),
		predicates:                append([]predicate.Moves{}, mq.predicates...),
		withToTrainedPokemonMove1: mq.withToTrainedPokemonMove1.Clone(),
		withToTrainedPokemonMove2: mq.withToTrainedPokemonMove2.Clone(),
		withToTrainedPokemonMove3: mq.withToTrainedPokemonMove3.Clone(),
		withToTrainedPokemonMove4: mq.withToTrainedPokemonMove4.Clone(),
		withMoveToPokemon:         mq.withMoveToPokemon.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithToTrainedPokemonMove1 tells the query-builder to eager-load the nodes that are connected to
// the "to_trained_pokemon_move1" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MovesQuery) WithToTrainedPokemonMove1(opts ...func(*TrainedPokemonDetailsQuery)) *MovesQuery {
	query := &TrainedPokemonDetailsQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withToTrainedPokemonMove1 = query
	return mq
}

// WithToTrainedPokemonMove2 tells the query-builder to eager-load the nodes that are connected to
// the "to_trained_pokemon_move2" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MovesQuery) WithToTrainedPokemonMove2(opts ...func(*TrainedPokemonDetailsQuery)) *MovesQuery {
	query := &TrainedPokemonDetailsQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withToTrainedPokemonMove2 = query
	return mq
}

// WithToTrainedPokemonMove3 tells the query-builder to eager-load the nodes that are connected to
// the "to_trained_pokemon_move3" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MovesQuery) WithToTrainedPokemonMove3(opts ...func(*TrainedPokemonDetailsQuery)) *MovesQuery {
	query := &TrainedPokemonDetailsQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withToTrainedPokemonMove3 = query
	return mq
}

// WithToTrainedPokemonMove4 tells the query-builder to eager-load the nodes that are connected to
// the "to_trained_pokemon_move4" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MovesQuery) WithToTrainedPokemonMove4(opts ...func(*TrainedPokemonDetailsQuery)) *MovesQuery {
	query := &TrainedPokemonDetailsQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withToTrainedPokemonMove4 = query
	return mq
}

// WithMoveToPokemon tells the query-builder to eager-load the nodes that are connected to
// the "move_to_pokemon" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MovesQuery) WithMoveToPokemon(opts ...func(*PokemonsQuery)) *MovesQuery {
	query := &PokemonsQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMoveToPokemon = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Moves.Query().
//		GroupBy(moves.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mq *MovesQuery) GroupBy(field string, fields ...string) *MovesGroupBy {
	group := &MovesGroupBy{config: mq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Moves.Query().
//		Select(moves.FieldName).
//		Scan(ctx, &v)
//
func (mq *MovesQuery) Select(fields ...string) *MovesSelect {
	mq.fields = append(mq.fields, fields...)
	return &MovesSelect{MovesQuery: mq}
}

func (mq *MovesQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mq.fields {
		if !moves.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MovesQuery) sqlAll(ctx context.Context) ([]*Moves, error) {
	var (
		nodes       = []*Moves{}
		_spec       = mq.querySpec()
		loadedTypes = [5]bool{
			mq.withToTrainedPokemonMove1 != nil,
			mq.withToTrainedPokemonMove2 != nil,
			mq.withToTrainedPokemonMove3 != nil,
			mq.withToTrainedPokemonMove4 != nil,
			mq.withMoveToPokemon != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Moves{config: mq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mq.withToTrainedPokemonMove1; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Moves)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ToTrainedPokemonMove1 = []*TrainedPokemonDetails{}
		}
		query.Where(predicate.TrainedPokemonDetails(func(s *sql.Selector) {
			s.Where(sql.InValues(moves.ToTrainedPokemonMove1Column, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.MoveId1
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "move_id1" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.ToTrainedPokemonMove1 = append(node.Edges.ToTrainedPokemonMove1, n)
		}
	}

	if query := mq.withToTrainedPokemonMove2; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Moves)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ToTrainedPokemonMove2 = []*TrainedPokemonDetails{}
		}
		query.Where(predicate.TrainedPokemonDetails(func(s *sql.Selector) {
			s.Where(sql.InValues(moves.ToTrainedPokemonMove2Column, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.MoveId2
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "move_id2" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.ToTrainedPokemonMove2 = append(node.Edges.ToTrainedPokemonMove2, n)
		}
	}

	if query := mq.withToTrainedPokemonMove3; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Moves)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ToTrainedPokemonMove3 = []*TrainedPokemonDetails{}
		}
		query.Where(predicate.TrainedPokemonDetails(func(s *sql.Selector) {
			s.Where(sql.InValues(moves.ToTrainedPokemonMove3Column, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.MoveId3
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "move_id3" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.ToTrainedPokemonMove3 = append(node.Edges.ToTrainedPokemonMove3, n)
		}
	}

	if query := mq.withToTrainedPokemonMove4; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Moves)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ToTrainedPokemonMove4 = []*TrainedPokemonDetails{}
		}
		query.Where(predicate.TrainedPokemonDetails(func(s *sql.Selector) {
			s.Where(sql.InValues(moves.ToTrainedPokemonMove4Column, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.MoveId4
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "move_id4" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.ToTrainedPokemonMove4 = append(node.Edges.ToTrainedPokemonMove4, n)
		}
	}

	if query := mq.withMoveToPokemon; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Moves, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.MoveToPokemon = []*Pokemons{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Moves)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   moves.MoveToPokemonTable,
				Columns: moves.MoveToPokemonPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(moves.MoveToPokemonPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, mq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "move_to_pokemon": %w`, err)
		}
		query.Where(pokemons.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "move_to_pokemon" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.MoveToPokemon = append(nodes[i].Edges.MoveToPokemon, n)
			}
		}
	}

	return nodes, nil
}

func (mq *MovesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MovesQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mq *MovesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   moves.Table,
			Columns: moves.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moves.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if unique := mq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, moves.FieldID)
		for i := range fields {
			if fields[i] != moves.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MovesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(moves.Table)
	columns := mq.fields
	if len(columns) == 0 {
		columns = moves.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MovesGroupBy is the group-by builder for Moves entities.
type MovesGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MovesGroupBy) Aggregate(fns ...AggregateFunc) *MovesGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mgb *MovesGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mgb *MovesGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MovesGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MovesGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mgb *MovesGroupBy) StringsX(ctx context.Context) []string {
	v, err := mgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MovesGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moves.Label}
	default:
		err = fmt.Errorf("ent: MovesGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mgb *MovesGroupBy) StringX(ctx context.Context) string {
	v, err := mgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MovesGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MovesGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mgb *MovesGroupBy) IntsX(ctx context.Context) []int {
	v, err := mgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MovesGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moves.Label}
	default:
		err = fmt.Errorf("ent: MovesGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mgb *MovesGroupBy) IntX(ctx context.Context) int {
	v, err := mgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MovesGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MovesGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mgb *MovesGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MovesGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moves.Label}
	default:
		err = fmt.Errorf("ent: MovesGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mgb *MovesGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mgb *MovesGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MovesGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mgb *MovesGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mgb *MovesGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moves.Label}
	default:
		err = fmt.Errorf("ent: MovesGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mgb *MovesGroupBy) BoolX(ctx context.Context) bool {
	v, err := mgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mgb *MovesGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mgb.fields {
		if !moves.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MovesGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql.Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
		for _, f := range mgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mgb.fields...)...)
}

// MovesSelect is the builder for selecting fields of Moves entities.
type MovesSelect struct {
	*MovesQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MovesSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	ms.sql = ms.MovesQuery.sqlQuery(ctx)
	return ms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ms *MovesSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ms *MovesSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MovesSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ms *MovesSelect) StringsX(ctx context.Context) []string {
	v, err := ms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ms *MovesSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moves.Label}
	default:
		err = fmt.Errorf("ent: MovesSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ms *MovesSelect) StringX(ctx context.Context) string {
	v, err := ms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ms *MovesSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MovesSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ms *MovesSelect) IntsX(ctx context.Context) []int {
	v, err := ms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ms *MovesSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moves.Label}
	default:
		err = fmt.Errorf("ent: MovesSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ms *MovesSelect) IntX(ctx context.Context) int {
	v, err := ms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ms *MovesSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MovesSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ms *MovesSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ms *MovesSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moves.Label}
	default:
		err = fmt.Errorf("ent: MovesSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ms *MovesSelect) Float64X(ctx context.Context) float64 {
	v, err := ms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ms *MovesSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MovesSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ms *MovesSelect) BoolsX(ctx context.Context) []bool {
	v, err := ms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ms *MovesSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moves.Label}
	default:
		err = fmt.Errorf("ent: MovesSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ms *MovesSelect) BoolX(ctx context.Context) bool {
	v, err := ms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ms *MovesSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ms.sql.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}