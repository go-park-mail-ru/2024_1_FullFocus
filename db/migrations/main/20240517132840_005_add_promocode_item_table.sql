-- +goose Up
-- +goose StatementBegin

SET search_path TO ozon, public;

CREATE TABLE IF NOT EXISTS promocode_item (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    promocode_type BIGINT NOT NULL REFERENCES promocode(id) ON DELETE CASCADE ON UPDATE CASCADE,
    profile_id BIGINT NOT NULL REFERENCES user_profile(id) ON DELETE CASCADE ON UPDATE CASCADE,
    code TEXT NOT NULL UNIQUE CHECK (char_length(code) BETWEEN 4 AND 8),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS promocode_item;

-- +goose StatementEnd
