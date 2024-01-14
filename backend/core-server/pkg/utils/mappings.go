package utils

import (
	"time"

	"github.com/SohamRatnaparkhi/git-store/backend/core-server/db/database"
	"github.com/SohamRatnaparkhi/git-store/backend/core-server/graph/model"
)

func MapRepository(repo database.Repository, success bool, message string) *model.RepoResponse {
	mappedRepo := &model.RepoResponse{
		Success: success,
		Message: &message,
		Data: &model.Repo{
			RepoID:      repo.RepoID.String(),
			UserID:      repo.UserID.String(),
			Name:        repo.Name,
			URL:         repo.Url.String,
			Platform:    repo.Platform,
			Visibility:  model.Visibility(repo.Visibility),
			IsRelease:   repo.IsRelease,
			IsBackup:    repo.IsBackup,
			Description: &repo.Description.String,
			CreatedAt:   timeToString(repo.CreatedAt),
			UpdatedAt:   timeToString(repo.UpdatedAt),
		},
	}
	return mappedRepo
}

func MapRepoList(repos []database.Repository, success bool, message string) *model.RepoListResponse {
	mappedRepos := make([]*model.Repo, len(repos))
	for i, repo := range repos {
		mappedRepos[i] = &model.Repo{
			RepoID:      repo.RepoID.String(),
			UserID:      repo.UserID.String(),
			Name:        repo.Name,
			URL:         repo.Url.String,
			Platform:    repo.Platform,
			Visibility:  model.Visibility(repo.Visibility),
			IsRelease:   repo.IsRelease,
			IsBackup:    repo.IsBackup,
			Description: &repo.Description.String,
			CreatedAt:   timeToString(repo.CreatedAt),
			UpdatedAt:   timeToString(repo.UpdatedAt),
		}
	}
	total := len(repos)
	mappedRepoList := &model.RepoListResponse{
		Success: success,
		Message: &message,
		Data: &model.RepoList{
			Repos: mappedRepos,
			Total: &total,
		},
	}
	return mappedRepoList
}

func timeToString(time time.Time) *string {
	if time.IsZero() {
		return nil
	}
	timeString := time.String()
	return &timeString
}
