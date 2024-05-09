-- +goose Up
-- +goose StatementBegin

SET search_path TO ozon, public;

TRUNCATE TABLE default_user CASCADE;
ALTER TABLE default_user DROP COLUMN user_login;
ALTER TABLE default_user ADD COLUMN email text NOT NULL UNIQUE;

ALTER TABLE user_profile DROP COLUMN email;
ALTER TABLE user_profile DROP COLUMN phone_number;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

SET search_path TO ozon, public;

TRUNCATE TABLE default_user CASCADE;
ALTER TABLE default_user DROP COLUMN email;
ALTER TABLE default_user ADD COLUMN user_login text NOT NULL CHECK(char_length(user_login) BETWEEN 4 AND 32);

ALTER TABLE user_profile ADD COLUMN email text NOT NULL DEFAULT '';
ALTER TABLE user_profile ADD COLUMN phone_number TEXT NOT NULL DEFAULT '';

-- +goose StatementEnd
