package schema

// This package defines the core schema of graph objects available to dnh.
//

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Scan holds the schema definition for the Scan entity.
type Example struct {
	ent.Schema
}

// Fields of the Scan.
func (Example) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Time("timestamp"),
	}
}

// Edges of the Scan.
func (Example) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("<identifier>", <struct>.Type).extra(),
		// One 2 One
		edge.To("ipaddress", IPAddress.Type).Unique(),
		// One 2 One Self
		edge.To("next", ASNInfo.Type),
		edge.To("dnsentry", DNSEntry.Type),
		edge.To("domain", Domain.Type),
		edge.To("paths", Path.Type),
		edge.To("nameserver", Nameserver.Type),
		edge.To("registrar", Registrar.Type),
		edge.To("whois", Whois.Type),
	}
}
