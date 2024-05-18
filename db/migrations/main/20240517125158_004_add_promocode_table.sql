-- +goose Up
-- +goose StatementBegin

SET search_path TO ozon, public;

CREATE TYPE promo_benefit_type AS ENUM (
    'percentage',
    'price discount',
    'free delivery'
);

CREATE TABLE IF NOT EXISTS promocode (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE CHECK (char_length(name) > 4),
    description TEXT NOT NULL DEFAULT '',
    min_sum_give INTEGER NOT NULL DEFAULT 0,
    min_sum_activation INTEGER NOT NULL DEFAULT 0,
    benefit_type promo_benefit_type NOT NULL,
    value INTEGER NOT NULL DEFAULT 0,
    ttl_hours SMALLINT NOT NULL CHECK (ttl_hours > 0)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS promocode;
DROP TYPE IF EXISTS promo_benefit_type;

-- +goose StatementEnd
