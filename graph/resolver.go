package graph

import "github.com/LamichhaneBibek/graphql-go/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthService service.AuthService
	UserService service.UserService
	PostService service.PostService
}
