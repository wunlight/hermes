-- name: CreateBrand :one
INSERT INTO brands (
    code,
    name
)
VALUES ($1, $2)
RETURNING
    id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at;



-- name: GetBrandByID :one
SELECT
    id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at
FROM brands
WHERE id = $1
  AND deleted_at IS NULL;



-- name: GetBrandByCode :one
SELECT
    id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at
FROM brands
WHERE code = $1
  AND deleted_at IS NULL;



-- name: UpdateBrand :one
UPDATE brands
SET
    code = $2,
    name = $3,
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING
    id,
    code,
    name,
    created_at,
    updated_at,
    deleted_at;



-- name: DeleteBrand :exec
UPDATE brands
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL;



-- name: ListBrands :many
SELECT *
FROM brands
WHERE deleted_at IS NULL
ORDER BY name ASC;


