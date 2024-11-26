// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
)

// IPAddress is the model entity for the IPAddress schema.
type IPAddress struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IP holds the value of the "ip" field.
	IP string `json:"ip,omitempty"`
	// Mask holds the value of the "mask" field.
	Mask string `json:"mask,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IPAddressQuery when eager-loading is set.
	Edges               IPAddressEdges `json:"edges"`
	dns_entry_ipaddress *int
	selectValues        sql.SelectValues
}

// IPAddressEdges holds the relations/edges for other nodes in the graph.
type IPAddressEdges struct {
	// Asninfo holds the value of the asninfo edge.
	Asninfo []*ASNInfo `json:"asninfo,omitempty"`
	// Scan holds the value of the scan edge.
	Scan []*Scan `json:"scan,omitempty"`
	// Dnsentry holds the value of the dnsentry edge.
	Dnsentry []*Scan `json:"dnsentry,omitempty"`
	// Domain holds the value of the domain edge.
	Domain []*Domain `json:"domain,omitempty"`
	// Nameserver holds the value of the nameserver edge.
	Nameserver []*Nameserver `json:"nameserver,omitempty"`
	// Registrar holds the value of the registrar edge.
	Registrar []*Registrar `json:"registrar,omitempty"`
	// Whois holds the value of the whois edge.
	Whois []*Whois `json:"whois,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// AsninfoOrErr returns the Asninfo value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) AsninfoOrErr() ([]*ASNInfo, error) {
	if e.loadedTypes[0] {
		return e.Asninfo, nil
	}
	return nil, &NotLoadedError{edge: "asninfo"}
}

// ScanOrErr returns the Scan value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) ScanOrErr() ([]*Scan, error) {
	if e.loadedTypes[1] {
		return e.Scan, nil
	}
	return nil, &NotLoadedError{edge: "scan"}
}

// DnsentryOrErr returns the Dnsentry value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) DnsentryOrErr() ([]*Scan, error) {
	if e.loadedTypes[2] {
		return e.Dnsentry, nil
	}
	return nil, &NotLoadedError{edge: "dnsentry"}
}

// DomainOrErr returns the Domain value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) DomainOrErr() ([]*Domain, error) {
	if e.loadedTypes[3] {
		return e.Domain, nil
	}
	return nil, &NotLoadedError{edge: "domain"}
}

// NameserverOrErr returns the Nameserver value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) NameserverOrErr() ([]*Nameserver, error) {
	if e.loadedTypes[4] {
		return e.Nameserver, nil
	}
	return nil, &NotLoadedError{edge: "nameserver"}
}

// RegistrarOrErr returns the Registrar value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) RegistrarOrErr() ([]*Registrar, error) {
	if e.loadedTypes[5] {
		return e.Registrar, nil
	}
	return nil, &NotLoadedError{edge: "registrar"}
}

// WhoisOrErr returns the Whois value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) WhoisOrErr() ([]*Whois, error) {
	if e.loadedTypes[6] {
		return e.Whois, nil
	}
	return nil, &NotLoadedError{edge: "whois"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IPAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ipaddress.FieldID:
			values[i] = new(sql.NullInt64)
		case ipaddress.FieldIP, ipaddress.FieldMask:
			values[i] = new(sql.NullString)
		case ipaddress.ForeignKeys[0]: // dns_entry_ipaddress
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IPAddress fields.
func (ia *IPAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ipaddress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ia.ID = int(value.Int64)
		case ipaddress.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				ia.IP = value.String
			}
		case ipaddress.FieldMask:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mask", values[i])
			} else if value.Valid {
				ia.Mask = value.String
			}
		case ipaddress.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field dns_entry_ipaddress", value)
			} else if value.Valid {
				ia.dns_entry_ipaddress = new(int)
				*ia.dns_entry_ipaddress = int(value.Int64)
			}
		default:
			ia.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IPAddress.
// This includes values selected through modifiers, order, etc.
func (ia *IPAddress) Value(name string) (ent.Value, error) {
	return ia.selectValues.Get(name)
}

// QueryAsninfo queries the "asninfo" edge of the IPAddress entity.
func (ia *IPAddress) QueryAsninfo() *ASNInfoQuery {
	return NewIPAddressClient(ia.config).QueryAsninfo(ia)
}

// QueryScan queries the "scan" edge of the IPAddress entity.
func (ia *IPAddress) QueryScan() *ScanQuery {
	return NewIPAddressClient(ia.config).QueryScan(ia)
}

// QueryDnsentry queries the "dnsentry" edge of the IPAddress entity.
func (ia *IPAddress) QueryDnsentry() *ScanQuery {
	return NewIPAddressClient(ia.config).QueryDnsentry(ia)
}

// QueryDomain queries the "domain" edge of the IPAddress entity.
func (ia *IPAddress) QueryDomain() *DomainQuery {
	return NewIPAddressClient(ia.config).QueryDomain(ia)
}

// QueryNameserver queries the "nameserver" edge of the IPAddress entity.
func (ia *IPAddress) QueryNameserver() *NameserverQuery {
	return NewIPAddressClient(ia.config).QueryNameserver(ia)
}

// QueryRegistrar queries the "registrar" edge of the IPAddress entity.
func (ia *IPAddress) QueryRegistrar() *RegistrarQuery {
	return NewIPAddressClient(ia.config).QueryRegistrar(ia)
}

// QueryWhois queries the "whois" edge of the IPAddress entity.
func (ia *IPAddress) QueryWhois() *WhoisQuery {
	return NewIPAddressClient(ia.config).QueryWhois(ia)
}

// Update returns a builder for updating this IPAddress.
// Note that you need to call IPAddress.Unwrap() before calling this method if this IPAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (ia *IPAddress) Update() *IPAddressUpdateOne {
	return NewIPAddressClient(ia.config).UpdateOne(ia)
}

// Unwrap unwraps the IPAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ia *IPAddress) Unwrap() *IPAddress {
	_tx, ok := ia.config.driver.(*txDriver)
	if !ok {
		panic("model_ent: IPAddress is not a transactional entity")
	}
	ia.config.driver = _tx.drv
	return ia
}

// String implements the fmt.Stringer.
func (ia *IPAddress) String() string {
	var builder strings.Builder
	builder.WriteString("IPAddress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ia.ID))
	builder.WriteString("ip=")
	builder.WriteString(ia.IP)
	builder.WriteString(", ")
	builder.WriteString("mask=")
	builder.WriteString(ia.Mask)
	builder.WriteByte(')')
	return builder.String()
}

// IPAddresses is a parsable slice of IPAddress.
type IPAddresses []*IPAddress
