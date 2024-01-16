package handler

import (
	"context"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/graph/model"
)

type Handlers interface {
	CreateRepoHandler(ctx context.Context, input model.CreateRepoInput) (*model.RepoResponse, error)
	GetRepoHandler(ctx context.Context, repoID *string, url *string) (*model.RepoResponse, error)
	UpdateRepoHandler(ctx context.Context, repoId string, input model.UpdateRepoInput) (*model.RepoResponse, error)
	DeleteRepoHandler(ctx context.Context, repoId string) (*model.RepoResponse, error)
}

type repoHandlers struct {
	dbQueries *database.Queries
}

func NewRepoHandler(dbQueries *database.Queries) *repoHandlers {
	return &repoHandlers{
		dbQueries: dbQueries,
	}
}
