-- +goose Up
-- +goose StatementBegin

-- ozon schema definition

CREATE SCHEMA ozon;
SET search_path TO ozon, public;

-- ozon.user definition

CREATE TABLE default_user (
	id uuid PRIMARY KEY,
	user_login text NOT NULL UNIQUE,
	password_hash text NOT NULL,
	created_at timetz DEFAULT now() NOT NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT user_check CHECK (char_length(user_login) BETWEEN 4 AND 32)
);

-- ozon.profile definition

CREATE TABLE user_profile (
	id uuid PRIMARY KEY,
	full_name text NOT NULL,
	email text NOT NULL,
	imgsrc text DEFAULT 'default-avatar.png' NOT NULL,
	phone_number text NOT NULL,
	user_id uuid NOT NULL REFERENCES default_user(id) ON DELETE CASCADE ON UPDATE CASCADE,
	created_at timetz DEFAULT now() NOT NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT email_lenght CHECK (char_length(email) BETWEEN 4 AND 255),
	CONSTRAINT email_valid CHECK (email ~* '^[a-z0-9\.\-]+@[a-z0-9\.\-]+\.[a-z]+$'),
	CONSTRAINT name_length CHECK (char_length(full_name) BETWEEN 5 AND 255),
	CONSTRAINT phone_length CHECK (char_length(phone_number) BETWEEN 5 AND 15),
	CONSTRAINT phone_valid CHECK (phone_number ~ '\+?[0-9]+')
);

-- ozon.product definition

CREATE TABLE product (
	id uuid PRIMARY KEY,
	product_name text NOT NULL,
	product_description text NULL,
	price numeric NOT NULL,
	imgsrc text DEFAULT 'default-product.png' NOT NULL,
	seller text NOT NULL,
	rating int4 DEFAULT 0 NOT NULL,
	category_id int4 NOT NULL,
	created_at timetz DEFAULT now() NOT NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT descriprion_length CHECK (char_length(product_description) BETWEEN 1 AND 255),
	CONSTRAINT name_length CHECK (char_length(product_name) BETWEEN 1 AND 50),
	CONSTRAINT price_positive CHECK (price > 0)
);

-- ozon.category definition

CREATE TABLE category (
	id smallserial PRIMARY KEY,
	category_name text NOT NULL UNIQUE,
	parent_id int2 NOT NULL REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT name_length CHECK (char_length(category_name) BETWEEN 1 AND 50)
);

-- ozon.product_category definition

CREATE TABLE product_category (
	product_id uuid NOT NULL REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
	category_id int2 NOT NULL REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE,
	product_category_pk int8 NOT NULL,
	created_at timetz DEFAULT now() NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT product_category_pk PRIMARY KEY (product_id, category_id)
);

-- ozon.order_item definition

CREATE TABLE order_item (
	id uuid PRIMARY KEY,
	product_id uuid NOT NULL REFERENCES product(id),
	count int2 DEFAULT 1 NOT NULL,
	created_at timetz DEFAULT now() NOT NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT count_positive CHECK (count >= 0)
);

-- ozon.ordering definition

CREATE TYPE ordering_status AS ENUM (
	'created',
	'cancelled',
	'ready'
);

CREATE TABLE ordering (
	id uuid PRIMARY KEY,
	sum int4 DEFAULT 0 NOT NULL,
	order_item uuid NOT NULL REFERENCES order_item(id),
	profile_id uuid NOT NULL REFERENCES user_profile(id),
	order_status ordering_status NOT NULL,
	created_at timetz DEFAULT now() NOT NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT sum_positive CHECK (sum >= 0)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS ordering;
DROP TYPE IF EXISTS ordering_status;
DROP TABLE IF EXISTS order_item;
DROP TABLE IF EXISTS user_profile;
DROP TABLE IF EXISTS default_user;
DROP TABLE IF EXISTS product_category;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS category;

DROP SCHEMA IF EXISTS ozon CASCADE;
SET search_path TO public;

-- +goose StatementEnd
