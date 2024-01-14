package handler

import (
	"context"
	"database/sql"
	"errors"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/graph/model"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/pkg/middlewares"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/pkg/utils"
	"github.com/google/uuid"
)

func (r *repoHandlers) CreateRepoHandler(ctx context.Context, input model.CreateRepoInput) (*model.RepoResponse, error) {
	userFromContext := ctx.Value(middlewares.UserClaims).(*database.User)
	userIdUUID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, err
	}
	if userFromContext.UserID != userIdUUID {
		return nil, errors.New("unauthorized")
	}
	repoId := uuid.New()
	visibility := input.Visibility.String()
	description := &sql.NullString{
		String: "",
		Valid:  false,
	}
	if input.Description != nil {
		description.String = *input.Description
		description.Valid = true
	}
	repo, err := r.dbQueries.CreateRepository(ctx, database.CreateRepositoryParams{
		RepoID: repoId,
		UserID: userIdUUID,
		Name:   input.Name,
		Url: sql.NullString{
			String: input.URL,
			Valid:  true,
		},
		Platform:    input.Platform,
		Visibility:  visibility,
		IsRelease:   input.IsRelease,
		IsBackup:    input.IsBackup,
		Description: *description,
	})

	if err != nil {
		return utils.MapRepository(database.Repository{}, false, err.Error()), nil
	}

	return utils.MapRepository(repo, true, "repo created successfully"), nil
}

func (r *repoHandlers) GetRepoHandler(ctx context.Context, repoID *string, url *string) (*model.RepoResponse, error) {
	userFromContext := ctx.Value(middlewares.UserClaims).(*database.User)

	if repoID != nil {
		repoIdUUID, err := uuid.Parse(*repoID)
		if err != nil {
			return nil, errors.New("invalid repo id")
		}
		repo, err := r.dbQueries.GetRepositoryByRepoId(ctx, repoIdUUID)
		if err != nil || repo.UserID != userFromContext.UserID {
			return nil, errors.New("unauthorized")
		}
		return utils.MapRepository(repo, true, "repo fetched successfully"), nil
	}

	if url != nil {
		repo, err := r.dbQueries.GetRepoByURl(ctx, sql.NullString{
			String: *url,
			Valid:  true,
		})
		if err != nil || repo.UserID != userFromContext.UserID {
			return nil, errors.New("unauthorized")
		}
	}

	return nil, errors.New("invalid input")
}
