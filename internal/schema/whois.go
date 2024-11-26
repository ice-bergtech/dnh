package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Whois holds the schema definition for the Whois entity.
type Whois struct {
	ent.Schema
}

// Fields of the Whois.
func (Whois) Fields() []ent.Field {
	return []ent.Field{
		field.String("query"),
		field.String("server"),
		field.String("raw"),
		field.String("country"),
		field.Time("created"),
		field.Time("updated"),
		field.Time("time_first"),
		field.Time("time_last"),
	}
}

// Edges of the Whois.
func (Whois) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("iprange", IPAddress.Type),
		edge.To("domain", Domain.Type),
		edge.To("asn", ASNInfo.Type),
		edge.To("registrar", Registrar.Type),
		edge.To("nameserver", Nameserver.Type),
		//
		edge.From("scan", Scan.Type).Ref("whois"),
	}
}
