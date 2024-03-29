-- name: CreateRepository :one
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
) RETURNING *;

-- name: GetRepositoryByRepoId :one
SELECT * FROM repository WHERE repo_id = $1;

-- name: GetRepoByURl :one
SELECT * FROM repository WHERE url = $1;

-- name: GetUserRepository :many
SELECT * FROM repository WHERE user_id = $1;

-- name: GetUserReposByPage :many
SELECT * FROM repository WHERE user_id = $1 LIMIT $2 OFFSET $3;

-- get repo pairs - (userid, (name), (visibility), (visibility, is_backup) (visibility, is_release), (platform), (platform, is_backup), (platform, is_release))

-- name: GetReposByUserId :many
SELECT * FROM repository WHERE user_id = $1;

-- name: GetReposByUserIdAndVisibility :many
SELECT * FROM repository WHERE user_id = $1 AND visibility = $2;

-- name: GetReposByUserIdAndPlatform :many
SELECT * FROM repository WHERE user_id = $1 AND platform = $2;

-- name: GetReposByUserIdAndVisibilityAndIsBackup :many
SELECT * FROM repository WHERE user_id = $1 AND visibility = $2 AND is_backup = $3;

-- name: GetReposByUserIdAndVisibilityAndIsRelease :many
SELECT * FROM repository WHERE user_id = $1 AND visibility = $2 AND is_release = $3;

-- name: GetReposByUserIdAndPlatformAndIsBackup :many
SELECT * FROM repository WHERE user_id = $1 AND platform = $2 AND is_backup = $3;

-- name: GetReposByUserIdAndPlatformAndIsRelease :many
SELECT * FROM repository WHERE user_id = $1 AND platform = $2 AND is_release = $3;

-- name: GetReposByUserIdAndName :many
SELECT * FROM repository WHERE user_id = $1 AND name = $2;

-- name: UpdateRepoByRepoId :one
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
RETURNING *;

-- name: DeleteRepoByRepoId :one
DELETE FROM repository WHERE repo_id = $1 RETURNING *;