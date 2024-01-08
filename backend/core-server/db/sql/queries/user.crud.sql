-- name: CreateUser :one
INSERT INTO users (
    user_id,
    local_username,
    local_password,
    oauth_id,
    oauth_provider,
    oauth_name,
    email,
    profile_picture
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING *;

-- name: GetUserByLocalUsername :one
SELECT * FROM users WHERE local_username = $1;

-- name: GetUserByOauthId :one
SELECT * FROM users WHERE oauth_id = $1;

-- name: GetUserByUserId :one
SELECT * FROM users WHERE user_id = $1;

-- name: UpdateUserPassword :one
UPDATE users SET local_password = $2 WHERE user_id = $1 RETURNING *;

-- name: UpdateUserWalletAddress :one
UPDATE users SET wallet_address = $2 WHERE user_id = $1 RETURNING *;

-- name: UpdateUserSecurityDetails :one
UPDATE users SET rsa_public_key = $2 AND hashed_secret = $3 WHERE user_id = $1 RETURNING *;

-- name: UpdateUserEmail :one
UPDATE users SET email = $2 WHERE user_id = $1 RETURNING *;

-- name: UpdateUser :one
UPDATE users SET
    local_username = $2,
    local_password = $3,
    oauth_id = $4,
    oauth_provider = $5,
    oauth_name = $6,
    email = $7,
    profile_picture = $8,
    wallet_address = $9,
    rsa_public_key = $10,
    hashed_secret = $11
WHERE user_id = $1 RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE user_id = $1 RETURNING *;