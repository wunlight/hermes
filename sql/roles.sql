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

-- name: CreateRolePermission :exec
INSERT INTO role_permissions (
    role_id,
    permission_id
)
VALUES (
    $1,
    $2
);