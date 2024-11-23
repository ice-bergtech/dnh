// Code generated by ent, DO NOT EDIT.

package domain

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the domain type in the database.
	Label = "domain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTimeFirst holds the string denoting the time_first field in the database.
	FieldTimeFirst = "time_first"
	// FieldTimeLast holds the string denoting the time_last field in the database.
	FieldTimeLast = "time_last"
	// EdgeNameserver holds the string denoting the nameserver edge name in mutations.
	EdgeNameserver = "nameserver"
	// EdgeDomain holds the string denoting the domain edge name in mutations.
	EdgeDomain = "domain"
	// EdgeDnsentry holds the string denoting the dnsentry edge name in mutations.
	EdgeDnsentry = "dnsentry"
	// EdgeIpaddress holds the string denoting the ipaddress edge name in mutations.
	EdgeIpaddress = "ipaddress"
	// EdgePath holds the string denoting the path edge name in mutations.
	EdgePath = "path"
	// Table holds the table name of the domain in the database.
	Table = "domains"
	// NameserverTable is the table that holds the nameserver relation/edge.
	NameserverTable = "nameservers"
	// NameserverInverseTable is the table name for the Nameserver entity.
	// It exists in this package in order to avoid circular dependency with the "nameserver" package.
	NameserverInverseTable = "nameservers"
	// NameserverColumn is the table column denoting the nameserver relation/edge.
	NameserverColumn = "domain_nameserver"
	// DomainTable is the table that holds the domain relation/edge. The primary key declared below.
	DomainTable = "domain_domain"
	// DnsentryTable is the table that holds the dnsentry relation/edge.
	DnsentryTable = "dns_entries"
	// DnsentryInverseTable is the table name for the DNSEntry entity.
	// It exists in this package in order to avoid circular dependency with the "dnsentry" package.
	DnsentryInverseTable = "dns_entries"
	// DnsentryColumn is the table column denoting the dnsentry relation/edge.
	DnsentryColumn = "domain_dnsentry"
	// IpaddressTable is the table that holds the ipaddress relation/edge.
	IpaddressTable = "ip_addresses"
	// IpaddressInverseTable is the table name for the IPAddress entity.
	// It exists in this package in order to avoid circular dependency with the "ipaddress" package.
	IpaddressInverseTable = "ip_addresses"
	// IpaddressColumn is the table column denoting the ipaddress relation/edge.
	IpaddressColumn = "domain_ipaddress"
	// PathTable is the table that holds the path relation/edge.
	PathTable = "paths"
	// PathInverseTable is the table name for the Path entity.
	// It exists in this package in order to avoid circular dependency with the "path" package.
	PathInverseTable = "paths"
	// PathColumn is the table column denoting the path relation/edge.
	PathColumn = "domain_path"
)

// Columns holds all SQL columns for domain fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTimeFirst,
	FieldTimeLast,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "domains"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"scan_domain",
	"whois_domain",
}

var (
	// DomainPrimaryKey and DomainColumn2 are the table columns denoting the
	// primary key for the domain relation (M2M).
	DomainPrimaryKey = []string{"domain_id", "domain_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Domain queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTimeFirst orders the results by the time_first field.
func ByTimeFirst(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeFirst, opts...).ToFunc()
}

// ByTimeLast orders the results by the time_last field.
func ByTimeLast(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeLast, opts...).ToFunc()
}

// ByNameserverCount orders the results by nameserver count.
func ByNameserverCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNameserverStep(), opts...)
	}
}

// ByNameserver orders the results by nameserver terms.
func ByNameserver(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNameserverStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDomainCount orders the results by domain count.
func ByDomainCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDomainStep(), opts...)
	}
}

// ByDomain orders the results by domain terms.
func ByDomain(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDomainStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDnsentryCount orders the results by dnsentry count.
func ByDnsentryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDnsentryStep(), opts...)
	}
}

// ByDnsentry orders the results by dnsentry terms.
func ByDnsentry(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDnsentryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIpaddressCount orders the results by ipaddress count.
func ByIpaddressCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIpaddressStep(), opts...)
	}
}

// ByIpaddress orders the results by ipaddress terms.
func ByIpaddress(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIpaddressStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPathCount orders the results by path count.
func ByPathCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPathStep(), opts...)
	}
}

// ByPath orders the results by path terms.
func ByPath(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPathStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newNameserverStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NameserverInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NameserverTable, NameserverColumn),
	)
}
func newDomainStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, DomainTable, DomainPrimaryKey...),
	)
}
func newDnsentryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DnsentryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DnsentryTable, DnsentryColumn),
	)
}
func newIpaddressStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IpaddressInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IpaddressTable, IpaddressColumn),
	)
}
func newPathStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PathInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PathTable, PathColumn),
	)
}
