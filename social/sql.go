package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"

	"github.com/pulkitbhardwaj/alpha/social/internal"
)

func NewSQLGraph(lc fx.Lifecycle) (internal.Graph, error) {
	var (
		graph internal.Graph
		err   error
	)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			graph, err = internal.Open("sqlite3", "file:ent?mode=memory&_fk=1")
			if err != nil {
				return err
			}
			fmt.Println("Starting sqlite3 Database")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return graph.Close()
		},
	})

	return graph, nil
}
