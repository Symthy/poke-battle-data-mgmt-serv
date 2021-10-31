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
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/abilities"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/forms"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/moves"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/predicate"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/autogen/ent/trainedpokemons"
)

// PokemonsQuery is the builder for querying Pokemons entities.
type PokemonsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Pokemons
	// eager-loading edges.
	withAbility1         *AbilitiesQuery
	withAbility2         *AbilitiesQuery
	withHiddenAbility    *AbilitiesQuery
	withForm             *FormsQuery
	withToTrainedPokemon *TrainedPokemonsQuery
	withPokemonToMove    *MovesQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PokemonsQuery builder.
func (pq *PokemonsQuery) Where(ps ...predicate.Pokemons) *PokemonsQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PokemonsQuery) Limit(limit int) *PokemonsQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PokemonsQuery) Offset(offset int) *PokemonsQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PokemonsQuery) Unique(unique bool) *PokemonsQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *PokemonsQuery) Order(o ...OrderFunc) *PokemonsQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryAbility1 chains the current query on the "ability1" edge.
func (pq *PokemonsQuery) QueryAbility1() *AbilitiesQuery {
	query := &AbilitiesQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pokemons.Table, pokemons.FieldID, selector),
			sqlgraph.To(abilities.Table, abilities.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, pokemons.Ability1Table, pokemons.Ability1Column),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAbility2 chains the current query on the "ability2" edge.
func (pq *PokemonsQuery) QueryAbility2() *AbilitiesQuery {
	query := &AbilitiesQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pokemons.Table, pokemons.FieldID, selector),
			sqlgraph.To(abilities.Table, abilities.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, pokemons.Ability2Table, pokemons.Ability2Column),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHiddenAbility chains the current query on the "hidden_ability" edge.
