-- +goose Up
-- +goose StatementBegin

SET search_path TO ozon, public;

ALTER TABLE user_profile DROP CONSTRAINT user_profile_email_check;
ALTER TABLE user_profile ALTER COLUMN email SET DEFAULT '';

ALTER TABLE user_profile DROP CONSTRAINT user_profile_phone_number_check;
ALTER TABLE user_profile ALTER COLUMN phone_number SET DEFAULT '';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
