-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS pg_stat_statements;
CREATE EXTENSION IF NOT EXISTS auto_explain;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP EXTENSION IF EXISTS pg_stat_statements;
DROP EXTENSION IF EXISTS auto_explain;

-- +goose StatementEnd
