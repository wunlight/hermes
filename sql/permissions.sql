-- name: CreatePermission :one
INSERT INTO permissions (
    code,
    description
)
VALUES (
    $1,
    $2
)
RETURNING
    id,
    code,
    description,
    created_at,
    updated_at;

-- name: GetPermissionByCode :one
SELECT
    id,
    code,
    description,
    created_at,
    updated_at
FROM permissions
WHERE code = $1;

-- name: GetPermissionByCodes :many
SELECT
    id,
    code,
    description,
    created_at,
    updated_at
FROM permissions
WHERE code = ANY(sqlc.arg(codes)::text[]);