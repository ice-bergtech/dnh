package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ASNInfo holds the schema definition for the ASNInfo entity.
type ASNInfo struct {
	ent.Schema
}

// Fields of the ASNInfo.
func (ASNInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("asn"),
		field.String("country"),
		field.String("registry"),
	}
}

// Edges of the ASNInfo.
func (ASNInfo) Edges() []ent.Edge {
	return nil
}
