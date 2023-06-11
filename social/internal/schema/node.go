package schema

import (
	"time"

	"fmt"
	"io"
	"strconv"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

// -------------------------------------------------
// Mixin definition

// Node implements the ent.Mixin for sharing
// time fields with package schemas.
type Node struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable(),

		field.Time("created_at").
			Immutable().
			Default(time.Now),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func MarshalUUID(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(u.String()))
	})
}

func UnmarshalUUID(v interface{}) (u uuid.UUID, err error) {
	s, ok := v.(string)
	if !ok {
		return u, fmt.Errorf("invalid type %T, expect string", v)
	}
	return uuid.Parse(s)
}
