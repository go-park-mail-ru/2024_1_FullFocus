-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA csat;
SET search_path TO csat, public;

CREATE TABLE poll (
    id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title TEXT NOT NULL
);

CREATE TABLE response (
    profile_id BIGINT NOT NULL,
    poll_id SMALLINT NOT NULL,
    rate SMALLINT NOT NULL,
    CONSTRAINT response_pk PRIMARY KEY (profile_id, poll_id),
    CONSTRAINT response_rate_check CHECK (((rate >= 1) AND (rate <= 10))),
    CONSTRAINT response_poll_id_fkey FOREIGN KEY (poll_id) REFERENCES poll(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS response;
DROP TABLE IF EXISTS poll;

DROP SCHEMA IF EXISTS csat CASCADE;
SET search_path TO public;

-- +goose StatementEnd
