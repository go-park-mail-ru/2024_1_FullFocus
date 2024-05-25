-- +goose Up
-- +goose StatementBegin

CREATE TYPE notification_type AS ENUM (
    'order_status_change'
);

CREATE TABLE IF NOT EXISTS notification (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    profile_id BIGINT NOT NULL REFERENCES user_profile(id) ON DELETE CASCADE ON UPDATE CASCADE,
    type notification_type NOT NULL,
    read_status BOOLEAN DEFAULT FALSE,
    payload JSONB NOT NULL CHECK (length(payload::TEXT) > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TYPE IF EXISTS notification_type;
DROP TABLE IF EXISTS notification;

-- +goose StatementEnd
