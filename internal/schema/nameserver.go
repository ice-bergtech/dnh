package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Nameserver holds the schema definition for the Nameserver entity.
type Nameserver struct {
	ent.Schema
}

// Fields of the Nameserver.
func (Nameserver) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("time_first"),
		field.Time("time_last"),
	}
}

// Edges of the Nameserver.
func (Nameserver) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ipaddress", IPAddress.Type),
		//
		edge.From("scan", Scan.Type).Ref("nameserver"),
		edge.From("dnsentry", Scan.Type).Ref("nameserver"),
		edge.From("domain", Domain.Type).Ref("nameserver"),
		edge.From("whois", Whois.Type).Ref("nameserver"),
	}
}
