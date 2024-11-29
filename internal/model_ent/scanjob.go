// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scanjob"
)

// ScanJob is the model entity for the ScanJob schema.
type ScanJob struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Scanid holds the value of the "scanid" field.
	Scanid string `json:"scanid,omitempty"`
	// Timeout holds the value of the "timeout" field.
	Timeout int `json:"timeout,omitempty"`
	// Status holds the value of the "status" field.
	Status scanjob.Status `json:"status,omitempty"`
	// Input holds the value of the "input" field.
	Input string `json:"input,omitempty"`
	// Output holds the value of the "output" field.
	Output string `json:"output,omitempty"`
	// Command holds the value of the "command" field.
	Command string `json:"command,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScanJobQuery when eager-loading is set.
	Edges        ScanJobEdges `json:"edges"`
	selectValues sql.SelectValues

	// Function defined by template.
	FuncPtr    *func()
	FuncParams []any
}

// ScanJobEdges holds the relations/edges for other nodes in the graph.
type ScanJobEdges struct {
	// Ipaddress holds the value of the ipaddress edge.
	Ipaddress []*IPAddress `json:"ipaddress,omitempty"`
	// Asninfo holds the value of the asninfo edge.
	Asninfo []*ASNInfo `json:"asninfo,omitempty"`
	// Dnsentry holds the value of the dnsentry edge.
	Dnsentry []*DNSEntry `json:"dnsentry,omitempty"`
	// Domain holds the value of the domain edge.
	Domain []*Domain `json:"domain,omitempty"`
	// Path holds the value of the path edge.
	Path []*Path `json:"path,omitempty"`
	// Nameserver holds the value of the nameserver edge.
	Nameserver []*Nameserver `json:"nameserver,omitempty"`
	// Registrar holds the value of the registrar edge.
	Registrar []*Registrar `json:"registrar,omitempty"`
	// Whois holds the value of the whois edge.
	Whois []*Whois `json:"whois,omitempty"`
	// Scan holds the value of the scan edge.
	Scan []*Scan `json:"scan,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// IpaddressOrErr returns the Ipaddress value or an error if the edge
// was not loaded in eager-loading.
func (e ScanJobEdges) IpaddressOrErr() ([]*IPAddress, error) {
	if e.loadedTypes[0] {
		return e.Ipaddress, nil
	}
	return nil, &NotLoadedError{edge: "ipaddress"}
}

// AsninfoOrErr returns the Asninfo value or an error if the edge
// was not loaded in eager-loading.
func (e ScanJobEdges) AsninfoOrErr() ([]*ASNInfo, error) {
	if e.loadedTypes[1] {
		return e.Asninfo, nil
	}
	return nil, &NotLoadedError{edge: "asninfo"}
}

// DnsentryOrErr returns the Dnsentry value or an error if the edge
// was not loaded in eager-loading.
func (e ScanJobEdges) DnsentryOrErr() ([]*DNSEntry, error) {
	if e.loadedTypes[2] {
		return e.Dnsentry, nil
	}
	return nil, &NotLoadedError{edge: "dnsentry"}
}

// DomainOrErr returns the Domain value or an error if the edge
// was not loaded in eager-loading.
func (e ScanJobEdges) DomainOrErr() ([]*Domain, error) {
	if e.loadedTypes[3] {
		return e.Domain, nil
	}
	return nil, &NotLoadedError{edge: "domain"}
}

// PathOrErr returns the Path value or an error if the edge
// was not loaded in eager-loading.
func (e ScanJobEdges) PathOrErr() ([]*Path, error) {
	if e.loadedTypes[4] {
		return e.Path, nil
	}
	return nil, &NotLoadedError{edge: "path"}
}

// NameserverOrErr returns the Nameserver value or an error if the edge
// was not loaded in eager-loading.
func (e ScanJobEdges) NameserverOrErr() ([]*Nameserver, error) {
	if e.loadedTypes[5] {
		return e.Nameserver, nil
	}
	return nil, &NotLoadedError{edge: "nameserver"}
}

// RegistrarOrErr returns the Registrar value or an error if the edge
// was not loaded in eager-loading.
func (e ScanJobEdges) RegistrarOrErr() ([]*Registrar, error) {
	if e.loadedTypes[6] {
		return e.Registrar, nil
	}
	return nil, &NotLoadedError{edge: "registrar"}
}

// WhoisOrErr returns the Whois value or an error if the edge
// was not loaded in eager-loading.
func (e ScanJobEdges) WhoisOrErr() ([]*Whois, error) {
	if e.loadedTypes[7] {
		return e.Whois, nil
	}
	return nil, &NotLoadedError{edge: "whois"}
}

