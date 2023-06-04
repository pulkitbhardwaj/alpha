package internal

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/google/uuid"
)

type Graph interface {
	entgql.TxOpener

	// Noder returns the graph node with the given ID and options
	Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (Noder, error)

	// Noders returns all the graph node with the given IDs and options
	Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error)

	// Close closes the database connection and prevents new graph queries.
	Close() error
}
