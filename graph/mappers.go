package graph

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/LamichhaneBibek/graphql-go/graph/model"
	"github.com/LamichhaneBibek/graphql-go/internal/auth"
	"github.com/LamichhaneBibek/graphql-go/internal/models"
)

func requireAuth(ctx context.Context) (uint, error) {
	id, ok := auth.UserIDFromCtx(ctx)
	if !ok {
		return 0, errors.New("unauthenticated")
	}
	return id, nil
}

func (r *Resolver) requireRole(ctx context.Context, userID uint, role string) error {
	ok, err := r.UserService.HasRole(userID, role)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("forbidden: insufficient permissions")
	}
	return nil
}

func parseID(id string) (uint, error) {
	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, errors.New("invalid ID")
	}
	return uint(n), nil
}

func mapUser(u *models.User) *model.User {
	if u == nil {
		return nil
	}
	gql := &model.User{
		ID:        fmt.Sprint(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
	if u.Role != nil {
		gql.Role = mapRole(u.Role)
	}
	return gql
}

func mapRole(r *models.Role) *model.Role {
	if r == nil {
		return nil
	}
	gql := &model.Role{
		ID:          fmt.Sprint(r.ID),
		Name:        r.Name,
		Permissions: make([]*model.Permission, len(r.Permissions)),
	}
	for i, p := range r.Permissions {
		gql.Permissions[i] = mapPermission(&p)
	}
	return gql
}

func mapPermission(p *models.Permission) *model.Permission {
	if p == nil {
		return nil
	}
	return &model.Permission{
		ID:       fmt.Sprint(p.ID),
		Name:     p.Name,
		Resource: p.Resource,
		Action:   p.Action,
	}
}

func mapPost(p *models.Post) *model.Post {
	if p == nil {
		return nil
	}
	return &model.Post{
		ID:        fmt.Sprint(p.ID),
		Title:     p.Title,
		Content:   p.Content,
		Published: p.Published,
		Author:    mapUser(&p.Author),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
