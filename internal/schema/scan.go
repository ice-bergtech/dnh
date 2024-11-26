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
		//M-2-O edge.From("scan", Scan.Type).Ref("<edge.To(name)>").Unique()
		edge.To("ipaddress", IPAddress.Type),
		edge.To("asninfo", ASNInfo.Type),
		edge.To("dnsentry", DNSEntry.Type),
		edge.To("domain", Domain.Type),
		edge.To("path", Path.Type),
		edge.To("nameserver", Nameserver.Type),
		edge.To("registrar", Registrar.Type),
		edge.To("whois", Whois.Type),
	}
}
