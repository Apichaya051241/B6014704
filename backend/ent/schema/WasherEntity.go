package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// WasherEntity holds the schema definition for the WasherEntity entity.
type WasherEntity struct {
	ent.Schema
}

// Fields of the WasherEntity.
func (WasherEntity) Fields() []ent.Field {
	return []ent.Field{
		field.String("wash").NotEmpty(),
	}
}

// Edges of the WasherEntity.
func (WasherEntity) Edges() []ent.Edge {
	return nil
}
