package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("password").Sensitive(),
		field.String("avatar").Optional(),
		field.String("email").Unique(),
		field.String("gender").Default("male").Comment("male 男, female 女"),
		field.String("role").Default("user").Comment("admin, user"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
