package schema

import (
	"entgo.io/contrib/entproto"
	_ "entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive().
			Annotations(
				entproto.Field(1),
			),
		field.String("name").
			Annotations(
				entproto.Field(2),
			).
			Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Service(),
	}
}
