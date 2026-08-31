-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sku TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,

    category_id UUID NOT NULL,
    brand_id UUID,
    unit_id UUID NOT NULL,

    min_stock INTEGER NOT NULL DEFAULT 0,

    weight NUMERIC(12, 3),
    length NUMERIC(12, 3),
    width NUMERIC(12, 3),

    description TEXT,

    status TEXT NOT NULL DEFAULT 'active',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_products_category
        FOREIGN KEY (category_id)
        REFERENCES categories(id)
        ON DELETE RESTRICT,

    CONSTRAINT fk_products_brand
        FOREIGN KEY (brand_id)
        REFERENCES brands(id)
        ON DELETE RESTRICT,

    CONSTRAINT fk_products_unit
        FOREIGN KEY (unit_id)
        REFERENCES units(id)
        ON DELETE RESTRICT,

    CONSTRAINT chk_products_min_stock
        CHECK (min_stock >= 0),

    CONSTRAINT chk_products_weight
        CHECK (weight IS NULL OR weight >= 0),

    CONSTRAINT chk_products_length
        CHECK (length IS NULL OR length >= 0),

    CONSTRAINT chk_products_width
        CHECK (width IS NULL OR width >= 0),

    CONSTRAINT chk_products_status
        CHECK (status IN ('active', 'inactive'))
);

CREATE INDEX IF NOT EXISTS idx_products_category_id
    ON products(category_id);

CREATE INDEX IF NOT EXISTS idx_products_brand_id
    ON products(brand_id);

CREATE INDEX IF NOT EXISTS idx_products_unit_id
    ON products(unit_id);

CREATE INDEX IF NOT EXISTS idx_products_status
    ON products(status);

CREATE INDEX IF NOT EXISTS idx_products_deleted_at
    ON products(deleted_at);

-- +goose Down
DROP TABLE IF EXISTS products;
