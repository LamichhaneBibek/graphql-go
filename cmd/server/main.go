package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LamichhaneBibek/graphql-go/config"
	"github.com/LamichhaneBibek/graphql-go/graph"
	"github.com/LamichhaneBibek/graphql-go/internal/models"
	"github.com/LamichhaneBibek/graphql-go/internal/repository"
	"github.com/LamichhaneBibek/graphql-go/internal/seed"
	"github.com/LamichhaneBibek/graphql-go/internal/service"
)

func main() {
	cfg := config.Load()
	db := config.ConnectDB(cfg)

	db.AutoMigrate(
		&models.Permission{},
		&models.Role{},
		&models.User{},
		&models.Post{},
	)

	seed.Roles(db)

	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	resolver := &graph.Resolver{
		AuthService: service.NewAuthService(userRepo),
		UserService: service.NewUserService(userRepo, roleRepo),
		PostService: service.NewPostService(postRepo),
	}

	srv := newServer(resolver, cfg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("server listening on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-quit
	log.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
