package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Example holds the schema definition for the Example entity.
type Example struct {
	ent.Schema
}

func (Example) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDAndCreatedAt{},
	}
}

// Fields of the Example.
func (Example) Fields() []ent.Field {
	return []ent.Field{
		field.String("ExampleString"),
		field.JSON("ExampleJSON", map[string]any{}),
	}
}

// Edges of the Example.
func (Example) Edges() []ent.Edge {
	return nil
}
