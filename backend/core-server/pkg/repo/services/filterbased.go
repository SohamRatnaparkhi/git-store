package services

import (
	"context"
	"errors"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/graph/model"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/pkg/middlewares"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/pkg/utils"
	"github.com/google/uuid"
)

func (r *repoServices) GetRepoBasedOnFilters(ctx context.Context, input *model.RepoFilterInput, pageNo *int, pageSize *int) (*model.RepoListResponse, error) {
	userFromContext := ctx.Value(middlewares.UserClaims).(*database.User)
	userIdUUID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, err
	}
	if userFromContext.UserID != userIdUUID {
		return nil, errors.New("unauthorized")
	}
	// types - (userid, (name), (visibility), (visibility, is_backup) (visibility, is_release), (platform), (platform, is_backup), (platform, is_release))

	if input.Name != nil {
		repos, err := r.dbQueries.GetReposByUserIdAndName(ctx, database.GetReposByUserIdAndNameParams{
			UserID: userIdUUID,
			Name:   *input.Name,
		})
		if err != nil {
			return nil, err
		}
		return utils.MapRepoList(repos, true, "repo list fetched successfully"), nil
	}

	if input.Visibility != nil {

		if input.IsBackup != nil {
			repos, err := r.dbQueries.GetReposByUserIdAndVisibilityAndIsBackup(ctx, database.GetReposByUserIdAndVisibilityAndIsBackupParams{
				UserID:   userIdUUID,
				IsBackup: *input.IsBackup,
			})
			if err != nil {
				return nil, err
			}
			return utils.MapRepoList(repos, true, "repo list fetched successfully"), nil
		}

		if input.IsRelease != nil {
			repos, err := r.dbQueries.GetReposByUserIdAndVisibilityAndIsRelease(ctx, database.GetReposByUserIdAndVisibilityAndIsReleaseParams{
				UserID:    userIdUUID,
				IsRelease: *input.IsRelease,
			})
			if err != nil {
				return nil, err
			}
			return utils.MapRepoList(repos, true, "repo list fetched successfully"), nil
		}

		repos, err := r.dbQueries.GetReposByUserIdAndVisibility(ctx, database.GetReposByUserIdAndVisibilityParams{
			UserID:     userIdUUID,
			Visibility: input.Visibility.String(),
		})
		if err != nil {
			return nil, err
		}
		return utils.MapRepoList(repos, true, "repo list fetched successfully"), nil
	}

	if input.Platform != nil {

		if input.IsBackup != nil {
			repos, err := r.dbQueries.GetReposByUserIdAndPlatformAndIsBackup(ctx, database.GetReposByUserIdAndPlatformAndIsBackupParams{
				UserID:   userIdUUID,
				IsBackup: *input.IsBackup,
			})
			if err != nil {
				return nil, err
			}
			return utils.MapRepoList(repos, true, "repo list fetched successfully"), nil
		}

		if input.IsRelease != nil {
			repos, err := r.dbQueries.GetReposByUserIdAndPlatformAndIsRelease(ctx, database.GetReposByUserIdAndPlatformAndIsReleaseParams{
				UserID:    userIdUUID,
				IsRelease: *input.IsRelease,
			})
			if err != nil {
				return nil, err
			}
			return utils.MapRepoList(repos, true, "repo list fetched successfully"), nil
		}

		repos, err := r.dbQueries.GetReposByUserIdAndPlatform(ctx, database.GetReposByUserIdAndPlatformParams{
			UserID:   userIdUUID,
			Platform: *input.Platform,
		})
		if err != nil {
			return nil, err
		}
		return utils.MapRepoList(repos, true, "repo list fetched successfully"), nil
	}
	return nil, errors.New("invalid input")
}
