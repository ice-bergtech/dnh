package model_ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --template="../schema/ent-templates.tmpl" --target ../model_ent ../schema

import (
	_ "github.com/ice-bergtech/dnh/src/internal/schema"
)
