package handler

import (
	"context"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/graph/model"
)

type Handlers interface {
	RegisterUserHandler(ctx context.Context, inputNormal *model.RegisterUserInput, inputOAuth *model.RegisterUserOAuthInput) (*database.User, error)
	UpdateUserHandler(ctx context.Context, updateUserInput *model.UpdateUserInput) (*database.User, error)
	DeleteUserHandler(ctx context.Context, userId string) (*database.User, error)
	GetUserHandler(ctx context.Context, userId string) (*database.User, error)
}

type userHandler struct {
	dbQueries *database.Queries
}

func NewUserHandler(dbQueries *database.Queries) *userHandler {
	return &userHandler{
		dbQueries: dbQueries,
	}
}
