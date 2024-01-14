package services

import (
	"context"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/graph/model"
)

type Services interface {
	GetRepoBasedOnFilters(ctx context.Context, input *model.RepoFilterInput, pageNo *int, pageSize *int) (*model.RepoListResponse, error)
}

type repoServices struct {
	dbQueries *database.Queries
}

func NewServices(dbQueries *database.Queries) *repoServices {
	return &repoServices{
		dbQueries: dbQueries,
	}
}
