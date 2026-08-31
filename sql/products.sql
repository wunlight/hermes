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
    description
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
    p.id,
    p.sku,
    p.name,
    c.id AS category_id,
    c.name AS category_name,
    b.id AS brand_id,
    b.name AS brand_name,
    u.id AS unit_id,
    u.name AS unit_name,
    p.min_stock,
    p.weight,
    p.length,
    p.width,
    p.description,
    p.status,
    p.created_at,
    p.updated_at,
    p.deleted_at
FROM products p
LEFT JOIN categories c ON c.id = p.category_id
LEFT JOIN brands b     ON b.id = p.brand_id
LEFT JOIN units u      ON u.id = p.unit_id
WHERE p.id = $1
  AND p.deleted_at IS NULL;



-- name: GetProductBySKU :one
SELECT
    p.id,
    p.sku,
    p.name,
    c.id AS category_id,
    c.name AS category_name,
    b.id AS brand_id,
    b.name AS brand_name,
    u.id AS unit_id,
    u.name AS unit_name,
    p.min_stock,
    p.weight,
    p.length,
    p.width,
    p.description,
    p.status,
    p.created_at,
    p.updated_at,
    p.deleted_at
FROM products p
LEFT JOIN categories c ON c.id = p.category_id
LEFT JOIN brands b     ON b.id = p.brand_id
LEFT JOIN units u      ON u.id = p.unit_id
WHERE p.sku = $1
  AND p.deleted_at IS NULL;



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



-- name: DeleteProduct :execresult
UPDATE products
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL;



-- name: ListProducts :many
SELECT
    p.id,
    p.sku,
    p.name,
    c.id AS category_id,
    c.name AS category_name,
    b.id AS brand_id,
    b.name AS brand_name,
    u.id AS unit_id,
    u.name AS unit_name,
    p.min_stock,
    p.weight,
    p.length,
    p.width,
    p.description,
    p.status,
    p.created_at,
    p.updated_at,
    p.deleted_at
FROM products p
LEFT JOIN categories c ON c.id = p.category_id
LEFT JOIN brands b     ON b.id = p.brand_id
LEFT JOIN units u      ON u.id = p.unit_id
WHERE p.deleted_at IS NULL
ORDER BY p.name ASC;


