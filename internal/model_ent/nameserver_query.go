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
	"github.com/ice-bergtech/dnh/src/internal/model_ent/dnsentry"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/domain"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/nameserver"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scanjob"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/whois"
)

// NameserverQuery is the builder for querying Nameserver entities.
type NameserverQuery struct {
	config
	ctx           *QueryContext
	order         []nameserver.OrderOption
	inters        []Interceptor
	predicates    []predicate.Nameserver
	withIpaddress *IPAddressQuery
	withScan      *ScanJobQuery
	withDnsentry  *DNSEntryQuery
	withDomain    *DomainQuery
	withWhois     *WhoisQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NameserverQuery builder.
func (nq *NameserverQuery) Where(ps ...predicate.Nameserver) *NameserverQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NameserverQuery) Limit(limit int) *NameserverQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NameserverQuery) Offset(offset int) *NameserverQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NameserverQuery) Unique(unique bool) *NameserverQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NameserverQuery) Order(o ...nameserver.OrderOption) *NameserverQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryIpaddress chains the current query on the "ipaddress" edge.
func (nq *NameserverQuery) QueryIpaddress() *IPAddressQuery {
	query := (&IPAddressClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nameserver.Table, nameserver.FieldID, selector),
			sqlgraph.To(ipaddress.Table, ipaddress.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, nameserver.IpaddressTable, nameserver.IpaddressPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScan chains the current query on the "scan" edge.
func (nq *NameserverQuery) QueryScan() *ScanJobQuery {
	query := (&ScanJobClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nameserver.Table, nameserver.FieldID, selector),
			sqlgraph.To(scanjob.Table, scanjob.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, nameserver.ScanTable, nameserver.ScanPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDnsentry chains the current query on the "dnsentry" edge.
func (nq *NameserverQuery) QueryDnsentry() *DNSEntryQuery {
	query := (&DNSEntryClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nameserver.Table, nameserver.FieldID, selector),
			sqlgraph.To(dnsentry.Table, dnsentry.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, nameserver.DnsentryTable, nameserver.DnsentryPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDomain chains the current query on the "domain" edge.
func (nq *NameserverQuery) QueryDomain() *DomainQuery {
	query := (&DomainClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nameserver.Table, nameserver.FieldID, selector),
			sqlgraph.To(domain.Table, domain.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, nameserver.DomainTable, nameserver.DomainPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWhois chains the current query on the "whois" edge.
func (nq *NameserverQuery) QueryWhois() *WhoisQuery {
	query := (&WhoisClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nameserver.Table, nameserver.FieldID, selector),
			sqlgraph.To(whois.Table, whois.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, nameserver.WhoisTable, nameserver.WhoisPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Nameserver entity from the query.
// Returns a *NotFoundError when no Nameserver was found.
func (nq *NameserverQuery) First(ctx context.Context) (*Nameserver, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nameserver.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NameserverQuery) FirstX(ctx context.Context) *Nameserver {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Nameserver ID from the query.
// Returns a *NotFoundError when no Nameserver ID was found.
func (nq *NameserverQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nameserver.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NameserverQuery) FirstIDX(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Nameserver entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Nameserver entity is found.
// Returns a *NotFoundError when no Nameserver entities are found.
func (nq *NameserverQuery) Only(ctx context.Context) (*Nameserver, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nameserver.Label}
	default:
		return nil, &NotSingularError{nameserver.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NameserverQuery) OnlyX(ctx context.Context) *Nameserver {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Nameserver ID in the query.
// Returns a *NotSingularError when more than one Nameserver ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NameserverQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nameserver.Label}
	default:
		err = &NotSingularError{nameserver.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NameserverQuery) OnlyIDX(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Nameservers.
func (nq *NameserverQuery) All(ctx context.Context) ([]*Nameserver, error) {
	ctx = setContextOp(ctx, nq.ctx, ent.OpQueryAll)
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Nameserver, *NameserverQuery]()
	return withInterceptors[[]*Nameserver](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NameserverQuery) AllX(ctx context.Context) []*Nameserver {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Nameserver IDs.
func (nq *NameserverQuery) IDs(ctx context.Context) (ids []int, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, ent.OpQueryIDs)
	if err = nq.Select(nameserver.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NameserverQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NameserverQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, ent.OpQueryCount)
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NameserverQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NameserverQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NameserverQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, ent.OpQueryExist)
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NameserverQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NameserverQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NameserverQuery) Clone() *NameserverQuery {
	if nq == nil {
		return nil
	}
	return &NameserverQuery{
		config:        nq.config,
		ctx:           nq.ctx.Clone(),
		order:         append([]nameserver.OrderOption{}, nq.order...),
		inters:        append([]Interceptor{}, nq.inters...),
		predicates:    append([]predicate.Nameserver{}, nq.predicates...),
		withIpaddress: nq.withIpaddress.Clone(),
		withScan:      nq.withScan.Clone(),
		withDnsentry:  nq.withDnsentry.Clone(),
		withDomain:    nq.withDomain.Clone(),
		withWhois:     nq.withWhois.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithIpaddress tells the query-builder to eager-load the nodes that are connected to
// the "ipaddress" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NameserverQuery) WithIpaddress(opts ...func(*IPAddressQuery)) *NameserverQuery {
	query := (&IPAddressClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withIpaddress = query
	return nq
}

// WithScan tells the query-builder to eager-load the nodes that are connected to
// the "scan" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NameserverQuery) WithScan(opts ...func(*ScanJobQuery)) *NameserverQuery {
	query := (&ScanJobClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withScan = query
	return nq
}

// WithDnsentry tells the query-builder to eager-load the nodes that are connected to
// the "dnsentry" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NameserverQuery) WithDnsentry(opts ...func(*DNSEntryQuery)) *NameserverQuery {
	query := (&DNSEntryClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withDnsentry = query
	return nq
}

// WithDomain tells the query-builder to eager-load the nodes that are connected to
// the "domain" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NameserverQuery) WithDomain(opts ...func(*DomainQuery)) *NameserverQuery {
	query := (&DomainClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withDomain = query
	return nq
}

// WithWhois tells the query-builder to eager-load the nodes that are connected to
// the "whois" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NameserverQuery) WithWhois(opts ...func(*WhoisQuery)) *NameserverQuery {
	query := (&WhoisClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withWhois = query
	return nq
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
//	client.Nameserver.Query().
//		GroupBy(nameserver.FieldName).
//		Aggregate(model_ent.Count()).
//		Scan(ctx, &v)
func (nq *NameserverQuery) GroupBy(field string, fields ...string) *NameserverGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NameserverGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = nameserver.Label
	grbuild.scan = grbuild.Scan
	return grbuild
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
//	client.Nameserver.Query().
//		Select(nameserver.FieldName).
//		Scan(ctx, &v)
func (nq *NameserverQuery) Select(fields ...string) *NameserverSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NameserverSelect{NameserverQuery: nq}
	sbuild.label = nameserver.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NameserverSelect configured with the given aggregations.
func (nq *NameserverQuery) Aggregate(fns ...AggregateFunc) *NameserverSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NameserverQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("model_ent: uninitialized interceptor (forgotten import model_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !nameserver.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model_ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NameserverQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Nameserver, error) {
	var (
		nodes       = []*Nameserver{}
		withFKs     = nq.withFKs
		_spec       = nq.querySpec()
		loadedTypes = [5]bool{
			nq.withIpaddress != nil,
			nq.withScan != nil,
			nq.withDnsentry != nil,
			nq.withDomain != nil,
			nq.withWhois != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, nameserver.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Nameserver).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Nameserver{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withIpaddress; query != nil {
		if err := nq.loadIpaddress(ctx, query, nodes,
			func(n *Nameserver) { n.Edges.Ipaddress = []*IPAddress{} },
			func(n *Nameserver, e *IPAddress) { n.Edges.Ipaddress = append(n.Edges.Ipaddress, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withScan; query != nil {
		if err := nq.loadScan(ctx, query, nodes,
			func(n *Nameserver) { n.Edges.Scan = []*ScanJob{} },
			func(n *Nameserver, e *ScanJob) { n.Edges.Scan = append(n.Edges.Scan, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withDnsentry; query != nil {
		if err := nq.loadDnsentry(ctx, query, nodes,
			func(n *Nameserver) { n.Edges.Dnsentry = []*DNSEntry{} },
			func(n *Nameserver, e *DNSEntry) { n.Edges.Dnsentry = append(n.Edges.Dnsentry, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withDomain; query != nil {
		if err := nq.loadDomain(ctx, query, nodes,
			func(n *Nameserver) { n.Edges.Domain = []*Domain{} },
			func(n *Nameserver, e *Domain) { n.Edges.Domain = append(n.Edges.Domain, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withWhois; query != nil {
		if err := nq.loadWhois(ctx, query, nodes,
			func(n *Nameserver) { n.Edges.Whois = []*Whois{} },
			func(n *Nameserver, e *Whois) { n.Edges.Whois = append(n.Edges.Whois, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NameserverQuery) loadIpaddress(ctx context.Context, query *IPAddressQuery, nodes []*Nameserver, init func(*Nameserver), assign func(*Nameserver, *IPAddress)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Nameserver)
	nids := make(map[int]map[*Nameserver]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(nameserver.IpaddressTable)
		s.Join(joinT).On(s.C(ipaddress.FieldID), joinT.C(nameserver.IpaddressPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(nameserver.IpaddressPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(nameserver.IpaddressPrimaryKey[0]))
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
					nids[inValue] = map[*Nameserver]struct{}{byID[outValue]: {}}
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
func (nq *NameserverQuery) loadScan(ctx context.Context, query *ScanJobQuery, nodes []*Nameserver, init func(*Nameserver), assign func(*Nameserver, *ScanJob)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Nameserver)
	nids := make(map[int]map[*Nameserver]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(nameserver.ScanTable)
		s.Join(joinT).On(s.C(scanjob.FieldID), joinT.C(nameserver.ScanPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(nameserver.ScanPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(nameserver.ScanPrimaryKey[1]))
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
					nids[inValue] = map[*Nameserver]struct{}{byID[outValue]: {}}
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
func (nq *NameserverQuery) loadDnsentry(ctx context.Context, query *DNSEntryQuery, nodes []*Nameserver, init func(*Nameserver), assign func(*Nameserver, *DNSEntry)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Nameserver)
	nids := make(map[int]map[*Nameserver]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(nameserver.DnsentryTable)
		s.Join(joinT).On(s.C(dnsentry.FieldID), joinT.C(nameserver.DnsentryPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(nameserver.DnsentryPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(nameserver.DnsentryPrimaryKey[1]))
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
					nids[inValue] = map[*Nameserver]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*DNSEntry](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "dnsentry" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (nq *NameserverQuery) loadDomain(ctx context.Context, query *DomainQuery, nodes []*Nameserver, init func(*Nameserver), assign func(*Nameserver, *Domain)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Nameserver)
	nids := make(map[int]map[*Nameserver]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(nameserver.DomainTable)
		s.Join(joinT).On(s.C(domain.FieldID), joinT.C(nameserver.DomainPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(nameserver.DomainPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(nameserver.DomainPrimaryKey[1]))
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
					nids[inValue] = map[*Nameserver]struct{}{byID[outValue]: {}}
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
func (nq *NameserverQuery) loadWhois(ctx context.Context, query *WhoisQuery, nodes []*Nameserver, init func(*Nameserver), assign func(*Nameserver, *Whois)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Nameserver)
	nids := make(map[int]map[*Nameserver]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(nameserver.WhoisTable)
		s.Join(joinT).On(s.C(whois.FieldID), joinT.C(nameserver.WhoisPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(nameserver.WhoisPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(nameserver.WhoisPrimaryKey[1]))
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
					nids[inValue] = map[*Nameserver]struct{}{byID[outValue]: {}}
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

func (nq *NameserverQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NameserverQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(nameserver.Table, nameserver.Columns, sqlgraph.NewFieldSpec(nameserver.FieldID, field.TypeInt))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nameserver.FieldID)
		for i := range fields {
			if fields[i] != nameserver.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NameserverQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(nameserver.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = nameserver.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NameserverGroupBy is the group-by builder for Nameserver entities.
type NameserverGroupBy struct {
	selector
	build *NameserverQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NameserverGroupBy) Aggregate(fns ...AggregateFunc) *NameserverGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NameserverGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, ent.OpQueryGroupBy)
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NameserverQuery, *NameserverGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NameserverGroupBy) sqlScan(ctx context.Context, root *NameserverQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NameserverSelect is the builder for selecting fields of Nameserver entities.
type NameserverSelect struct {
	*NameserverQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NameserverSelect) Aggregate(fns ...AggregateFunc) *NameserverSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NameserverSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, ent.OpQuerySelect)
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NameserverQuery, *NameserverSelect](ctx, ns.NameserverQuery, ns, ns.inters, v)
}

func (ns *NameserverSelect) sqlScan(ctx context.Context, root *NameserverQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
