-- name: CreateRole :one
INSERT INTO roles (
    code,
    name,
    description
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING
    id,
    code,
    name,
    description,
    created_at,
    updated_at;