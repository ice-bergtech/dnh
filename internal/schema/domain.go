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
		field.Ints("ports"),
		field.Time("time_first"),
		field.Time("time_last"),
	}
}

// Edges of the Domain.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nameserver", Nameserver.Type),
		edge.To("subdomain", Domain.Type),
		edge.To("ipaddress", IPAddress.Type),
		edge.To("path", Path.Type),
		//
		edge.From("scan", ScanJob.Type).Ref("domain"),
		edge.From("dnsentry", DNSEntry.Type).Ref("domain"),
		edge.From("registrar", Registrar.Type).Ref("domain"),
		edge.From("whois", Whois.Type).Ref("domain"),
	}
}
