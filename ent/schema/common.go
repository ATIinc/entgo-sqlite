package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type IDAndCreatedAt struct {
	mixin.Schema
}

func (IDAndCreatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("CreatedAt").Default(func() time.Time {
			return time.Now()
		}),
	}
}
