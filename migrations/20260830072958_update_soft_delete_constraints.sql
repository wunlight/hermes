-- +goose Up
ALTER TABLE categories
DROP CONSTRAINT categories_code_key;

CREATE UNIQUE INDEX idx_categories_code_active
    ON categories(code)
    WHERE deleted_at IS NULL;

ALTER TABLE brands
DROP CONSTRAINT brands_code_key;

CREATE UNIQUE INDEX idx_brands_code_active
    ON brands(code)
    WHERE deleted_at IS NULL;

DROP INDEX IF EXISTS idx_products_deleted_at;

-- +goose Down
CREATE INDEX idx_products_deleted_at
    ON products(deleted_at);

DROP INDEX IF EXISTS idx_brands_code_active;

ALTER TABLE brands
ADD CONSTRAINT brands_code_key UNIQUE (code);

DROP INDEX IF EXISTS idx_categories_code_active;

ALTER TABLE categories
ADD CONSTRAINT categories_code_key UNIQUE (code);
