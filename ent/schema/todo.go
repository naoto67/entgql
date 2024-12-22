package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/naoto67/entgql/ent/schema/puuid"
)

type Todo struct {
	ent.Schema
}

// Fields returns todo fields.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput),
			),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
				"Pending", "PENDING",
			).
			Annotations(
				entgql.OrderField("STATUS"),
			),
		field.Int("priority").
			Default(0).
			Annotations(
				entgql.OrderField("PRIORITY_ORDER"),
				entgql.MapsTo("priorityOrder"),
			),
		field.Text("text").
			NotEmpty().
			Annotations(
				entgql.OrderField("TEXT"),
			),
		field.Bytes("blob").
			Annotations(
				entgql.Skip(),
			).
			Optional(),
		field.JSON("init", map[string]any{}).
			Optional().
			Annotations(entgql.Type("Map")),
		field.Int("value").
			Default(0),
	}
}

// Annotations returns Todo annotations.
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField().Description("This is the todo item"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// "TD" declared once.
		puuid.MixinWithPrefix("TD"),
	}
}
