// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/predicate"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"
)

// CityQuery is the builder for querying City entities.
type CityQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.City
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (cq *CityQuery) Where(ps ...predicate.City) *CityQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CityQuery) Limit(limit int) *CityQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CityQuery) Offset(offset int) *CityQuery {
	cq.offset = &offset
	return cq
}

// Order adds an order step to the query.
func (cq *CityQuery) Order(o ...Order) *CityQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryStreets chains the current query on the streets edge.
func (cq *CityQuery) QueryStreets() *StreetQuery {
	query := &StreetQuery{config: cq.config}
	t1 := sql.Table(street.Table)
	t2 := cq.sqlQuery()
	t2.Select(t2.C(city.FieldID))
	query.sql = sql.Select().
		From(t1).
		Join(t2).
		On(t1.C(city.StreetsColumn), t2.C(city.FieldID))
	return query
}

// First returns the first City entity in the query. Returns *ErrNotFound when no city was found.
func (cq *CityQuery) First(ctx context.Context) (*City, error) {
	cs, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(cs) == 0 {
		return nil, &ErrNotFound{city.Label}
	}
	return cs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CityQuery) FirstX(ctx context.Context) *City {
	c, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return c
}

// FirstID returns the first City id in the query. Returns *ErrNotFound when no id was found.
func (cq *CityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{city.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (cq *CityQuery) FirstXID(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only City entity in the query, returns an error if not exactly one entity was returned.
func (cq *CityQuery) Only(ctx context.Context) (*City, error) {
	cs, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(cs) {
	case 1:
		return cs[0], nil
	case 0:
		return nil, &ErrNotFound{city.Label}
	default:
		return nil, &ErrNotSingular{city.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CityQuery) OnlyX(ctx context.Context) *City {
	c, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// OnlyID returns the only City id in the query, returns an error if not exactly one id was returned.
func (cq *CityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{city.Label}
	default:
		err = &ErrNotSingular{city.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (cq *CityQuery) OnlyXID(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cities.
func (cq *CityQuery) All(ctx context.Context) ([]*City, error) {
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CityQuery) AllX(ctx context.Context) []*City {
	cs, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return cs
}

// IDs executes the query and returns a list of City ids.
func (cq *CityQuery) IDs(ctx context.Context) ([]int, error) {
	return cq.sqlIDs(ctx)
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CityQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CityQuery) Count(ctx context.Context) (int, error) {
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CityQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CityQuery) Exist(ctx context.Context) (bool, error) {
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CityQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CityQuery) Clone() *CityQuery {
	return &CityQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]Order{}, cq.order...),
		unique:     append([]string{}, cq.unique...),
		predicates: append([]predicate.City{}, cq.predicates...),
		// clone intermediate queries.
		sql: cq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.City.Query().
//		GroupBy(city.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *CityQuery) GroupBy(field string, fields ...string) *CityGroupBy {
	group := &CityGroupBy{config: cq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = cq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.City.Query().
//		Select(city.FieldName).
//		Scan(ctx, &v)
//
func (cq *CityQuery) Select(field string, fields ...string) *CitySelect {
	selector := &CitySelect{config: cq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = cq.sqlQuery()
	return selector
}

func (cq *CityQuery) sqlAll(ctx context.Context) ([]*City, error) {
	rows := &sql.Rows{}
	selector := cq.sqlQuery()
	if unique := cq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := cq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var cs Cities
	if err := cs.FromRows(rows); err != nil {
		return nil, err
	}
	cs.config(cq.config)
	return cs, nil
}

func (cq *CityQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := cq.sqlQuery()
	unique := []string{city.FieldID}
	if len(cq.unique) > 0 {
		unique = cq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := cq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("ent: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("ent: failed reading count: %v", err)
	}
	return n, nil
}

func (cq *CityQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (cq *CityQuery) sqlIDs(ctx context.Context) ([]int, error) {
	vs, err := cq.sqlAll(ctx)
	if err != nil {
		return nil, err
	}
	var ids []int
	for _, v := range vs {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

func (cq *CityQuery) sqlQuery() *sql.Selector {
	t1 := sql.Table(city.Table)
	selector := sql.Select(t1.Columns(city.Columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(city.Columns...)...)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt64)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CityGroupBy is the builder for group-by City entities.
type CityGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CityGroupBy) Aggregate(fns ...Aggregate) *CityGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scan the result into the given value.
func (cgb *CityGroupBy) Scan(ctx context.Context, v interface{}) error {
	return cgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cgb *CityGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (cgb *CityGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CityGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cgb *CityGroupBy) StringsX(ctx context.Context) []string {
	v, err := cgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (cgb *CityGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CityGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cgb *CityGroupBy) IntsX(ctx context.Context) []int {
	v, err := cgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (cgb *CityGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CityGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cgb *CityGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (cgb *CityGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CityGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cgb *CityGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cgb *CityGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cgb.sqlQuery().Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *CityGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql
	columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
	columns = append(columns, cgb.fields...)
	for _, fn := range cgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(cgb.fields...)
}

// CitySelect is the builder for select fields of City entities.
type CitySelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (cs *CitySelect) Scan(ctx context.Context, v interface{}) error {
	return cs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cs *CitySelect) ScanX(ctx context.Context, v interface{}) {
	if err := cs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (cs *CitySelect) Strings(ctx context.Context) ([]string, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CitySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cs *CitySelect) StringsX(ctx context.Context) []string {
	v, err := cs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (cs *CitySelect) Ints(ctx context.Context) ([]int, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CitySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cs *CitySelect) IntsX(ctx context.Context) []int {
	v, err := cs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (cs *CitySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CitySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cs *CitySelect) Float64sX(ctx context.Context) []float64 {
	v, err := cs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (cs *CitySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CitySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cs *CitySelect) BoolsX(ctx context.Context) []bool {
	v, err := cs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cs *CitySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sqlQuery().Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cs *CitySelect) sqlQuery() sql.Querier {
	view := "city_view"
	return sql.Select(cs.fields...).From(cs.sql.As(view))
}
