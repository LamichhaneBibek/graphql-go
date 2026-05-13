package main

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/LamichhaneBibek/graphql-go/config"
	"github.com/LamichhaneBibek/graphql-go/graph"
	"github.com/LamichhaneBibek/graphql-go/internal/auth"
)

func newServer(resolver *graph.Resolver, cfg *config.Config) *http.Server {
	router := http.NewServeMux()

	gqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: resolver}),
	)
	gqlHandler.AddTransport(transport.POST{})
	gqlHandler.Use(extension.Introspection{})
	gqlHandler.Use(extension.FixedComplexityLimit(100))

	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	router.Handle("/query", auth.Middleware(gqlHandler))
	router.Handle("/health", http.HandlerFunc(healthHandler))

	return &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
