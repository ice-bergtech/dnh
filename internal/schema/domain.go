package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

// Fields of the Domain.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("time_first"),
		field.Time("time_last"),
	}
}

// Edges of the Domain.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nameserver", Nameserver.Type),
		edge.To("domain", Domain.Type),
		edge.To("dnsentry", DNSEntry.Type),
		edge.To("ipaddress", IPAddress.Type),
		edge.To("path", Path.Type),
	}
}
