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
	"github.com/ice-bergtech/dnh/src/internal/model_ent/domain"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/nameserver"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/registrar"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scan"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/whois"
)

// WhoisQuery is the builder for querying Whois entities.
type WhoisQuery struct {
	config
	ctx            *QueryContext
	order          []whois.OrderOption
	inters         []Interceptor
	predicates     []predicate.Whois
	withIprange    *IPAddressQuery
	withDomain     *DomainQuery
	withAsn        *ASNInfoQuery
	withRegistrar  *RegistrarQuery
	withNameserver *NameserverQuery
	withScan       *ScanQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WhoisQuery builder.
func (wq *WhoisQuery) Where(ps ...predicate.Whois) *WhoisQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WhoisQuery) Limit(limit int) *WhoisQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WhoisQuery) Offset(offset int) *WhoisQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WhoisQuery) Unique(unique bool) *WhoisQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WhoisQuery) Order(o ...whois.OrderOption) *WhoisQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryIprange chains the current query on the "iprange" edge.
func (wq *WhoisQuery) QueryIprange() *IPAddressQuery {
	query := (&IPAddressClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(whois.Table, whois.FieldID, selector),
			sqlgraph.To(ipaddress.Table, ipaddress.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, whois.IprangeTable, whois.IprangePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDomain chains the current query on the "domain" edge.
func (wq *WhoisQuery) QueryDomain() *DomainQuery {
	query := (&DomainClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(whois.Table, whois.FieldID, selector),
			sqlgraph.To(domain.Table, domain.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, whois.DomainTable, whois.DomainPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAsn chains the current query on the "asn" edge.
func (wq *WhoisQuery) QueryAsn() *ASNInfoQuery {
	query := (&ASNInfoClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(whois.Table, whois.FieldID, selector),
			sqlgraph.To(asninfo.Table, asninfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, whois.AsnTable, whois.AsnPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRegistrar chains the current query on the "registrar" edge.
func (wq *WhoisQuery) QueryRegistrar() *RegistrarQuery {
	query := (&RegistrarClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(whois.Table, whois.FieldID, selector),
			sqlgraph.To(registrar.Table, registrar.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, whois.RegistrarTable, whois.RegistrarPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNameserver chains the current query on the "nameserver" edge.
func (wq *WhoisQuery) QueryNameserver() *NameserverQuery {
	query := (&NameserverClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(whois.Table, whois.FieldID, selector),
			sqlgraph.To(nameserver.Table, nameserver.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, whois.NameserverTable, whois.NameserverPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScan chains the current query on the "scan" edge.
func (wq *WhoisQuery) QueryScan() *ScanQuery {
	query := (&ScanClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(whois.Table, whois.FieldID, selector),
			sqlgraph.To(scan.Table, scan.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, whois.ScanTable, whois.ScanPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Whois entity from the query.
// Returns a *NotFoundError when no Whois was found.
func (wq *WhoisQuery) First(ctx context.Context) (*Whois, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{whois.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WhoisQuery) FirstX(ctx context.Context) *Whois {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Whois ID from the query.
// Returns a *NotFoundError when no Whois ID was found.
func (wq *WhoisQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{whois.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WhoisQuery) FirstIDX(ctx context.Context) int {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Whois entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Whois entity is found.
// Returns a *NotFoundError when no Whois entities are found.
func (wq *WhoisQuery) Only(ctx context.Context) (*Whois, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{whois.Label}
	default:
		return nil, &NotSingularError{whois.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WhoisQuery) OnlyX(ctx context.Context) *Whois {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Whois ID in the query.
// Returns a *NotSingularError when more than one Whois ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WhoisQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{whois.Label}
	default:
		err = &NotSingularError{whois.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WhoisQuery) OnlyIDX(ctx context.Context) int {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WhoisSlice.
func (wq *WhoisQuery) All(ctx context.Context) ([]*Whois, error) {
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryAll)
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Whois, *WhoisQuery]()
	return withInterceptors[[]*Whois](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WhoisQuery) AllX(ctx context.Context) []*Whois {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Whois IDs.
func (wq *WhoisQuery) IDs(ctx context.Context) (ids []int, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryIDs)
	if err = wq.Select(whois.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WhoisQuery) IDsX(ctx context.Context) []int {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WhoisQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryCount)
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WhoisQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WhoisQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WhoisQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryExist)
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WhoisQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WhoisQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WhoisQuery) Clone() *WhoisQuery {
	if wq == nil {
		return nil
	}
	return &WhoisQuery{
		config:         wq.config,
		ctx:            wq.ctx.Clone(),
		order:          append([]whois.OrderOption{}, wq.order...),
		inters:         append([]Interceptor{}, wq.inters...),
		predicates:     append([]predicate.Whois{}, wq.predicates...),
		withIprange:    wq.withIprange.Clone(),
		withDomain:     wq.withDomain.Clone(),
		withAsn:        wq.withAsn.Clone(),
		withRegistrar:  wq.withRegistrar.Clone(),
		withNameserver: wq.withNameserver.Clone(),
		withScan:       wq.withScan.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithIprange tells the query-builder to eager-load the nodes that are connected to
// the "iprange" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WhoisQuery) WithIprange(opts ...func(*IPAddressQuery)) *WhoisQuery {
	query := (&IPAddressClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withIprange = query
	return wq
}

// WithDomain tells the query-builder to eager-load the nodes that are connected to
// the "domain" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WhoisQuery) WithDomain(opts ...func(*DomainQuery)) *WhoisQuery {
	query := (&DomainClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withDomain = query
	return wq
}

// WithAsn tells the query-builder to eager-load the nodes that are connected to
// the "asn" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WhoisQuery) WithAsn(opts ...func(*ASNInfoQuery)) *WhoisQuery {
	query := (&ASNInfoClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withAsn = query
	return wq
}

// WithRegistrar tells the query-builder to eager-load the nodes that are connected to
// the "registrar" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WhoisQuery) WithRegistrar(opts ...func(*RegistrarQuery)) *WhoisQuery {
	query := (&RegistrarClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withRegistrar = query
	return wq
}

// WithNameserver tells the query-builder to eager-load the nodes that are connected to
// the "nameserver" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WhoisQuery) WithNameserver(opts ...func(*NameserverQuery)) *WhoisQuery {
	query := (&NameserverClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withNameserver = query
	return wq
}

// WithScan tells the query-builder to eager-load the nodes that are connected to
// the "scan" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WhoisQuery) WithScan(opts ...func(*ScanQuery)) *WhoisQuery {
	query := (&ScanClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withScan = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Query string `json:"query,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Whois.Query().
//		GroupBy(whois.FieldQuery).
//		Aggregate(model_ent.Count()).
//		Scan(ctx, &v)
func (wq *WhoisQuery) GroupBy(field string, fields ...string) *WhoisGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WhoisGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = whois.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Query string `json:"query,omitempty"`
//	}
//
//	client.Whois.Query().
//		Select(whois.FieldQuery).
//		Scan(ctx, &v)
func (wq *WhoisQuery) Select(fields ...string) *WhoisSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WhoisSelect{WhoisQuery: wq}
	sbuild.label = whois.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WhoisSelect configured with the given aggregations.
func (wq *WhoisQuery) Aggregate(fns ...AggregateFunc) *WhoisSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WhoisQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("model_ent: uninitialized interceptor (forgotten import model_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !whois.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model_ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WhoisQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Whois, error) {
	var (
		nodes       = []*Whois{}
		withFKs     = wq.withFKs
		_spec       = wq.querySpec()
		loadedTypes = [6]bool{
			wq.withIprange != nil,
			wq.withDomain != nil,
			wq.withAsn != nil,
			wq.withRegistrar != nil,
			wq.withNameserver != nil,
			wq.withScan != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, whois.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Whois).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Whois{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withIprange; query != nil {
		if err := wq.loadIprange(ctx, query, nodes,
			func(n *Whois) { n.Edges.Iprange = []*IPAddress{} },
			func(n *Whois, e *IPAddress) { n.Edges.Iprange = append(n.Edges.Iprange, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withDomain; query != nil {
		if err := wq.loadDomain(ctx, query, nodes,
			func(n *Whois) { n.Edges.Domain = []*Domain{} },
			func(n *Whois, e *Domain) { n.Edges.Domain = append(n.Edges.Domain, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withAsn; query != nil {
		if err := wq.loadAsn(ctx, query, nodes,
			func(n *Whois) { n.Edges.Asn = []*ASNInfo{} },
			func(n *Whois, e *ASNInfo) { n.Edges.Asn = append(n.Edges.Asn, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withRegistrar; query != nil {
		if err := wq.loadRegistrar(ctx, query, nodes,
			func(n *Whois) { n.Edges.Registrar = []*Registrar{} },
			func(n *Whois, e *Registrar) { n.Edges.Registrar = append(n.Edges.Registrar, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withNameserver; query != nil {
		if err := wq.loadNameserver(ctx, query, nodes,
			func(n *Whois) { n.Edges.Nameserver = []*Nameserver{} },
			func(n *Whois, e *Nameserver) { n.Edges.Nameserver = append(n.Edges.Nameserver, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withScan; query != nil {
		if err := wq.loadScan(ctx, query, nodes,
			func(n *Whois) { n.Edges.Scan = []*Scan{} },
			func(n *Whois, e *Scan) { n.Edges.Scan = append(n.Edges.Scan, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WhoisQuery) loadIprange(ctx context.Context, query *IPAddressQuery, nodes []*Whois, init func(*Whois), assign func(*Whois, *IPAddress)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Whois)
	nids := make(map[int]map[*Whois]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(whois.IprangeTable)
		s.Join(joinT).On(s.C(ipaddress.FieldID), joinT.C(whois.IprangePrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(whois.IprangePrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(whois.IprangePrimaryKey[0]))
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
					nids[inValue] = map[*Whois]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "iprange" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (wq *WhoisQuery) loadDomain(ctx context.Context, query *DomainQuery, nodes []*Whois, init func(*Whois), assign func(*Whois, *Domain)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Whois)
	nids := make(map[int]map[*Whois]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(whois.DomainTable)
		s.Join(joinT).On(s.C(domain.FieldID), joinT.C(whois.DomainPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(whois.DomainPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(whois.DomainPrimaryKey[0]))
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
					nids[inValue] = map[*Whois]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Domain](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "domain" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (wq *WhoisQuery) loadAsn(ctx context.Context, query *ASNInfoQuery, nodes []*Whois, init func(*Whois), assign func(*Whois, *ASNInfo)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Whois)
	nids := make(map[int]map[*Whois]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(whois.AsnTable)
		s.Join(joinT).On(s.C(asninfo.FieldID), joinT.C(whois.AsnPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(whois.AsnPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(whois.AsnPrimaryKey[0]))
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
					nids[inValue] = map[*Whois]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*ASNInfo](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "asn" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (wq *WhoisQuery) loadRegistrar(ctx context.Context, query *RegistrarQuery, nodes []*Whois, init func(*Whois), assign func(*Whois, *Registrar)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Whois)
	nids := make(map[int]map[*Whois]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(whois.RegistrarTable)
		s.Join(joinT).On(s.C(registrar.FieldID), joinT.C(whois.RegistrarPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(whois.RegistrarPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(whois.RegistrarPrimaryKey[0]))
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
					nids[inValue] = map[*Whois]struct{}{byID[outValue]: {}}
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
func (wq *WhoisQuery) loadNameserver(ctx context.Context, query *NameserverQuery, nodes []*Whois, init func(*Whois), assign func(*Whois, *Nameserver)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Whois)
	nids := make(map[int]map[*Whois]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(whois.NameserverTable)
		s.Join(joinT).On(s.C(nameserver.FieldID), joinT.C(whois.NameserverPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(whois.NameserverPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(whois.NameserverPrimaryKey[0]))
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
					nids[inValue] = map[*Whois]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Nameserver](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "nameserver" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (wq *WhoisQuery) loadScan(ctx context.Context, query *ScanQuery, nodes []*Whois, init func(*Whois), assign func(*Whois, *Scan)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Whois)
	nids := make(map[int]map[*Whois]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(whois.ScanTable)
		s.Join(joinT).On(s.C(scan.FieldID), joinT.C(whois.ScanPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(whois.ScanPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(whois.ScanPrimaryKey[1]))
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
					nids[inValue] = map[*Whois]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Scan](ctx, query, qr, query.inters)
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

func (wq *WhoisQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WhoisQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(whois.Table, whois.Columns, sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, whois.FieldID)
		for i := range fields {
			if fields[i] != whois.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WhoisQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(whois.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = whois.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WhoisGroupBy is the group-by builder for Whois entities.
type WhoisGroupBy struct {
	selector
	build *WhoisQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WhoisGroupBy) Aggregate(fns ...AggregateFunc) *WhoisGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WhoisGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, ent.OpQueryGroupBy)
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WhoisQuery, *WhoisGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WhoisGroupBy) sqlScan(ctx context.Context, root *WhoisQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WhoisSelect is the builder for selecting fields of Whois entities.
type WhoisSelect struct {
	*WhoisQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WhoisSelect) Aggregate(fns ...AggregateFunc) *WhoisSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WhoisSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, ent.OpQuerySelect)
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WhoisQuery, *WhoisSelect](ctx, ws.WhoisQuery, ws, ws.inters, v)
}

func (ws *WhoisSelect) sqlScan(ctx context.Context, root *WhoisQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
