-- name: GetAllUsers :many
SELECT * FROM users WHERE deleted_at IS NULL;

-- name: CountUsers :one
SELECT COUNT(*) FROM users WHERE deleted_at IS NULL;

-- name: ListUsers :many
SELECT * FROM users  WHERE deleted_at IS NULL GROUP BY 1 offset $1 LIMIT $2;

-- name: CreateUser :one
INSERT INTO users (username, password, avatar, email) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL;

-- name: GetUserByName :one
SELECT * FROM users WHERE username = $1 AND deleted_at IS NULL;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1 AND deleted_at IS NULL;

-- name: UpdateUserById :one
UPDATE users SET username = $1, password = $2, avatar = $3, email = $4, updated_at = now() WHERE id = $5 AND deleted_at IS NULL RETURNING *;