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
	Edges                IPAddressEdges `json:"edges"`
	domain_ipaddress     *int
	nameserver_ipaddress *int
	scan_ipaddress       *int
	whois_iprange        *int
	selectValues         sql.SelectValues
}

// IPAddressEdges holds the relations/edges for other nodes in the graph.
type IPAddressEdges struct {
	// Asninfo holds the value of the asninfo edge.
	Asninfo []*ASNInfo `json:"asninfo,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AsninfoOrErr returns the Asninfo value or an error if the edge
// was not loaded in eager-loading.
func (e IPAddressEdges) AsninfoOrErr() ([]*ASNInfo, error) {
	if e.loadedTypes[0] {
		return e.Asninfo, nil
	}
	return nil, &NotLoadedError{edge: "asninfo"}
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
		case ipaddress.ForeignKeys[0]: // domain_ipaddress
			values[i] = new(sql.NullInt64)
		case ipaddress.ForeignKeys[1]: // nameserver_ipaddress
			values[i] = new(sql.NullInt64)
		case ipaddress.ForeignKeys[2]: // scan_ipaddress
			values[i] = new(sql.NullInt64)
		case ipaddress.ForeignKeys[3]: // whois_iprange
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
				return fmt.Errorf("unexpected type %T for edge-field domain_ipaddress", value)
			} else if value.Valid {
				ia.domain_ipaddress = new(int)
				*ia.domain_ipaddress = int(value.Int64)
			}
		case ipaddress.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field nameserver_ipaddress", value)
			} else if value.Valid {
				ia.nameserver_ipaddress = new(int)
				*ia.nameserver_ipaddress = int(value.Int64)
			}
		case ipaddress.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field scan_ipaddress", value)
			} else if value.Valid {
				ia.scan_ipaddress = new(int)
				*ia.scan_ipaddress = int(value.Int64)
			}
		case ipaddress.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field whois_iprange", value)
			} else if value.Valid {
				ia.whois_iprange = new(int)
				*ia.whois_iprange = int(value.Int64)
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
