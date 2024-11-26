// Code generated by ent, DO NOT EDIT.

package registrar

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the registrar type in the database.
	Label = "registrar"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCountryCode holds the string denoting the country_code field in the database.
	FieldCountryCode = "country_code"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldFax holds the string denoting the fax field in the database.
	FieldFax = "fax"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldTimeFirst holds the string denoting the time_first field in the database.
	FieldTimeFirst = "time_first"
	// FieldTimeLast holds the string denoting the time_last field in the database.
	FieldTimeLast = "time_last"
	// EdgeIpaddress holds the string denoting the ipaddress edge name in mutations.
	EdgeIpaddress = "ipaddress"
	// EdgeDomain holds the string denoting the domain edge name in mutations.
	EdgeDomain = "domain"
	// EdgeAsninfo holds the string denoting the asninfo edge name in mutations.
	EdgeAsninfo = "asninfo"
	// EdgeScan holds the string denoting the scan edge name in mutations.
	EdgeScan = "scan"
	// EdgeWhois holds the string denoting the whois edge name in mutations.
	EdgeWhois = "whois"
	// Table holds the table name of the registrar in the database.
	Table = "registrars"
	// IpaddressTable is the table that holds the ipaddress relation/edge. The primary key declared below.
	IpaddressTable = "registrar_ipaddress"
	// IpaddressInverseTable is the table name for the IPAddress entity.
	// It exists in this package in order to avoid circular dependency with the "ipaddress" package.
	IpaddressInverseTable = "ip_addresses"
	// DomainTable is the table that holds the domain relation/edge. The primary key declared below.
	DomainTable = "registrar_domain"
	// DomainInverseTable is the table name for the Domain entity.
	// It exists in this package in order to avoid circular dependency with the "domain" package.
	DomainInverseTable = "domains"
	// AsninfoTable is the table that holds the asninfo relation/edge. The primary key declared below.
	AsninfoTable = "registrar_asninfo"
	// AsninfoInverseTable is the table name for the ASNInfo entity.
	// It exists in this package in order to avoid circular dependency with the "asninfo" package.
	AsninfoInverseTable = "asn_infos"
	// ScanTable is the table that holds the scan relation/edge. The primary key declared below.
	ScanTable = "scan_registrar"
	// ScanInverseTable is the table name for the Scan entity.
	// It exists in this package in order to avoid circular dependency with the "scan" package.
	ScanInverseTable = "scans"
	// WhoisTable is the table that holds the whois relation/edge. The primary key declared below.
	WhoisTable = "whois_registrar"
	// WhoisInverseTable is the table name for the Whois entity.
	// It exists in this package in order to avoid circular dependency with the "whois" package.
	WhoisInverseTable = "whois"
)

// Columns holds all SQL columns for registrar fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldURL,
	FieldCountryCode,
	FieldPhone,
	FieldFax,
	FieldAddress,
	FieldSource,
	FieldTimeFirst,
	FieldTimeLast,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "registrars"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"example_registrar",
}

var (
	// IpaddressPrimaryKey and IpaddressColumn2 are the table columns denoting the
	// primary key for the ipaddress relation (M2M).
	IpaddressPrimaryKey = []string{"registrar_id", "ip_address_id"}
	// DomainPrimaryKey and DomainColumn2 are the table columns denoting the
	// primary key for the domain relation (M2M).
	DomainPrimaryKey = []string{"registrar_id", "domain_id"}
	// AsninfoPrimaryKey and AsninfoColumn2 are the table columns denoting the
	// primary key for the asninfo relation (M2M).
	AsninfoPrimaryKey = []string{"registrar_id", "asn_info_id"}
	// ScanPrimaryKey and ScanColumn2 are the table columns denoting the
	// primary key for the scan relation (M2M).
	ScanPrimaryKey = []string{"scan_id", "registrar_id"}
	// WhoisPrimaryKey and WhoisColumn2 are the table columns denoting the
	// primary key for the whois relation (M2M).
	WhoisPrimaryKey = []string{"whois_id", "registrar_id"}
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

// OrderOption defines the ordering options for the Registrar queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByCountryCode orders the results by the country_code field.
func ByCountryCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountryCode, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByFax orders the results by the fax field.
func ByFax(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFax, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
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

// ByAsninfoCount orders the results by asninfo count.
func ByAsninfoCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAsninfoStep(), opts...)
	}
}

// ByAsninfo orders the results by asninfo terms.
func ByAsninfo(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAsninfoStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newDomainStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DomainInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, DomainTable, DomainPrimaryKey...),
	)
}
func newAsninfoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AsninfoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AsninfoTable, AsninfoPrimaryKey...),
	)
}
func newScanStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScanInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ScanTable, ScanPrimaryKey...),
	)
}
func newWhoisStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WhoisInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, WhoisTable, WhoisPrimaryKey...),
	)
}
