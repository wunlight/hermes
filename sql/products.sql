-- name: CreateProduct :one
INSERT INTO products (
    sku,
    name,
    category_id,
    brand_id,
    unit_id,
    min_stock,
    weight,
    length,
    width,
    description,
    status
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING
    id,
    sku,
    name,
    category_id,
    brand_id,
    unit_id,
    min_stock,
    weight,
    length,
    width,
    description,
    status,
    created_at,
    updated_at,
    deleted_at;



-- name: GetProductByID :one
SELECT
    id,
    sku,
    name,
    category_id,
    brand_id,
    unit_id,
    min_stock,
    weight,
    length,
    width,
    description,
    status,
    created_at,
    updated_at,
    deleted_at
FROM products
WHERE id = $1
  AND deleted_at IS NULL;



-- name: GetProductBySKU :one
SELECT
    id,
    sku,
    name,
    category_id,
    brand_id,
    unit_id,
    min_stock,
    weight,
    length,
    width,
    description,
    status,
    created_at,
    updated_at,
    deleted_at
FROM products
WHERE sku = $1
  AND deleted_at IS NULL;



-- name: UpdateProduct :one
UPDATE products
SET
    sku = $2,
    name = $3,
    category_id = $4,
    brand_id = $5,
    unit_id = $6,
    min_stock = $7,
    weight = $8,
    length = $9,
    width = $10,
    description = $11,
    status = $12,
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING
    id,
    sku,
    name,
    category_id,
    brand_id,
    unit_id,
    min_stock,
    weight,
    length,
    width,
    description,
    status,
    created_at,
    updated_at,
    deleted_at;



-- name: DeleteProduct :exec
UPDATE products
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL;



-- name: ListProducts :many
SELECT
    id,
    sku,
    name,
    category_id,
    brand_id,
    unit_id,
    min_stock,
    weight,
    length,
    width,
    description,
    status,
    created_at,
    updated_at,
    deleted_at
FROM products
WHERE deleted_at IS NULL
ORDER BY name ASC;



-- name: ListProductsByCategory :many
SELECT
    id,
    sku,
    name,
    category_id,
    brand_id,
    unit_id,
    min_stock,
    weight,
    length,
    width,
    description,
    status,
    created_at,
    updated_at,
    deleted_at
FROM products
WHERE category_id = $1
  AND deleted_at IS NULL
ORDER BY name ASC;



-- name: ListProductsByBrand :many
SELECT
    id,
    sku,
    name,
    category_id,
    brand_id,
    unit_id,
    min_stock,
    weight,
    length,
    width,
    description,
    status,
    created_at,
    updated_at,
    deleted_at
FROM products
WHERE brand_id = $1
  AND deleted_at IS NULL
ORDER BY name ASC;



-- name: ListProductsByStatus :many
SELECT
    id,
    sku,
    name,
    category_id,
    brand_id,
    unit_id,
    min_stock,
    weight,
    length,
    width,
    description,
    status,
    created_at,
    updated_at,
    deleted_at
FROM products
WHERE status = $1
  AND deleted_at IS NULL
ORDER BY name ASC;


