// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/domain"
)

// Domain is the model entity for the Domain schema.
type Domain struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Ports holds the value of the "ports" field.
	Ports []int `json:"ports,omitempty"`
	// TimeFirst holds the value of the "time_first" field.
	TimeFirst time.Time `json:"time_first,omitempty"`
	// TimeLast holds the value of the "time_last" field.
	TimeLast time.Time `json:"time_last,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DomainQuery when eager-loading is set.
	Edges          DomainEdges `json:"edges"`
	example_domain *string
	selectValues   sql.SelectValues
}

// DomainEdges holds the relations/edges for other nodes in the graph.
type DomainEdges struct {
	// Nameserver holds the value of the nameserver edge.
	Nameserver []*Nameserver `json:"nameserver,omitempty"`
	// Subdomain holds the value of the subdomain edge.
	Subdomain []*Domain `json:"subdomain,omitempty"`
	// Ipaddress holds the value of the ipaddress edge.
	Ipaddress []*IPAddress `json:"ipaddress,omitempty"`
	// Path holds the value of the path edge.
	Path []*Path `json:"path,omitempty"`
	// Scan holds the value of the scan edge.
	Scan []*ScanJob `json:"scan,omitempty"`
	// Dnsentry holds the value of the dnsentry edge.
	Dnsentry []*DNSEntry `json:"dnsentry,omitempty"`
	// Registrar holds the value of the registrar edge.
	Registrar []*Registrar `json:"registrar,omitempty"`
	// Whois holds the value of the whois edge.
	Whois []*Whois `json:"whois,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// NameserverOrErr returns the Nameserver value or an error if the edge
// was not loaded in eager-loading.
func (e DomainEdges) NameserverOrErr() ([]*Nameserver, error) {
	if e.loadedTypes[0] {
		return e.Nameserver, nil
	}
	return nil, &NotLoadedError{edge: "nameserver"}
}

// SubdomainOrErr returns the Subdomain value or an error if the edge
// was not loaded in eager-loading.
func (e DomainEdges) SubdomainOrErr() ([]*Domain, error) {
	if e.loadedTypes[1] {
		return e.Subdomain, nil
	}
	return nil, &NotLoadedError{edge: "subdomain"}
}

// IpaddressOrErr returns the Ipaddress value or an error if the edge
// was not loaded in eager-loading.
func (e DomainEdges) IpaddressOrErr() ([]*IPAddress, error) {
	if e.loadedTypes[2] {
		return e.Ipaddress, nil
	}
	return nil, &NotLoadedError{edge: "ipaddress"}
}

// PathOrErr returns the Path value or an error if the edge
// was not loaded in eager-loading.
func (e DomainEdges) PathOrErr() ([]*Path, error) {
	if e.loadedTypes[3] {
		return e.Path, nil
	}
	return nil, &NotLoadedError{edge: "path"}
}

// ScanOrErr returns the Scan value or an error if the edge
// was not loaded in eager-loading.
func (e DomainEdges) ScanOrErr() ([]*ScanJob, error) {
	if e.loadedTypes[4] {
		return e.Scan, nil
	}
	return nil, &NotLoadedError{edge: "scan"}
}

// DnsentryOrErr returns the Dnsentry value or an error if the edge
// was not loaded in eager-loading.
func (e DomainEdges) DnsentryOrErr() ([]*DNSEntry, error) {
	if e.loadedTypes[5] {
		return e.Dnsentry, nil
	}
	return nil, &NotLoadedError{edge: "dnsentry"}
}

// RegistrarOrErr returns the Registrar value or an error if the edge
// was not loaded in eager-loading.
func (e DomainEdges) RegistrarOrErr() ([]*Registrar, error) {
	if e.loadedTypes[6] {
		return e.Registrar, nil
	}
	return nil, &NotLoadedError{edge: "registrar"}
}

// WhoisOrErr returns the Whois value or an error if the edge
// was not loaded in eager-loading.
func (e DomainEdges) WhoisOrErr() ([]*Whois, error) {
	if e.loadedTypes[7] {
		return e.Whois, nil
	}
	return nil, &NotLoadedError{edge: "whois"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Domain) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case domain.FieldPorts:
			values[i] = new([]byte)
		case domain.FieldID:
			values[i] = new(sql.NullInt64)
		case domain.FieldName:
			values[i] = new(sql.NullString)
		case domain.FieldTimeFirst, domain.FieldTimeLast:
			values[i] = new(sql.NullTime)
		case domain.ForeignKeys[0]: // example_domain
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Domain fields.
func (d *Domain) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case domain.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case domain.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case domain.FieldPorts:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ports", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &d.Ports); err != nil {
					return fmt.Errorf("unmarshal field ports: %w", err)
				}
			}
		case domain.FieldTimeFirst:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time_first", values[i])
			} else if value.Valid {
				d.TimeFirst = value.Time
			}
		case domain.FieldTimeLast:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time_last", values[i])
			} else if value.Valid {
				d.TimeLast = value.Time
			}
		case domain.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field example_domain", values[i])
			} else if value.Valid {
				d.example_domain = new(string)
				*d.example_domain = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Domain.
// This includes values selected through modifiers, order, etc.
func (d *Domain) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryNameserver queries the "nameserver" edge of the Domain entity.
func (d *Domain) QueryNameserver() *NameserverQuery {
	return NewDomainClient(d.config).QueryNameserver(d)
}

// QuerySubdomain queries the "subdomain" edge of the Domain entity.
func (d *Domain) QuerySubdomain() *DomainQuery {
	return NewDomainClient(d.config).QuerySubdomain(d)
}

// QueryIpaddress queries the "ipaddress" edge of the Domain entity.
func (d *Domain) QueryIpaddress() *IPAddressQuery {
	return NewDomainClient(d.config).QueryIpaddress(d)
}

// QueryPath queries the "path" edge of the Domain entity.
func (d *Domain) QueryPath() *PathQuery {
	return NewDomainClient(d.config).QueryPath(d)
}

// QueryScan queries the "scan" edge of the Domain entity.
func (d *Domain) QueryScan() *ScanJobQuery {
	return NewDomainClient(d.config).QueryScan(d)
}

// QueryDnsentry queries the "dnsentry" edge of the Domain entity.
func (d *Domain) QueryDnsentry() *DNSEntryQuery {
	return NewDomainClient(d.config).QueryDnsentry(d)
}

// QueryRegistrar queries the "registrar" edge of the Domain entity.
func (d *Domain) QueryRegistrar() *RegistrarQuery {
	return NewDomainClient(d.config).QueryRegistrar(d)
}

// QueryWhois queries the "whois" edge of the Domain entity.
func (d *Domain) QueryWhois() *WhoisQuery {
	return NewDomainClient(d.config).QueryWhois(d)
}

// Update returns a builder for updating this Domain.
// Note that you need to call Domain.Unwrap() before calling this method if this Domain
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Domain) Update() *DomainUpdateOne {
	return NewDomainClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Domain entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Domain) Unwrap() *Domain {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("model_ent: Domain is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Domain) String() string {
	var builder strings.Builder
	builder.WriteString("Domain(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("ports=")
	builder.WriteString(fmt.Sprintf("%v", d.Ports))
	builder.WriteString(", ")
	builder.WriteString("time_first=")
	builder.WriteString(d.TimeFirst.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("time_last=")
	builder.WriteString(d.TimeLast.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Domains is a parsable slice of Domain.
type Domains []*Domain
