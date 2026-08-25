-- name: CreateUnit :one
INSERT INTO units (
    code,
    name
)
VALUES (
    $1,
    $2
)
RETURNING
    id,
    code,
    name;



-- name: GetUnitByID :one
SELECT
    id,
    code,
    name
FROM units
WHERE id = $1;



-- name: GetUnitByCode :one
SELECT
    id,
    code,
    name
FROM units
WHERE code = $1;



-- name: UpdateUnit :one
UPDATE units
SET
    code = $2,
    name = $3
WHERE id = $1
RETURNING
    id,
    code,
    name;



-- name: DeleteUnit :exec
DELETE FROM units
WHERE id = $1;



-- name: ListUnits :many
SELECT
    id,
    code,
    name
FROM units
ORDER BY name ASC;


