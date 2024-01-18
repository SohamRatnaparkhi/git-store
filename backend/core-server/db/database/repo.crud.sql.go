// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: repo.crud.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createRepository = `-- name: CreateRepository :one
INSERT INTO repository (
    repo_id,
    installation_id,
    user_id,
    name,
    url,
    platform,
    visibility,
    is_release,
    is_backup,
    description
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
) RETURNING repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at
`

type CreateRepositoryParams struct {
	RepoID         uuid.UUID
	InstallationID uuid.UUID
	UserID         uuid.UUID
	Name           string
	Url            sql.NullString
	Platform       string
	Visibility     string
	IsRelease      bool
	IsBackup       bool
	Description    sql.NullString
}

func (q *Queries) CreateRepository(ctx context.Context, arg CreateRepositoryParams) (Repository, error) {
	row := q.db.QueryRowContext(ctx, createRepository,
		arg.RepoID,
		arg.InstallationID,
		arg.UserID,
		arg.Name,
		arg.Url,
		arg.Platform,
		arg.Visibility,
		arg.IsRelease,
		arg.IsBackup,
		arg.Description,
	)
	var i Repository
	err := row.Scan(
		&i.RepoID,
		&i.InstallationID,
		&i.UserID,
		&i.Name,
		&i.Url,
		&i.Platform,
		&i.Visibility,
		&i.IsRelease,
		&i.IsBackup,
		&i.IsApp,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRepoByRepoId = `-- name: DeleteRepoByRepoId :one
DELETE FROM repository WHERE repo_id = $1 RETURNING repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at
`

func (q *Queries) DeleteRepoByRepoId(ctx context.Context, repoID uuid.UUID) (Repository, error) {
	row := q.db.QueryRowContext(ctx, deleteRepoByRepoId, repoID)
	var i Repository
	err := row.Scan(
		&i.RepoID,
		&i.InstallationID,
		&i.UserID,
		&i.Name,
		&i.Url,
		&i.Platform,
		&i.Visibility,
		&i.IsRelease,
		&i.IsBackup,
		&i.IsApp,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRepoByURl = `-- name: GetRepoByURl :one
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE url = $1
`

func (q *Queries) GetRepoByURl(ctx context.Context, url sql.NullString) (Repository, error) {
	row := q.db.QueryRowContext(ctx, getRepoByURl, url)
	var i Repository
	err := row.Scan(
		&i.RepoID,
		&i.InstallationID,
		&i.UserID,
		&i.Name,
		&i.Url,
		&i.Platform,
		&i.Visibility,
		&i.IsRelease,
		&i.IsBackup,
		&i.IsApp,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getReposByUserId = `-- name: GetReposByUserId :many

SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE user_id = $1
`

// get repo pairs - (userid, (name), (visibility), (visibility, is_backup) (visibility, is_release), (platform), (platform, is_backup), (platform, is_release))
func (q *Queries) GetReposByUserId(ctx context.Context, userID uuid.UUID) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, getReposByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repository
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.RepoID,
			&i.InstallationID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Platform,
			&i.Visibility,
			&i.IsRelease,
			&i.IsBackup,
			&i.IsApp,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReposByUserIdAndName = `-- name: GetReposByUserIdAndName :many
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE user_id = $1 AND name = $2
`

type GetReposByUserIdAndNameParams struct {
	UserID uuid.UUID
	Name   string
}

func (q *Queries) GetReposByUserIdAndName(ctx context.Context, arg GetReposByUserIdAndNameParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, getReposByUserIdAndName, arg.UserID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repository
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.RepoID,
			&i.InstallationID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Platform,
			&i.Visibility,
			&i.IsRelease,
			&i.IsBackup,
			&i.IsApp,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReposByUserIdAndPlatform = `-- name: GetReposByUserIdAndPlatform :many
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE user_id = $1 AND platform = $2
`

type GetReposByUserIdAndPlatformParams struct {
	UserID   uuid.UUID
	Platform string
}

func (q *Queries) GetReposByUserIdAndPlatform(ctx context.Context, arg GetReposByUserIdAndPlatformParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, getReposByUserIdAndPlatform, arg.UserID, arg.Platform)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repository
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.RepoID,
			&i.InstallationID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Platform,
			&i.Visibility,
			&i.IsRelease,
			&i.IsBackup,
			&i.IsApp,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReposByUserIdAndPlatformAndIsBackup = `-- name: GetReposByUserIdAndPlatformAndIsBackup :many
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE user_id = $1 AND platform = $2 AND is_backup = $3
`

type GetReposByUserIdAndPlatformAndIsBackupParams struct {
	UserID   uuid.UUID
	Platform string
	IsBackup bool
}

func (q *Queries) GetReposByUserIdAndPlatformAndIsBackup(ctx context.Context, arg GetReposByUserIdAndPlatformAndIsBackupParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, getReposByUserIdAndPlatformAndIsBackup, arg.UserID, arg.Platform, arg.IsBackup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repository
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.RepoID,
			&i.InstallationID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Platform,
			&i.Visibility,
			&i.IsRelease,
			&i.IsBackup,
			&i.IsApp,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReposByUserIdAndPlatformAndIsRelease = `-- name: GetReposByUserIdAndPlatformAndIsRelease :many
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE user_id = $1 AND platform = $2 AND is_release = $3
`

type GetReposByUserIdAndPlatformAndIsReleaseParams struct {
	UserID    uuid.UUID
	Platform  string
	IsRelease bool
}

func (q *Queries) GetReposByUserIdAndPlatformAndIsRelease(ctx context.Context, arg GetReposByUserIdAndPlatformAndIsReleaseParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, getReposByUserIdAndPlatformAndIsRelease, arg.UserID, arg.Platform, arg.IsRelease)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repository
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.RepoID,
			&i.InstallationID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Platform,
			&i.Visibility,
			&i.IsRelease,
			&i.IsBackup,
			&i.IsApp,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReposByUserIdAndVisibility = `-- name: GetReposByUserIdAndVisibility :many
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE user_id = $1 AND visibility = $2
`

type GetReposByUserIdAndVisibilityParams struct {
	UserID     uuid.UUID
	Visibility string
}

func (q *Queries) GetReposByUserIdAndVisibility(ctx context.Context, arg GetReposByUserIdAndVisibilityParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, getReposByUserIdAndVisibility, arg.UserID, arg.Visibility)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repository
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.RepoID,
			&i.InstallationID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Platform,
			&i.Visibility,
			&i.IsRelease,
			&i.IsBackup,
			&i.IsApp,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReposByUserIdAndVisibilityAndIsBackup = `-- name: GetReposByUserIdAndVisibilityAndIsBackup :many
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE user_id = $1 AND visibility = $2 AND is_backup = $3
`

type GetReposByUserIdAndVisibilityAndIsBackupParams struct {
	UserID     uuid.UUID
	Visibility string
	IsBackup   bool
}

func (q *Queries) GetReposByUserIdAndVisibilityAndIsBackup(ctx context.Context, arg GetReposByUserIdAndVisibilityAndIsBackupParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, getReposByUserIdAndVisibilityAndIsBackup, arg.UserID, arg.Visibility, arg.IsBackup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repository
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.RepoID,
			&i.InstallationID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Platform,
			&i.Visibility,
			&i.IsRelease,
			&i.IsBackup,
			&i.IsApp,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReposByUserIdAndVisibilityAndIsRelease = `-- name: GetReposByUserIdAndVisibilityAndIsRelease :many
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE user_id = $1 AND visibility = $2 AND is_release = $3
`

type GetReposByUserIdAndVisibilityAndIsReleaseParams struct {
	UserID     uuid.UUID
	Visibility string
	IsRelease  bool
}

func (q *Queries) GetReposByUserIdAndVisibilityAndIsRelease(ctx context.Context, arg GetReposByUserIdAndVisibilityAndIsReleaseParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, getReposByUserIdAndVisibilityAndIsRelease, arg.UserID, arg.Visibility, arg.IsRelease)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repository
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.RepoID,
			&i.InstallationID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Platform,
			&i.Visibility,
			&i.IsRelease,
			&i.IsBackup,
			&i.IsApp,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRepositoryByRepoId = `-- name: GetRepositoryByRepoId :one
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE repo_id = $1
`

func (q *Queries) GetRepositoryByRepoId(ctx context.Context, repoID uuid.UUID) (Repository, error) {
	row := q.db.QueryRowContext(ctx, getRepositoryByRepoId, repoID)
	var i Repository
	err := row.Scan(
		&i.RepoID,
		&i.InstallationID,
		&i.UserID,
		&i.Name,
		&i.Url,
		&i.Platform,
		&i.Visibility,
		&i.IsRelease,
		&i.IsBackup,
		&i.IsApp,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserReposByPage = `-- name: GetUserReposByPage :many
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE user_id = $1 LIMIT $2 OFFSET $3
`

type GetUserReposByPageParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

func (q *Queries) GetUserReposByPage(ctx context.Context, arg GetUserReposByPageParams) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, getUserReposByPage, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repository
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.RepoID,
			&i.InstallationID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Platform,
			&i.Visibility,
			&i.IsRelease,
			&i.IsBackup,
			&i.IsApp,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserRepository = `-- name: GetUserRepository :many
SELECT repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at FROM repository WHERE user_id = $1
`

func (q *Queries) GetUserRepository(ctx context.Context, userID uuid.UUID) ([]Repository, error) {
	rows, err := q.db.QueryContext(ctx, getUserRepository, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repository
	for rows.Next() {
		var i Repository
		if err := rows.Scan(
			&i.RepoID,
			&i.InstallationID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Platform,
			&i.Visibility,
			&i.IsRelease,
			&i.IsBackup,
			&i.IsApp,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRepoByRepoId = `-- name: UpdateRepoByRepoId :one
UPDATE repository SET
    name = $2,
    url = $3,
    platform = $4,
    visibility = $5,
    is_release = $6,
    is_backup = $7,
    description = $8,
    updated_at = CURRENT_TIMESTAMP
WHERE repo_id = $1
RETURNING repo_id, installation_id, user_id, name, url, platform, visibility, is_release, is_backup, is_app, description, created_at, updated_at
`

type UpdateRepoByRepoIdParams struct {
	RepoID      uuid.UUID
	Name        string
	Url         sql.NullString
	Platform    string
	Visibility  string
	IsRelease   bool
	IsBackup    bool
	Description sql.NullString
}

func (q *Queries) UpdateRepoByRepoId(ctx context.Context, arg UpdateRepoByRepoIdParams) (Repository, error) {
	row := q.db.QueryRowContext(ctx, updateRepoByRepoId,
		arg.RepoID,
		arg.Name,
		arg.Url,
		arg.Platform,
		arg.Visibility,
		arg.IsRelease,
		arg.IsBackup,
		arg.Description,
	)
	var i Repository
	err := row.Scan(
		&i.RepoID,
		&i.InstallationID,
		&i.UserID,
		&i.Name,
		&i.Url,
		&i.Platform,
		&i.Visibility,
		&i.IsRelease,
		&i.IsBackup,
		&i.IsApp,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