func (pq *PokemonsQuery) QueryHiddenAbility() *AbilitiesQuery {
	query := &AbilitiesQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pokemons.Table, pokemons.FieldID, selector),
			sqlgraph.To(abilities.Table, abilities.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, pokemons.HiddenAbilityTable, pokemons.HiddenAbilityColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryForm chains the current query on the "form" edge.
func (pq *PokemonsQuery) QueryForm() *FormsQuery {
	query := &FormsQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pokemons.Table, pokemons.FieldID, selector),
			sqlgraph.To(forms.Table, forms.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, pokemons.FormTable, pokemons.FormColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryToTrainedPokemon chains the current query on the "to_trained_pokemon" edge.
func (pq *PokemonsQuery) QueryToTrainedPokemon() *TrainedPokemonsQuery {
	query := &TrainedPokemonsQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pokemons.Table, pokemons.FieldID, selector),
			sqlgraph.To(trainedpokemons.Table, trainedpokemons.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, pokemons.ToTrainedPokemonTable, pokemons.ToTrainedPokemonColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPokemonToMove chains the current query on the "pokemon_to_move" edge.
func (pq *PokemonsQuery) QueryPokemonToMove() *MovesQuery {
	query := &MovesQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pokemons.Table, pokemons.FieldID, selector),
			sqlgraph.To(moves.Table, moves.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, pokemons.PokemonToMoveTable, pokemons.PokemonToMovePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Pokemons entity from the query.
// Returns a *NotFoundError when no Pokemons was found.
func (pq *PokemonsQuery) First(ctx context.Context) (*Pokemons, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pokemons.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PokemonsQuery) FirstX(ctx context.Context) *Pokemons {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pokemons ID from the query.
// Returns a *NotFoundError when no Pokemons ID was found.
func (pq *PokemonsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pokemons.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PokemonsQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pokemons entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Pokemons entity is not found.
// Returns a *NotFoundError when no Pokemons entities are found.
func (pq *PokemonsQuery) Only(ctx context.Context) (*Pokemons, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pokemons.Label}
	default:
		return nil, &NotSingularError{pokemons.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PokemonsQuery) OnlyX(ctx context.Context) *Pokemons {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pokemons ID in the query.
// Returns a *NotSingularError when exactly one Pokemons ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pq *PokemonsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pokemons.Label}
	default:
		err = &NotSingularError{pokemons.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PokemonsQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PokemonsSlice.
func (pq *PokemonsQuery) All(ctx context.Context) ([]*Pokemons, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PokemonsQuery) AllX(ctx context.Context) []*Pokemons {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pokemons IDs.
func (pq *PokemonsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(pokemons.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PokemonsQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PokemonsQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PokemonsQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PokemonsQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PokemonsQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PokemonsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PokemonsQuery) Clone() *PokemonsQuery {
	if pq == nil {
		return nil
	}
	return &PokemonsQuery{
		config:               pq.config,
		limit:                pq.limit,
		offset:               pq.offset,
		order:                append([]OrderFunc{}, pq.order...),
		predicates:           append([]predicate.Pokemons{}, pq.predicates...),
		withAbility1:         pq.withAbility1.Clone(),
		withAbility2:         pq.withAbility2.Clone(),
		withHiddenAbility:    pq.withHiddenAbility.Clone(),
		withForm:             pq.withForm.Clone(),
		withToTrainedPokemon: pq.withToTrainedPokemon.Clone(),
		withPokemonToMove:    pq.withPokemonToMove.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithAbility1 tells the query-builder to eager-load the nodes that are connected to
// the "ability1" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PokemonsQuery) WithAbility1(opts ...func(*AbilitiesQuery)) *PokemonsQuery {
	query := &AbilitiesQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withAbility1 = query
	return pq
}

// WithAbility2 tells the query-builder to eager-load the nodes that are connected to
// the "ability2" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PokemonsQuery) WithAbility2(opts ...func(*AbilitiesQuery)) *PokemonsQuery {
	query := &AbilitiesQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withAbility2 = query
	return pq
}

// WithHiddenAbility tells the query-builder to eager-load the nodes that are connected to
// the "hidden_ability" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PokemonsQuery) WithHiddenAbility(opts ...func(*AbilitiesQuery)) *PokemonsQuery {
	query := &AbilitiesQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withHiddenAbility = query
	return pq
}

// WithForm tells the query-builder to eager-load the nodes that are connected to
// the "form" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PokemonsQuery) WithForm(opts ...func(*FormsQuery)) *PokemonsQuery {
	query := &FormsQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withForm = query
	return pq
}

// WithToTrainedPokemon tells the query-builder to eager-load the nodes that are connected to
// the "to_trained_pokemon" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PokemonsQuery) WithToTrainedPokemon(opts ...func(*TrainedPokemonsQuery)) *PokemonsQuery {
	query := &TrainedPokemonsQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withToTrainedPokemon = query
	return pq
}

// WithPokemonToMove tells the query-builder to eager-load the nodes that are connected to
// the "pokemon_to_move" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PokemonsQuery) WithPokemonToMove(opts ...func(*MovesQuery)) *PokemonsQuery {
	query := &MovesQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withPokemonToMove = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PokedexNo int `json:"pokedex_no,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Pokemons.Query().
//		GroupBy(pokemons.FieldPokedexNo).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PokemonsQuery) GroupBy(field string, fields ...string) *PokemonsGroupBy {
	group := &PokemonsGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PokedexNo int `json:"pokedex_no,omitempty"`
//	}
//
//	client.Pokemons.Query().
//		Select(pokemons.FieldPokedexNo).
//		Scan(ctx, &v)
//
func (pq *PokemonsQuery) Select(fields ...string) *PokemonsSelect {
	pq.fields = append(pq.fields, fields...)
	return &PokemonsSelect{PokemonsQuery: pq}
}

func (pq *PokemonsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !pokemons.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PokemonsQuery) sqlAll(ctx context.Context) ([]*Pokemons, error) {
	var (
		nodes       = []*Pokemons{}
		_spec       = pq.querySpec()
		loadedTypes = [6]bool{
			pq.withAbility1 != nil,
			pq.withAbility2 != nil,
			pq.withHiddenAbility != nil,
			pq.withForm != nil,
			pq.withToTrainedPokemon != nil,
			pq.withPokemonToMove != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Pokemons{config: pq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withAbility1; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Pokemons)
		for i := range nodes {
			fk := nodes[i].AbilityId1
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(abilities.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "ability_id1" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Ability1 = n
			}
		}
	}

	if query := pq.withAbility2; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Pokemons)
		for i := range nodes {
			fk := nodes[i].AbilityId2
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(abilities.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "ability_id2" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Ability2 = n
			}
		}
	}

	if query := pq.withHiddenAbility; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Pokemons)
		for i := range nodes {
			fk := nodes[i].HiddenAbilityID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(abilities.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "hidden_ability_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.HiddenAbility = n
			}
		}
	}

	if query := pq.withForm; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Pokemons)
		for i := range nodes {
			fk := nodes[i].FormNo
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(forms.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "form_no" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Form = n
			}
		}
	}

	if query := pq.withToTrainedPokemon; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Pokemons)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ToTrainedPokemon = []*TrainedPokemons{}
		}
		query.Where(predicate.TrainedPokemons(func(s *sql.Selector) {
			s.Where(sql.InValues(pokemons.ToTrainedPokemonColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.PokemonID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "pokemon_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.ToTrainedPokemon = append(node.Edges.ToTrainedPokemon, n)
		}
	}

	if query := pq.withPokemonToMove; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Pokemons, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.PokemonToMove = []*Moves{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Pokemons)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   pokemons.PokemonToMoveTable,
				Columns: pokemons.PokemonToMovePrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(pokemons.PokemonToMovePrimaryKey[1], fks...))
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
		if err := sqlgraph.QueryEdges(ctx, pq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "pokemon_to_move": %w`, err)
		}
		query.Where(moves.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "pokemon_to_move" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.PokemonToMove = append(nodes[i].Edges.PokemonToMove, n)
			}
		}
	}

	return nodes, nil
}

func (pq *PokemonsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PokemonsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pq *PokemonsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pokemons.Table,
			Columns: pokemons.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pokemons.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pokemons.FieldID)
		for i := range fields {
			if fields[i] != pokemons.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PokemonsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pokemons.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = pokemons.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PokemonsGroupBy is the group-by builder for Pokemons entities.
type PokemonsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PokemonsGroupBy) Aggregate(fns ...AggregateFunc) *PokemonsGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PokemonsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PokemonsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PokemonsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PokemonsGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemons.Label}
	default:
		err = fmt.Errorf("ent: PokemonsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *PokemonsGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PokemonsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PokemonsGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemons.Label}
	default:
		err = fmt.Errorf("ent: PokemonsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *PokemonsGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PokemonsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PokemonsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemons.Label}
	default:
		err = fmt.Errorf("ent: PokemonsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *PokemonsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PokemonsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PokemonsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PokemonsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemons.Label}
	default:
		err = fmt.Errorf("ent: PokemonsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *PokemonsGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PokemonsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pgb.fields {
		if !pokemons.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PokemonsGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// PokemonsSelect is the builder for selecting fields of Pokemons entities.
type PokemonsSelect struct {
	*PokemonsQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PokemonsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.PokemonsQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PokemonsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ps *PokemonsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PokemonsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PokemonsSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ps *PokemonsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemons.Label}
	default:
		err = fmt.Errorf("ent: PokemonsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *PokemonsSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ps *PokemonsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PokemonsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PokemonsSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ps *PokemonsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemons.Label}
	default:
		err = fmt.Errorf("ent: PokemonsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *PokemonsSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ps *PokemonsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PokemonsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PokemonsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ps *PokemonsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemons.Label}
	default:
		err = fmt.Errorf("ent: PokemonsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *PokemonsSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ps *PokemonsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PokemonsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PokemonsSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ps *PokemonsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pokemons.Label}
	default:
		err = fmt.Errorf("ent: PokemonsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *PokemonsSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PokemonsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}