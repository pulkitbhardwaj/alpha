package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"

	"github.com/pulkitbhardwaj/alpha/social/internal"
)

type GraphQLParams struct {
	fx.In

	Schema graphql.ExecutableSchema
	Graph  *internal.Client
	Store  sessions.Store
}

func Transport(params GraphQLParams) http.Handler {
	mux := mux.NewRouter()
	srv := handler.New(params.Schema)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			// Get the token from payload
			any := initPayload["authToken"]
			token, ok := any.(string)
			if !ok || token == "" {
				return nil, errors.New("authToken not found in transport payload")
			}

			// Perform token verification and authentication...
			userId := "john.doe" // e.g. userId, err := GetUserFromAuthentication(token)

			// put it in context
			ctxNew := context.WithValue(ctx, authCtxKey, userId)

			return ctxNew, nil
		},
	})
	srv.SetQueryCache(lru.New(1000))
	// srv.Use(entgql.Transactioner{TxOpener: params.Graph})
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	mux.
		Handle("/", playground.Handler("GraphQL playground", "/")).Methods("GET")
	mux.
		Handle("/graphql", srv).Methods("POST")

	return mux
}
