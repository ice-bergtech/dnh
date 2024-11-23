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
	}
}
