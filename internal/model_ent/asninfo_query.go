// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/asninfo"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/registrar"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scanjob"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/whois"
)

// ASNInfoQuery is the builder for querying ASNInfo entities.
type ASNInfoQuery struct {
	config
	ctx           *QueryContext
	order         []asninfo.OrderOption
	inters        []Interceptor
	predicates    []predicate.ASNInfo
	withScan      *ScanJobQuery
	withIpaddress *IPAddressQuery
	withRegistrar *RegistrarQuery
	withWhois     *WhoisQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ASNInfoQuery builder.
func (aiq *ASNInfoQuery) Where(ps ...predicate.ASNInfo) *ASNInfoQuery {
	aiq.predicates = append(aiq.predicates, ps...)
	return aiq
}

// Limit the number of records to be returned by this query.
func (aiq *ASNInfoQuery) Limit(limit int) *ASNInfoQuery {
	aiq.ctx.Limit = &limit
	return aiq
}

// Offset to start from.
func (aiq *ASNInfoQuery) Offset(offset int) *ASNInfoQuery {
	aiq.ctx.Offset = &offset
	return aiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aiq *ASNInfoQuery) Unique(unique bool) *ASNInfoQuery {
	aiq.ctx.Unique = &unique
	return aiq
}

// Order specifies how the records should be ordered.
func (aiq *ASNInfoQuery) Order(o ...asninfo.OrderOption) *ASNInfoQuery {
	aiq.order = append(aiq.order, o...)
	return aiq
}

// QueryScan chains the current query on the "scan" edge.
func (aiq *ASNInfoQuery) QueryScan() *ScanJobQuery {
	query := (&ScanJobClient{config: aiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asninfo.Table, asninfo.FieldID, selector),
			sqlgraph.To(scanjob.Table, scanjob.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, asninfo.ScanTable, asninfo.ScanPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIpaddress chains the current query on the "ipaddress" edge.
func (aiq *ASNInfoQuery) QueryIpaddress() *IPAddressQuery {
	query := (&IPAddressClient{config: aiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asninfo.Table, asninfo.FieldID, selector),
			sqlgraph.To(ipaddress.Table, ipaddress.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, asninfo.IpaddressTable, asninfo.IpaddressPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRegistrar chains the current query on the "registrar" edge.
func (aiq *ASNInfoQuery) QueryRegistrar() *RegistrarQuery {
	query := (&RegistrarClient{config: aiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asninfo.Table, asninfo.FieldID, selector),
			sqlgraph.To(registrar.Table, registrar.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, asninfo.RegistrarTable, asninfo.RegistrarPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWhois chains the current query on the "whois" edge.
func (aiq *ASNInfoQuery) QueryWhois() *WhoisQuery {
	query := (&WhoisClient{config: aiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asninfo.Table, asninfo.FieldID, selector),
			sqlgraph.To(whois.Table, whois.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, asninfo.WhoisTable, asninfo.WhoisPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ASNInfo entity from the query.
// Returns a *NotFoundError when no ASNInfo was found.
func (aiq *ASNInfoQuery) First(ctx context.Context) (*ASNInfo, error) {
	nodes, err := aiq.Limit(1).All(setContextOp(ctx, aiq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asninfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aiq *ASNInfoQuery) FirstX(ctx context.Context) *ASNInfo {
	node, err := aiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ASNInfo ID from the query.
// Returns a *NotFoundError when no ASNInfo ID was found.
func (aiq *ASNInfoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aiq.Limit(1).IDs(setContextOp(ctx, aiq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asninfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aiq *ASNInfoQuery) FirstIDX(ctx context.Context) int {
	id, err := aiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ASNInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ASNInfo entity is found.
// Returns a *NotFoundError when no ASNInfo entities are found.
func (aiq *ASNInfoQuery) Only(ctx context.Context) (*ASNInfo, error) {
	nodes, err := aiq.Limit(2).All(setContextOp(ctx, aiq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asninfo.Label}
	default:
		return nil, &NotSingularError{asninfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aiq *ASNInfoQuery) OnlyX(ctx context.Context) *ASNInfo {
	node, err := aiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ASNInfo ID in the query.
// Returns a *NotSingularError when more than one ASNInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (aiq *ASNInfoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aiq.Limit(2).IDs(setContextOp(ctx, aiq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asninfo.Label}
	default:
		err = &NotSingularError{asninfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aiq *ASNInfoQuery) OnlyIDX(ctx context.Context) int {
	id, err := aiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ASNInfos.
func (aiq *ASNInfoQuery) All(ctx context.Context) ([]*ASNInfo, error) {
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryAll)
	if err := aiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ASNInfo, *ASNInfoQuery]()
	return withInterceptors[[]*ASNInfo](ctx, aiq, qr, aiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aiq *ASNInfoQuery) AllX(ctx context.Context) []*ASNInfo {
	nodes, err := aiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ASNInfo IDs.
func (aiq *ASNInfoQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aiq.ctx.Unique == nil && aiq.path != nil {
		aiq.Unique(true)
	}
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryIDs)
	if err = aiq.Select(asninfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aiq *ASNInfoQuery) IDsX(ctx context.Context) []int {
	ids, err := aiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aiq *ASNInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryCount)
	if err := aiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aiq, querierCount[*ASNInfoQuery](), aiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aiq *ASNInfoQuery) CountX(ctx context.Context) int {
	count, err := aiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aiq *ASNInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryExist)
	switch _, err := aiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aiq *ASNInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := aiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ASNInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aiq *ASNInfoQuery) Clone() *ASNInfoQuery {
	if aiq == nil {
		return nil
	}
	return &ASNInfoQuery{
		config:        aiq.config,
		ctx:           aiq.ctx.Clone(),
		order:         append([]asninfo.OrderOption{}, aiq.order...),
		inters:        append([]Interceptor{}, aiq.inters...),
		predicates:    append([]predicate.ASNInfo{}, aiq.predicates...),
		withScan:      aiq.withScan.Clone(),
		withIpaddress: aiq.withIpaddress.Clone(),
		withRegistrar: aiq.withRegistrar.Clone(),
		withWhois:     aiq.withWhois.Clone(),
		// clone intermediate query.
		sql:  aiq.sql.Clone(),
		path: aiq.path,
	}
}

// WithScan tells the query-builder to eager-load the nodes that are connected to
// the "scan" edge. The optional arguments are used to configure the query builder of the edge.
func (aiq *ASNInfoQuery) WithScan(opts ...func(*ScanJobQuery)) *ASNInfoQuery {
	query := (&ScanJobClient{config: aiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aiq.withScan = query
	return aiq
}

// WithIpaddress tells the query-builder to eager-load the nodes that are connected to
// the "ipaddress" edge. The optional arguments are used to configure the query builder of the edge.
func (aiq *ASNInfoQuery) WithIpaddress(opts ...func(*IPAddressQuery)) *ASNInfoQuery {
	query := (&IPAddressClient{config: aiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aiq.withIpaddress = query
	return aiq
}

// WithRegistrar tells the query-builder to eager-load the nodes that are connected to
// the "registrar" edge. The optional arguments are used to configure the query builder of the edge.
func (aiq *ASNInfoQuery) WithRegistrar(opts ...func(*RegistrarQuery)) *ASNInfoQuery {
	query := (&RegistrarClient{config: aiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aiq.withRegistrar = query
	return aiq
}

// WithWhois tells the query-builder to eager-load the nodes that are connected to
// the "whois" edge. The optional arguments are used to configure the query builder of the edge.
func (aiq *ASNInfoQuery) WithWhois(opts ...func(*WhoisQuery)) *ASNInfoQuery {
	query := (&WhoisClient{config: aiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aiq.withWhois = query
	return aiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Asn int `json:"asn,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ASNInfo.Query().
//		GroupBy(asninfo.FieldAsn).
//		Aggregate(model_ent.Count()).
//		Scan(ctx, &v)
func (aiq *ASNInfoQuery) GroupBy(field string, fields ...string) *ASNInfoGroupBy {
	aiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ASNInfoGroupBy{build: aiq}
	grbuild.flds = &aiq.ctx.Fields
	grbuild.label = asninfo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Asn int `json:"asn,omitempty"`
//	}
//
//	client.ASNInfo.Query().
//		Select(asninfo.FieldAsn).
//		Scan(ctx, &v)
func (aiq *ASNInfoQuery) Select(fields ...string) *ASNInfoSelect {
	aiq.ctx.Fields = append(aiq.ctx.Fields, fields...)
	sbuild := &ASNInfoSelect{ASNInfoQuery: aiq}
	sbuild.label = asninfo.Label
	sbuild.flds, sbuild.scan = &aiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ASNInfoSelect configured with the given aggregations.
func (aiq *ASNInfoQuery) Aggregate(fns ...AggregateFunc) *ASNInfoSelect {
	return aiq.Select().Aggregate(fns...)
}

func (aiq *ASNInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aiq.inters {
		if inter == nil {
			return fmt.Errorf("model_ent: uninitialized interceptor (forgotten import model_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aiq); err != nil {
				return err
			}
		}
	}
	for _, f := range aiq.ctx.Fields {
		if !asninfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model_ent: invalid field %q for query", f)}
		}
	}
	if aiq.path != nil {
		prev, err := aiq.path(ctx)
		if err != nil {
			return err
		}
		aiq.sql = prev
	}
	return nil
}

func (aiq *ASNInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ASNInfo, error) {
	var (
		nodes       = []*ASNInfo{}
		withFKs     = aiq.withFKs
		_spec       = aiq.querySpec()
		loadedTypes = [4]bool{
			aiq.withScan != nil,
			aiq.withIpaddress != nil,
			aiq.withRegistrar != nil,
			aiq.withWhois != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, asninfo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ASNInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ASNInfo{config: aiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aiq.withScan; query != nil {
		if err := aiq.loadScan(ctx, query, nodes,
			func(n *ASNInfo) { n.Edges.Scan = []*ScanJob{} },
			func(n *ASNInfo, e *ScanJob) { n.Edges.Scan = append(n.Edges.Scan, e) }); err != nil {
			return nil, err
		}
	}
	if query := aiq.withIpaddress; query != nil {
		if err := aiq.loadIpaddress(ctx, query, nodes,
			func(n *ASNInfo) { n.Edges.Ipaddress = []*IPAddress{} },
			func(n *ASNInfo, e *IPAddress) { n.Edges.Ipaddress = append(n.Edges.Ipaddress, e) }); err != nil {
			return nil, err
		}
	}
	if query := aiq.withRegistrar; query != nil {
		if err := aiq.loadRegistrar(ctx, query, nodes,
			func(n *ASNInfo) { n.Edges.Registrar = []*Registrar{} },
			func(n *ASNInfo, e *Registrar) { n.Edges.Registrar = append(n.Edges.Registrar, e) }); err != nil {
			return nil, err
		}
	}
	if query := aiq.withWhois; query != nil {
		if err := aiq.loadWhois(ctx, query, nodes,
			func(n *ASNInfo) { n.Edges.Whois = []*Whois{} },
			func(n *ASNInfo, e *Whois) { n.Edges.Whois = append(n.Edges.Whois, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aiq *ASNInfoQuery) loadScan(ctx context.Context, query *ScanJobQuery, nodes []*ASNInfo, init func(*ASNInfo), assign func(*ASNInfo, *ScanJob)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ASNInfo)
	nids := make(map[int]map[*ASNInfo]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(asninfo.ScanTable)
		s.Join(joinT).On(s.C(scanjob.FieldID), joinT.C(asninfo.ScanPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(asninfo.ScanPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(asninfo.ScanPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ASNInfo]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*ScanJob](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "scan" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aiq *ASNInfoQuery) loadIpaddress(ctx context.Context, query *IPAddressQuery, nodes []*ASNInfo, init func(*ASNInfo), assign func(*ASNInfo, *IPAddress)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ASNInfo)
	nids := make(map[int]map[*ASNInfo]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(asninfo.IpaddressTable)
		s.Join(joinT).On(s.C(ipaddress.FieldID), joinT.C(asninfo.IpaddressPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(asninfo.IpaddressPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(asninfo.IpaddressPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ASNInfo]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*IPAddress](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "ipaddress" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aiq *ASNInfoQuery) loadRegistrar(ctx context.Context, query *RegistrarQuery, nodes []*ASNInfo, init func(*ASNInfo), assign func(*ASNInfo, *Registrar)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ASNInfo)
	nids := make(map[int]map[*ASNInfo]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(asninfo.RegistrarTable)
		s.Join(joinT).On(s.C(registrar.FieldID), joinT.C(asninfo.RegistrarPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(asninfo.RegistrarPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(asninfo.RegistrarPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ASNInfo]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Registrar](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "registrar" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aiq *ASNInfoQuery) loadWhois(ctx context.Context, query *WhoisQuery, nodes []*ASNInfo, init func(*ASNInfo), assign func(*ASNInfo, *Whois)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ASNInfo)
	nids := make(map[int]map[*ASNInfo]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(asninfo.WhoisTable)
		s.Join(joinT).On(s.C(whois.FieldID), joinT.C(asninfo.WhoisPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(asninfo.WhoisPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(asninfo.WhoisPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ASNInfo]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Whois](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "whois" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (aiq *ASNInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aiq.querySpec()
	_spec.Node.Columns = aiq.ctx.Fields
	if len(aiq.ctx.Fields) > 0 {
		_spec.Unique = aiq.ctx.Unique != nil && *aiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aiq.driver, _spec)
}

func (aiq *ASNInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(asninfo.Table, asninfo.Columns, sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt))
	_spec.From = aiq.sql
	if unique := aiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aiq.path != nil {
		_spec.Unique = true
	}
	if fields := aiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asninfo.FieldID)
		for i := range fields {
			if fields[i] != asninfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aiq *ASNInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aiq.driver.Dialect())
	t1 := builder.Table(asninfo.Table)
	columns := aiq.ctx.Fields
	if len(columns) == 0 {
		columns = asninfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aiq.sql != nil {
		selector = aiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aiq.ctx.Unique != nil && *aiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aiq.predicates {
		p(selector)
	}
	for _, p := range aiq.order {
		p(selector)
	}
	if offset := aiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ASNInfoGroupBy is the group-by builder for ASNInfo entities.
type ASNInfoGroupBy struct {
	selector
	build *ASNInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aigb *ASNInfoGroupBy) Aggregate(fns ...AggregateFunc) *ASNInfoGroupBy {
	aigb.fns = append(aigb.fns, fns...)
	return aigb
}

// Scan applies the selector query and scans the result into the given value.
func (aigb *ASNInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aigb.build.ctx, ent.OpQueryGroupBy)
	if err := aigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ASNInfoQuery, *ASNInfoGroupBy](ctx, aigb.build, aigb, aigb.build.inters, v)
}

func (aigb *ASNInfoGroupBy) sqlScan(ctx context.Context, root *ASNInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aigb.fns))
	for _, fn := range aigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aigb.flds)+len(aigb.fns))
		for _, f := range *aigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ASNInfoSelect is the builder for selecting fields of ASNInfo entities.
type ASNInfoSelect struct {
	*ASNInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ais *ASNInfoSelect) Aggregate(fns ...AggregateFunc) *ASNInfoSelect {
	ais.fns = append(ais.fns, fns...)
	return ais
}

// Scan applies the selector query and scans the result into the given value.
func (ais *ASNInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ais.ctx, ent.OpQuerySelect)
	if err := ais.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ASNInfoQuery, *ASNInfoSelect](ctx, ais.ASNInfoQuery, ais, ais.inters, v)
}

func (ais *ASNInfoSelect) sqlScan(ctx context.Context, root *ASNInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ais.fns))
	for _, fn := range ais.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ais.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ais.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
