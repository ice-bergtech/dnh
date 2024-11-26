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
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scan"
)

// DNSEntryQuery is the builder for querying DNSEntry entities.
type DNSEntryQuery struct {
	config
	ctx            *QueryContext
	order          []dnsentry.OrderOption
	inters         []Interceptor
	predicates     []predicate.DNSEntry
	withDomain     *DomainQuery
	withIpaddress  *IPAddressQuery
	withNameserver *NameserverQuery
	withScan       *ScanQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DNSEntryQuery builder.
func (deq *DNSEntryQuery) Where(ps ...predicate.DNSEntry) *DNSEntryQuery {
	deq.predicates = append(deq.predicates, ps...)
	return deq
}

// Limit the number of records to be returned by this query.
func (deq *DNSEntryQuery) Limit(limit int) *DNSEntryQuery {
	deq.ctx.Limit = &limit
	return deq
}

// Offset to start from.
func (deq *DNSEntryQuery) Offset(offset int) *DNSEntryQuery {
	deq.ctx.Offset = &offset
	return deq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (deq *DNSEntryQuery) Unique(unique bool) *DNSEntryQuery {
	deq.ctx.Unique = &unique
	return deq
}

// Order specifies how the records should be ordered.
func (deq *DNSEntryQuery) Order(o ...dnsentry.OrderOption) *DNSEntryQuery {
	deq.order = append(deq.order, o...)
	return deq
}

// QueryDomain chains the current query on the "domain" edge.
func (deq *DNSEntryQuery) QueryDomain() *DomainQuery {
	query := (&DomainClient{config: deq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := deq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := deq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dnsentry.Table, dnsentry.FieldID, selector),
			sqlgraph.To(domain.Table, domain.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dnsentry.DomainTable, dnsentry.DomainColumn),
		)
		fromU = sqlgraph.SetNeighbors(deq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIpaddress chains the current query on the "ipaddress" edge.
func (deq *DNSEntryQuery) QueryIpaddress() *IPAddressQuery {
	query := (&IPAddressClient{config: deq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := deq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := deq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dnsentry.Table, dnsentry.FieldID, selector),
			sqlgraph.To(ipaddress.Table, ipaddress.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dnsentry.IpaddressTable, dnsentry.IpaddressColumn),
		)
		fromU = sqlgraph.SetNeighbors(deq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNameserver chains the current query on the "nameserver" edge.
func (deq *DNSEntryQuery) QueryNameserver() *NameserverQuery {
	query := (&NameserverClient{config: deq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := deq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := deq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dnsentry.Table, dnsentry.FieldID, selector),
			sqlgraph.To(nameserver.Table, nameserver.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dnsentry.NameserverTable, dnsentry.NameserverColumn),
		)
		fromU = sqlgraph.SetNeighbors(deq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScan chains the current query on the "scan" edge.
func (deq *DNSEntryQuery) QueryScan() *ScanQuery {
	query := (&ScanClient{config: deq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := deq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := deq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dnsentry.Table, dnsentry.FieldID, selector),
			sqlgraph.To(scan.Table, scan.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, dnsentry.ScanTable, dnsentry.ScanPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(deq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DNSEntry entity from the query.
// Returns a *NotFoundError when no DNSEntry was found.
func (deq *DNSEntryQuery) First(ctx context.Context) (*DNSEntry, error) {
	nodes, err := deq.Limit(1).All(setContextOp(ctx, deq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dnsentry.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (deq *DNSEntryQuery) FirstX(ctx context.Context) *DNSEntry {
	node, err := deq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DNSEntry ID from the query.
// Returns a *NotFoundError when no DNSEntry ID was found.
func (deq *DNSEntryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = deq.Limit(1).IDs(setContextOp(ctx, deq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dnsentry.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (deq *DNSEntryQuery) FirstIDX(ctx context.Context) int {
	id, err := deq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DNSEntry entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DNSEntry entity is found.
// Returns a *NotFoundError when no DNSEntry entities are found.
func (deq *DNSEntryQuery) Only(ctx context.Context) (*DNSEntry, error) {
	nodes, err := deq.Limit(2).All(setContextOp(ctx, deq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dnsentry.Label}
	default:
		return nil, &NotSingularError{dnsentry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (deq *DNSEntryQuery) OnlyX(ctx context.Context) *DNSEntry {
	node, err := deq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DNSEntry ID in the query.
// Returns a *NotSingularError when more than one DNSEntry ID is found.
// Returns a *NotFoundError when no entities are found.
func (deq *DNSEntryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = deq.Limit(2).IDs(setContextOp(ctx, deq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dnsentry.Label}
	default:
		err = &NotSingularError{dnsentry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (deq *DNSEntryQuery) OnlyIDX(ctx context.Context) int {
	id, err := deq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DNSEntries.
func (deq *DNSEntryQuery) All(ctx context.Context) ([]*DNSEntry, error) {
	ctx = setContextOp(ctx, deq.ctx, ent.OpQueryAll)
	if err := deq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DNSEntry, *DNSEntryQuery]()
	return withInterceptors[[]*DNSEntry](ctx, deq, qr, deq.inters)
}

// AllX is like All, but panics if an error occurs.
func (deq *DNSEntryQuery) AllX(ctx context.Context) []*DNSEntry {
	nodes, err := deq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DNSEntry IDs.
func (deq *DNSEntryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if deq.ctx.Unique == nil && deq.path != nil {
		deq.Unique(true)
	}
	ctx = setContextOp(ctx, deq.ctx, ent.OpQueryIDs)
	if err = deq.Select(dnsentry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (deq *DNSEntryQuery) IDsX(ctx context.Context) []int {
	ids, err := deq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (deq *DNSEntryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, deq.ctx, ent.OpQueryCount)
	if err := deq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, deq, querierCount[*DNSEntryQuery](), deq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (deq *DNSEntryQuery) CountX(ctx context.Context) int {
	count, err := deq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (deq *DNSEntryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, deq.ctx, ent.OpQueryExist)
	switch _, err := deq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (deq *DNSEntryQuery) ExistX(ctx context.Context) bool {
	exist, err := deq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DNSEntryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (deq *DNSEntryQuery) Clone() *DNSEntryQuery {
	if deq == nil {
		return nil
	}
	return &DNSEntryQuery{
		config:         deq.config,
		ctx:            deq.ctx.Clone(),
		order:          append([]dnsentry.OrderOption{}, deq.order...),
		inters:         append([]Interceptor{}, deq.inters...),
		predicates:     append([]predicate.DNSEntry{}, deq.predicates...),
		withDomain:     deq.withDomain.Clone(),
		withIpaddress:  deq.withIpaddress.Clone(),
		withNameserver: deq.withNameserver.Clone(),
		withScan:       deq.withScan.Clone(),
		// clone intermediate query.
		sql:  deq.sql.Clone(),
		path: deq.path,
	}
}

// WithDomain tells the query-builder to eager-load the nodes that are connected to
// the "domain" edge. The optional arguments are used to configure the query builder of the edge.
func (deq *DNSEntryQuery) WithDomain(opts ...func(*DomainQuery)) *DNSEntryQuery {
	query := (&DomainClient{config: deq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	deq.withDomain = query
	return deq
}

// WithIpaddress tells the query-builder to eager-load the nodes that are connected to
// the "ipaddress" edge. The optional arguments are used to configure the query builder of the edge.
func (deq *DNSEntryQuery) WithIpaddress(opts ...func(*IPAddressQuery)) *DNSEntryQuery {
	query := (&IPAddressClient{config: deq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	deq.withIpaddress = query
	return deq
}

// WithNameserver tells the query-builder to eager-load the nodes that are connected to
// the "nameserver" edge. The optional arguments are used to configure the query builder of the edge.
func (deq *DNSEntryQuery) WithNameserver(opts ...func(*NameserverQuery)) *DNSEntryQuery {
	query := (&NameserverClient{config: deq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	deq.withNameserver = query
	return deq
}

// WithScan tells the query-builder to eager-load the nodes that are connected to
// the "scan" edge. The optional arguments are used to configure the query builder of the edge.
func (deq *DNSEntryQuery) WithScan(opts ...func(*ScanQuery)) *DNSEntryQuery {
	query := (&ScanClient{config: deq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	deq.withScan = query
	return deq
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
//	client.DNSEntry.Query().
//		GroupBy(dnsentry.FieldName).
//		Aggregate(model_ent.Count()).
//		Scan(ctx, &v)
func (deq *DNSEntryQuery) GroupBy(field string, fields ...string) *DNSEntryGroupBy {
	deq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DNSEntryGroupBy{build: deq}
	grbuild.flds = &deq.ctx.Fields
	grbuild.label = dnsentry.Label
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
//	client.DNSEntry.Query().
//		Select(dnsentry.FieldName).
//		Scan(ctx, &v)
func (deq *DNSEntryQuery) Select(fields ...string) *DNSEntrySelect {
	deq.ctx.Fields = append(deq.ctx.Fields, fields...)
	sbuild := &DNSEntrySelect{DNSEntryQuery: deq}
	sbuild.label = dnsentry.Label
	sbuild.flds, sbuild.scan = &deq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DNSEntrySelect configured with the given aggregations.
func (deq *DNSEntryQuery) Aggregate(fns ...AggregateFunc) *DNSEntrySelect {
	return deq.Select().Aggregate(fns...)
}

func (deq *DNSEntryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range deq.inters {
		if inter == nil {
			return fmt.Errorf("model_ent: uninitialized interceptor (forgotten import model_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, deq); err != nil {
				return err
			}
		}
	}
	for _, f := range deq.ctx.Fields {
		if !dnsentry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model_ent: invalid field %q for query", f)}
		}
	}
	if deq.path != nil {
		prev, err := deq.path(ctx)
		if err != nil {
			return err
		}
		deq.sql = prev
	}
	return nil
}

func (deq *DNSEntryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DNSEntry, error) {
	var (
		nodes       = []*DNSEntry{}
		withFKs     = deq.withFKs
		_spec       = deq.querySpec()
		loadedTypes = [4]bool{
			deq.withDomain != nil,
			deq.withIpaddress != nil,
			deq.withNameserver != nil,
			deq.withScan != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, dnsentry.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DNSEntry).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DNSEntry{config: deq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, deq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := deq.withDomain; query != nil {
		if err := deq.loadDomain(ctx, query, nodes,
			func(n *DNSEntry) { n.Edges.Domain = []*Domain{} },
			func(n *DNSEntry, e *Domain) { n.Edges.Domain = append(n.Edges.Domain, e) }); err != nil {
			return nil, err
		}
	}
	if query := deq.withIpaddress; query != nil {
		if err := deq.loadIpaddress(ctx, query, nodes,
			func(n *DNSEntry) { n.Edges.Ipaddress = []*IPAddress{} },
			func(n *DNSEntry, e *IPAddress) { n.Edges.Ipaddress = append(n.Edges.Ipaddress, e) }); err != nil {
			return nil, err
		}
	}
	if query := deq.withNameserver; query != nil {
		if err := deq.loadNameserver(ctx, query, nodes,
			func(n *DNSEntry) { n.Edges.Nameserver = []*Nameserver{} },
			func(n *DNSEntry, e *Nameserver) { n.Edges.Nameserver = append(n.Edges.Nameserver, e) }); err != nil {
			return nil, err
		}
	}
	if query := deq.withScan; query != nil {
		if err := deq.loadScan(ctx, query, nodes,
			func(n *DNSEntry) { n.Edges.Scan = []*Scan{} },
			func(n *DNSEntry, e *Scan) { n.Edges.Scan = append(n.Edges.Scan, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (deq *DNSEntryQuery) loadDomain(ctx context.Context, query *DomainQuery, nodes []*DNSEntry, init func(*DNSEntry), assign func(*DNSEntry, *Domain)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*DNSEntry)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Domain(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(dnsentry.DomainColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.dns_entry_domain
		if fk == nil {
			return fmt.Errorf(`foreign-key "dns_entry_domain" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "dns_entry_domain" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (deq *DNSEntryQuery) loadIpaddress(ctx context.Context, query *IPAddressQuery, nodes []*DNSEntry, init func(*DNSEntry), assign func(*DNSEntry, *IPAddress)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*DNSEntry)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.IPAddress(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(dnsentry.IpaddressColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.dns_entry_ipaddress
		if fk == nil {
			return fmt.Errorf(`foreign-key "dns_entry_ipaddress" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "dns_entry_ipaddress" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (deq *DNSEntryQuery) loadNameserver(ctx context.Context, query *NameserverQuery, nodes []*DNSEntry, init func(*DNSEntry), assign func(*DNSEntry, *Nameserver)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*DNSEntry)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Nameserver(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(dnsentry.NameserverColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.dns_entry_nameserver
		if fk == nil {
			return fmt.Errorf(`foreign-key "dns_entry_nameserver" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "dns_entry_nameserver" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (deq *DNSEntryQuery) loadScan(ctx context.Context, query *ScanQuery, nodes []*DNSEntry, init func(*DNSEntry), assign func(*DNSEntry, *Scan)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*DNSEntry)
	nids := make(map[int]map[*DNSEntry]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(dnsentry.ScanTable)
		s.Join(joinT).On(s.C(scan.FieldID), joinT.C(dnsentry.ScanPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(dnsentry.ScanPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(dnsentry.ScanPrimaryKey[1]))
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
					nids[inValue] = map[*DNSEntry]struct{}{byID[outValue]: {}}
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

func (deq *DNSEntryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := deq.querySpec()
	_spec.Node.Columns = deq.ctx.Fields
	if len(deq.ctx.Fields) > 0 {
		_spec.Unique = deq.ctx.Unique != nil && *deq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, deq.driver, _spec)
}

func (deq *DNSEntryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(dnsentry.Table, dnsentry.Columns, sqlgraph.NewFieldSpec(dnsentry.FieldID, field.TypeInt))
	_spec.From = deq.sql
	if unique := deq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if deq.path != nil {
		_spec.Unique = true
	}
	if fields := deq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dnsentry.FieldID)
		for i := range fields {
			if fields[i] != dnsentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := deq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := deq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := deq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := deq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (deq *DNSEntryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(deq.driver.Dialect())
	t1 := builder.Table(dnsentry.Table)
	columns := deq.ctx.Fields
	if len(columns) == 0 {
		columns = dnsentry.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if deq.sql != nil {
		selector = deq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if deq.ctx.Unique != nil && *deq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range deq.predicates {
		p(selector)
	}
	for _, p := range deq.order {
		p(selector)
	}
	if offset := deq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := deq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DNSEntryGroupBy is the group-by builder for DNSEntry entities.
type DNSEntryGroupBy struct {
	selector
	build *DNSEntryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (degb *DNSEntryGroupBy) Aggregate(fns ...AggregateFunc) *DNSEntryGroupBy {
	degb.fns = append(degb.fns, fns...)
	return degb
}

// Scan applies the selector query and scans the result into the given value.
func (degb *DNSEntryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, degb.build.ctx, ent.OpQueryGroupBy)
	if err := degb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DNSEntryQuery, *DNSEntryGroupBy](ctx, degb.build, degb, degb.build.inters, v)
}

func (degb *DNSEntryGroupBy) sqlScan(ctx context.Context, root *DNSEntryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(degb.fns))
	for _, fn := range degb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*degb.flds)+len(degb.fns))
		for _, f := range *degb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*degb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := degb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DNSEntrySelect is the builder for selecting fields of DNSEntry entities.
type DNSEntrySelect struct {
	*DNSEntryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (des *DNSEntrySelect) Aggregate(fns ...AggregateFunc) *DNSEntrySelect {
	des.fns = append(des.fns, fns...)
	return des
}

// Scan applies the selector query and scans the result into the given value.
func (des *DNSEntrySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, des.ctx, ent.OpQuerySelect)
	if err := des.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DNSEntryQuery, *DNSEntrySelect](ctx, des.DNSEntryQuery, des, des.inters, v)
}

func (des *DNSEntrySelect) sqlScan(ctx context.Context, root *DNSEntryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(des.fns))
	for _, fn := range des.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*des.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := des.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
