package internal

import (
	"context"
	"fmt"
	"log"

	"ariga.io/entcache"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
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

func SQLGraph(lc fx.Lifecycle) (Graph, error) {
	// Open the database connection.
	db, err := sql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening database", err)
	}

	// Decorates the sql.Driver with entcache.Driver.
	// Create an ent.Client.
	graph := NewClient(Driver(entcache.NewDriver(db)))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Creating migrations")

			// Tell the entcache.Driver to skip the caching layer
			// when running the schema migration.
			if graph.Schema.Create(entcache.Skip(ctx)); err != nil {
				return err
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return graph.Close()
		},
	})

	return graph, nil
}
