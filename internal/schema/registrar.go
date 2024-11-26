package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Registrar holds the schema definition for the Registrar entity.
type Registrar struct {
	ent.Schema
}

// Fields of the Registrar.
func (Registrar) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("url"),
		field.String("country_code"),
		field.String("phone"),
		field.String("fax"),
		field.String("address"),
		field.String("source"),
		field.Time("time_first"),
		field.Time("time_last"),
	}
}

// Edges of the Registrar.
func (Registrar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ipaddress", IPAddress.Type),
		edge.To("domain", Domain.Type),
		edge.To("asninfo", ASNInfo.Type),
		//
		edge.From("scan", Scan.Type).Ref("registrar"),
		edge.From("whois", Whois.Type).Ref("registrar"),
	}
}
