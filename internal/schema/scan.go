package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Scan holds the schema definition for the Scan entity.
type Scan struct {
	ent.Schema
}

// Fields of the Scan.
func (Scan) Fields() []ent.Field {
	return []ent.Field{
		field.String("scanid").Unique(),
		field.String("input"),
		field.String("type"),
		field.Time("timestamp"),
	}
}

// Edges of the Scan.
func (Scan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("scanjob", ScanJob.Type),
	}
}

// Scan holds the schema definition for the Scan entity.
type ScanJob struct {
	ent.Schema
}

type Status int

const (
	Pending Status = iota
	Running
	Completed
	Failed
)

func (s Status) Str() string {
	switch s {
	case Running:
		return "Running"
	case Completed:
		return "Completed"
	case Failed:
		return "Failed"
	case Pending:
		fallthrough
	default:
		return "Pending"
	}
}

// Fields of the Scan.
func (ScanJob) Fields() []ent.Field {
	return []ent.Field{
		field.String("scanid"),
		field.Int("timeout"),
		field.Enum("status").Values(Pending.Str(), Running.Str(), Completed.Str(), Failed.Str()), // Enum values
		field.String("input"),
		field.String("output"),
		field.String("command"),
		field.Time("timestamp"),
	}
}

// Edges of the Scan.
func (ScanJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ipaddress", IPAddress.Type),
		edge.To("asninfo", ASNInfo.Type),
		edge.To("dnsentry", DNSEntry.Type),
		edge.To("domain", Domain.Type),
		edge.To("path", Path.Type),
		edge.To("nameserver", Nameserver.Type),
		edge.To("registrar", Registrar.Type),
		edge.To("whois", Whois.Type),
		edge.From("scan", Scan.Type).Ref("scanjob"),
	}
}
