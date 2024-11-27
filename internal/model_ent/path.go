// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/path"
)

// Path is the model entity for the Path schema.
type Path struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PathQuery when eager-loading is set.
	Edges         PathEdges `json:"edges"`
	example_paths *string
	selectValues  sql.SelectValues
}

// PathEdges holds the relations/edges for other nodes in the graph.
type PathEdges struct {
	// Domain holds the value of the domain edge.
	Domain []*Domain `json:"domain,omitempty"`
	// Scan holds the value of the scan edge.
	Scan []*ScanJob `json:"scan,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DomainOrErr returns the Domain value or an error if the edge
// was not loaded in eager-loading.
func (e PathEdges) DomainOrErr() ([]*Domain, error) {
	if e.loadedTypes[0] {
		return e.Domain, nil
	}
	return nil, &NotLoadedError{edge: "domain"}
}

// ScanOrErr returns the Scan value or an error if the edge
// was not loaded in eager-loading.
func (e PathEdges) ScanOrErr() ([]*ScanJob, error) {
	if e.loadedTypes[1] {
		return e.Scan, nil
	}
	return nil, &NotLoadedError{edge: "scan"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Path) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case path.FieldID:
			values[i] = new(sql.NullInt64)
		case path.FieldPath:
			values[i] = new(sql.NullString)
		case path.ForeignKeys[0]: // example_paths
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Path fields.
func (pa *Path) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case path.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case path.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				pa.Path = value.String
			}
		case path.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field example_paths", values[i])
			} else if value.Valid {
				pa.example_paths = new(string)
				*pa.example_paths = value.String
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Path.
// This includes values selected through modifiers, order, etc.
func (pa *Path) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// QueryDomain queries the "domain" edge of the Path entity.
func (pa *Path) QueryDomain() *DomainQuery {
	return NewPathClient(pa.config).QueryDomain(pa)
}

// QueryScan queries the "scan" edge of the Path entity.
func (pa *Path) QueryScan() *ScanJobQuery {
	return NewPathClient(pa.config).QueryScan(pa)
}

// Update returns a builder for updating this Path.
// Note that you need to call Path.Unwrap() before calling this method if this Path
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Path) Update() *PathUpdateOne {
	return NewPathClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Path entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Path) Unwrap() *Path {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("model_ent: Path is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Path) String() string {
	var builder strings.Builder
	builder.WriteString("Path(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("path=")
	builder.WriteString(pa.Path)
	builder.WriteByte(')')
	return builder.String()
}

// Paths is a parsable slice of Path.
type Paths []*Path
