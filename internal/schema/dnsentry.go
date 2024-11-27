package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DNSEntry holds the schema definition for the DNSEntry entity.
type DNSEntry struct {
	ent.Schema
}

// Fields of the DNSEntry.
func (DNSEntry) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type"),
		field.String("value"),
		field.Int("ttl"),
		field.Time("time_first"),
		field.Time("time_last"),
	}
}

// Edges of the DNSEntry.
func (DNSEntry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("domain", Domain.Type),
		edge.To("ipaddress", IPAddress.Type),
		edge.To("nameserver", Nameserver.Type),
		//
		edge.From("scan", ScanJob.Type).Ref("dnsentry"),
	}
}
