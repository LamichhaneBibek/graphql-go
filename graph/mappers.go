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
	return &model.User{
		ID:        fmt.Sprint(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
