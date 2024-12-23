package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Path holds the schema definition for the Path entity.
type Path struct {
	ent.Schema
}

// Fields of the Path.
func (Path) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
	}
}

// Edges of the Path.
func (Path) Edges() []ent.Edge {
	return []ent.Edge{
		//
		edge.From("domain", Domain.Type).Ref("path"),
		edge.From("scan", ScanJob.Type).Ref("path"),
	}
}
