-- +goose Up

-- +goose StatementBegin

CREATE TABLE poll (
  id int4 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1 NO CYCLE) NOT NULL,
  title text NOT NULL,
  CONSTRAINT poll_pkey PRIMARY KEY (id)
);

CREATE TABLE response (
  profile_id int4 NOT NULL,
  poll_id int4 NOT NULL,
  rate int4 NOT NULL,
  CONSTRAINT response_pk PRIMARY KEY (profile_id, poll_id),
  CONSTRAINT response_rate_check CHECK (((rate >= 1) AND (rate <= 10))),
  CONSTRAINT response_poll_id_fkey FOREIGN KEY (poll_id) REFERENCES poll(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS response;
DROP TABLE IF EXISTS poll;

-- +goose StatementEnd
