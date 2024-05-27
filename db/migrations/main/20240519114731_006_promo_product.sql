-- +goose Up
-- +goose StatementBegin

SET search_path TO ozon, public;

CREATE TYPE promotion_product_sale_type AS ENUM (
    'percentSale',
    'priceSale',
    'finalPrice'
);

CREATE TABLE promo_product (
    product_id INT PRIMARY KEY REFERENCES product(id),
    benefit_type promotion_product_sale_type NOT NULL,
    benefit_value INT NOT NULL CHECK(benefit_value > 0)
);

ALTER TABLE product
ADD COLUMN on_sale BOOLEAN DEFAULT FALSE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

SET search_path TO ozon, public;

DROP TABLE promo_product;
DROP TYPE promotion_product_sale_type;

ALTER TABLE product
DROP COLUMN on_sale;

-- +goose StatementEnd
