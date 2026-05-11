package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/LamichhaneBibek/graphql-go/config"
	"github.com/LamichhaneBibek/graphql-go/graph"
	"github.com/LamichhaneBibek/graphql-go/internal/auth"
	"github.com/LamichhaneBibek/graphql-go/internal/models"
)

// cmd/server/main.go
func main() {
	cfg := config.Load() // load from env vars
	db := config.ConnectDB(cfg)
	db.AutoMigrate(&models.User{})

	resolver := &graph.Resolver{DB: db} // inject DB into resolver
	srv := newServer(resolver, cfg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-quit
	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}

func newServer(resolver *graph.Resolver, cfg *config.Config) *http.Server {
	router := http.NewServeMux()

	gqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: resolver}),
	)
	gqlHandler.AddTransport(transport.POST{})
	gqlHandler.Use(extension.Introspection{})
	gqlHandler.Use(extension.FixedComplexityLimit(100))

	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	router.Handle("/query", auth.Middleware(gqlHandler))      // ← use auth.Middleware
	router.Handle("/health", http.HandlerFunc(healthHandler)) // ← define below

	return &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
