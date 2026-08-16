-- name: CreateUser :one
INSERT INTO users (
    email,
    password_hash,
    name
)
VALUES ($1, $2, $3)
RETURNING
    id,
    email,
    password_hash,
    name,
    created_at,
    updated_at,
    deleted_at;



-- name: GetUserByID :one
SELECT
    id,
    email,
    password_hash,
    name,
    created_at,
    updated_at,
    deleted_at
FROM users
WHERE id = $1
  AND deleted_at IS NULL
LIMIT 1;



-- name: GetUserByEmail :one
SELECT
    id,
    email,
    password_hash,
    name,
    created_at,
    updated_at,
    deleted_at
FROM users
WHERE email = $1
  AND deleted_at IS NULL
LIMIT 1;



-- name: UpdateUser :one
UPDATE users
SET
    email = $2,
    name = $3,
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING
    id,
    email,
    password_hash,
    name,
    created_at,
    updated_at,
    deleted_at;



-- name: DeleteUser :exec
UPDATE users
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL;



-- name: ListUsers :many
SELECT
    id,
    email,
    password_hash,
    name,
    created_at,
    updated_at,
    deleted_at
FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC;


