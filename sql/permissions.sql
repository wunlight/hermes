-- name: CreatePermission :one
INSERT INTO permissions (
    code,
    description
)
VALUES (
    $1,
    $2
)
RETURNING *;