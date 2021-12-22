package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CopyMsg holds the schema definition for the CopyMsg entity.
type CopyMsg struct {
	ent.Schema
}

// Fields of the CopyMsg.
func (CopyMsg) Fields() []ent.Field {
	return []ent.Field{
		field.String("body"),
		field.String("from").Optional(), // 来源那个机器
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the CopyMsg.
func (CopyMsg) Edges() []ent.Edge {
	return nil
}