// ScanOrErr returns the Scan value or an error if the edge
// was not loaded in eager-loading.
func (e ScanJobEdges) ScanOrErr() ([]*Scan, error) {
	if e.loadedTypes[8] {
		return e.Scan, nil
	}
	return nil, &NotLoadedError{edge: "scan"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ScanJob) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case scanjob.FieldID, scanjob.FieldTimeout:
			values[i] = new(sql.NullInt64)
		case scanjob.FieldScanid, scanjob.FieldStatus, scanjob.FieldInput, scanjob.FieldOutput, scanjob.FieldCommand:
			values[i] = new(sql.NullString)
		case scanjob.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ScanJob fields.
func (sj *ScanJob) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case scanjob.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sj.ID = int(value.Int64)
		case scanjob.FieldScanid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scanid", values[i])
			} else if value.Valid {
				sj.Scanid = value.String
			}
		case scanjob.FieldTimeout:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field timeout", values[i])
			} else if value.Valid {
				sj.Timeout = int(value.Int64)
			}
		case scanjob.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sj.Status = scanjob.Status(value.String)
			}
		case scanjob.FieldInput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field input", values[i])
			} else if value.Valid {
				sj.Input = value.String
			}
		case scanjob.FieldOutput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field output", values[i])
			} else if value.Valid {
				sj.Output = value.String
			}
		case scanjob.FieldCommand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field command", values[i])
			} else if value.Valid {
				sj.Command = value.String
			}
		case scanjob.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				sj.Timestamp = value.Time
			}
		default:
			sj.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ScanJob.
// This includes values selected through modifiers, order, etc.
func (sj *ScanJob) Value(name string) (ent.Value, error) {
	return sj.selectValues.Get(name)
}

// QueryIpaddress queries the "ipaddress" edge of the ScanJob entity.
func (sj *ScanJob) QueryIpaddress() *IPAddressQuery {
	return NewScanJobClient(sj.config).QueryIpaddress(sj)
}

// QueryAsninfo queries the "asninfo" edge of the ScanJob entity.
func (sj *ScanJob) QueryAsninfo() *ASNInfoQuery {
	return NewScanJobClient(sj.config).QueryAsninfo(sj)
}

// QueryDnsentry queries the "dnsentry" edge of the ScanJob entity.
func (sj *ScanJob) QueryDnsentry() *DNSEntryQuery {
	return NewScanJobClient(sj.config).QueryDnsentry(sj)
}

// QueryDomain queries the "domain" edge of the ScanJob entity.
func (sj *ScanJob) QueryDomain() *DomainQuery {
	return NewScanJobClient(sj.config).QueryDomain(sj)
}

// QueryPath queries the "path" edge of the ScanJob entity.
func (sj *ScanJob) QueryPath() *PathQuery {
	return NewScanJobClient(sj.config).QueryPath(sj)
}

// QueryNameserver queries the "nameserver" edge of the ScanJob entity.
func (sj *ScanJob) QueryNameserver() *NameserverQuery {
	return NewScanJobClient(sj.config).QueryNameserver(sj)
}

// QueryRegistrar queries the "registrar" edge of the ScanJob entity.
func (sj *ScanJob) QueryRegistrar() *RegistrarQuery {
	return NewScanJobClient(sj.config).QueryRegistrar(sj)
}

// QueryWhois queries the "whois" edge of the ScanJob entity.
func (sj *ScanJob) QueryWhois() *WhoisQuery {
	return NewScanJobClient(sj.config).QueryWhois(sj)
}

// QueryScan queries the "scan" edge of the ScanJob entity.
func (sj *ScanJob) QueryScan() *ScanQuery {
	return NewScanJobClient(sj.config).QueryScan(sj)
}

// Update returns a builder for updating this ScanJob.
// Note that you need to call ScanJob.Unwrap() before calling this method if this ScanJob
// was returned from a transaction, and the transaction was committed or rolled back.
func (sj *ScanJob) Update() *ScanJobUpdateOne {
	return NewScanJobClient(sj.config).UpdateOne(sj)
}

// Unwrap unwraps the ScanJob entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sj *ScanJob) Unwrap() *ScanJob {
	_tx, ok := sj.config.driver.(*txDriver)
	if !ok {
		panic("model_ent: ScanJob is not a transactional entity")
	}
	sj.config.driver = _tx.drv
	return sj
}

// String implements the fmt.Stringer.
func (sj *ScanJob) String() string {
	var builder strings.Builder
	builder.WriteString("ScanJob(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sj.ID))
	builder.WriteString("scanid=")
	builder.WriteString(sj.Scanid)
	builder.WriteString(", ")
	builder.WriteString("timeout=")
	builder.WriteString(fmt.Sprintf("%v", sj.Timeout))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", sj.Status))
	builder.WriteString(", ")
	builder.WriteString("input=")
	builder.WriteString(sj.Input)
	builder.WriteString(", ")
	builder.WriteString("output=")
	builder.WriteString(sj.Output)
	builder.WriteString(", ")
	builder.WriteString("command=")
	builder.WriteString(sj.Command)
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(sj.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ScanJobs is a parsable slice of ScanJob.
type ScanJobs []*ScanJob