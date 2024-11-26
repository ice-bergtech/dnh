// Code generated by ent, DO NOT EDIT.

package nameserver

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the nameserver type in the database.
	Label = "nameserver"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTimeFirst holds the string denoting the time_first field in the database.
	FieldTimeFirst = "time_first"
	// FieldTimeLast holds the string denoting the time_last field in the database.
	FieldTimeLast = "time_last"
	// EdgeIpaddress holds the string denoting the ipaddress edge name in mutations.
	EdgeIpaddress = "ipaddress"
	// EdgeScan holds the string denoting the scan edge name in mutations.
	EdgeScan = "scan"
	// EdgeDnsentry holds the string denoting the dnsentry edge name in mutations.
	EdgeDnsentry = "dnsentry"
	// EdgeDomain holds the string denoting the domain edge name in mutations.
	EdgeDomain = "domain"
	// EdgeWhois holds the string denoting the whois edge name in mutations.
	EdgeWhois = "whois"
	// Table holds the table name of the nameserver in the database.
	Table = "nameservers"
	// IpaddressTable is the table that holds the ipaddress relation/edge. The primary key declared below.
	IpaddressTable = "nameserver_ipaddress"
	// IpaddressInverseTable is the table name for the IPAddress entity.
	// It exists in this package in order to avoid circular dependency with the "ipaddress" package.
	IpaddressInverseTable = "ip_addresses"
	// ScanTable is the table that holds the scan relation/edge. The primary key declared below.
	ScanTable = "scan_nameserver"
	// ScanInverseTable is the table name for the Scan entity.
	// It exists in this package in order to avoid circular dependency with the "scan" package.
	ScanInverseTable = "scans"
	// DnsentryTable is the table that holds the dnsentry relation/edge. The primary key declared below.
	DnsentryTable = "scan_nameserver"
	// DnsentryInverseTable is the table name for the Scan entity.
	// It exists in this package in order to avoid circular dependency with the "scan" package.
	DnsentryInverseTable = "scans"
	// DomainTable is the table that holds the domain relation/edge. The primary key declared below.
	DomainTable = "domain_nameserver"
	// DomainInverseTable is the table name for the Domain entity.
	// It exists in this package in order to avoid circular dependency with the "domain" package.
	DomainInverseTable = "domains"
	// WhoisTable is the table that holds the whois relation/edge. The primary key declared below.
	WhoisTable = "whois_nameserver"
	// WhoisInverseTable is the table name for the Whois entity.
	// It exists in this package in order to avoid circular dependency with the "whois" package.
	WhoisInverseTable = "whois"
)

// Columns holds all SQL columns for nameserver fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTimeFirst,
	FieldTimeLast,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "nameservers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"dns_entry_nameserver",
	"example_nameserver",
}

var (
	// IpaddressPrimaryKey and IpaddressColumn2 are the table columns denoting the
	// primary key for the ipaddress relation (M2M).
	IpaddressPrimaryKey = []string{"nameserver_id", "ip_address_id"}
	// ScanPrimaryKey and ScanColumn2 are the table columns denoting the
	// primary key for the scan relation (M2M).
	ScanPrimaryKey = []string{"scan_id", "nameserver_id"}
	// DnsentryPrimaryKey and DnsentryColumn2 are the table columns denoting the
	// primary key for the dnsentry relation (M2M).
	DnsentryPrimaryKey = []string{"scan_id", "nameserver_id"}
	// DomainPrimaryKey and DomainColumn2 are the table columns denoting the
	// primary key for the domain relation (M2M).
	DomainPrimaryKey = []string{"domain_id", "nameserver_id"}
	// WhoisPrimaryKey and WhoisColumn2 are the table columns denoting the
	// primary key for the whois relation (M2M).
	WhoisPrimaryKey = []string{"whois_id", "nameserver_id"}
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

// OrderOption defines the ordering options for the Nameserver queries.
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

// ByScanCount orders the results by scan count.
func ByScanCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newScanStep(), opts...)
	}
}

// ByScan orders the results by scan terms.
func ByScan(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newScanStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByWhoisCount orders the results by whois count.
func ByWhoisCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWhoisStep(), opts...)
	}
}

// ByWhois orders the results by whois terms.
func ByWhois(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWhoisStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newIpaddressStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IpaddressInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, IpaddressTable, IpaddressPrimaryKey...),
	)
}
func newScanStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScanInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ScanTable, ScanPrimaryKey...),
	)
}
func newDnsentryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DnsentryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, DnsentryTable, DnsentryPrimaryKey...),
	)
}
func newDomainStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DomainInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, DomainTable, DomainPrimaryKey...),
	)
}
func newWhoisStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WhoisInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, WhoisTable, WhoisPrimaryKey...),
	)
}
