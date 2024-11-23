package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IPAddress holds the schema definition for the IPAddress entity.
type IPAddress struct {
	ent.Schema
}

// Fields of the IPAddress.
func (IPAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip"),
		field.String("mask"),
	}
}

func (IPAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("asninfo", ASNInfo.Type),
	}
}
