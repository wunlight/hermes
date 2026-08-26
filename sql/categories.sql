-- name: CreateCategory :one
INSERT INTO categories (
    parent_id,
    code,
    name
)
VALUES ($1, $2, $3)
RETURNING
    id,
    parent_id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at;



-- name: GetCategoryByID :one
SELECT
    id,
    parent_id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at
FROM categories
WHERE id = $1
  AND deleted_at IS NULL;



-- name: GetCategoryByCode :one
SELECT
    id,
    parent_id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at
FROM categories
WHERE code = $1
  AND deleted_at IS NULL;



-- name: UpdateCategory :one
UPDATE categories
SET
    parent_id = $2,
    code = $3,
    name = $4,
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING
    id,
    parent_id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at;



-- name: DeleteCategory :execresult
UPDATE categories
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL;



-- name: ListCategories :many
SELECT
    id,
    parent_id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at
FROM categories
WHERE deleted_at IS NULL
ORDER BY name ASC;



-- name: ListRootCategories :many
SELECT
    id,
    parent_id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at
FROM categories
WHERE parent_id IS NULL
  AND deleted_at IS NULL
ORDER BY name ASC;



-- name: ListChildCategories :many
SELECT
    id,
    parent_id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at
FROM categories
WHERE parent_id = $1
  AND deleted_at IS NULL
ORDER BY name ASC;


