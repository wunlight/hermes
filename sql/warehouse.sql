-- name: CreateWarehouse :one
INSERT INTO warehouses (
    code,
    name,
    description
) VALUES (
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
    updated_at,
    deleted_at;



-- name: GetWarehouseByID :one
SELECT
    id,
    code,
    name,
    description,
    created_at,
    updated_at,
    deleted_at
FROM warehouses
WHERE id = $1
  AND deleted_at IS NULL;



-- name: GetWarehouseByCode :one
SELECT
    id,
    code,
    name,
    description,
    created_at,
    updated_at,
    deleted_at
FROM warehouses
WHERE code = $1
  AND deleted_at IS NULL;



-- name: UpdateWarehouse :one
UPDATE warehouses
SET
    code = $2,
    name = $3,
    description = $4,
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING
    id,
    code,
    name,
    description,
    created_at,
    updated_at,
    deleted_at;



-- name: DeleteWarehouse :execresult
UPDATE warehouses
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL;



-- name: ListWarehouses :many
SELECT
    id,
    code,
    name,
    description,
    created_at,
    updated_at,
    deleted_at
FROM warehouses
WHERE deleted_at IS NULL
ORDER BY created_at DESC;


