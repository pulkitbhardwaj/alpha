//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("migrate called")

	ex, err := entgql.NewExtension(
		entgql.WithRelaySpec(true),
		entgql.WithNodeDescriptor(true),
		entgql.WithWhereFilters(true),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaGenerator(),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithSchemaPath("internal.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	if err := entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

	// Create a local migration directory.
	dir, err := migrate.NewLocalDir("./migrations")
	if err != nil {
		log.Fatalln(err)
	}
	// Load the graph.
	graph, err := entc.LoadGraph("./schema", &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	tbls, err := graph.Tables()
	if err != nil {
		log.Fatalln(err)
	}
	// Open connection to the database.
	drv, err := sql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalln(err)
	}
	// Inspect the current database state and compare it with the graph.
	m, err := schema.NewMigrate(drv, schema.WithDir(dir))
	if err != nil {
		log.Fatalln(err)
	}
	if err := m.NamedDiff(context.Background(), os.Args[1], tbls...); err != nil {
		log.Fatalln(err)
	}
}
