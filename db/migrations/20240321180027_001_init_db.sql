-- +goose Up
-- +goose StatementBegin

-- ozon schema definition

CREATE SCHEMA ozon;
SET search_path TO ozon, public;

-- ozon.user definition

CREATE TABLE default_user (
	id SERIAL PRIMARY KEY,
	user_login TEXT NOT NULL UNIQUE CHECK (char_length(user_login) BETWEEN 4 AND 32),
	password_hash TEXT NOT NULL CHECK (char_length(user_login) BETWEEN 8 AND 255),
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NOT NULL
);

-- ozon.profile definition

CREATE TABLE user_profile (
	id int4 PRIMARY KEY REFERENCES default_user(id) ON DELETE CASCADE ON UPDATE CASCADE,
	full_name TEXT NOT NULL CHECK (char_length(full_name) BETWEEN 1 AND 255),
	email TEXT NOT NULL CHECK ((char_length(email) BETWEEN 4 AND 255) AND (email ~* '^[a-z0-9\.\-]+@[a-z0-9\.\-]+\.[a-z]+$')),
	phone_number TEXT NOT NULL CHECK ((char_length(phone_number) BETWEEN 5 AND 15) AND (phone_number ~ '\+?[0-9]+')),
	imgsrc TEXT DEFAULT '',
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NOT NULL
);

-- ozon.product definition

CREATE TABLE product (
	id SERIAL PRIMARY KEY,
	product_name TEXT NOT NULL CHECK (char_length(product_name) BETWEEN 1 AND 50),
	product_description TEXT DEFAULT NULL CHECK (char_length(product_description) BETWEEN 1 AND 255),
	price int4 NOT NULL CHECK (price > 0),
	imgsrc TEXT DEFAULT '',
	seller TEXT NOT NULL CHECK (char_length(seller) > 0),
	rating int2 DEFAULT 0 NOT NULL CHECK (rating >= 0),
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NOT NULL
);

-- ozon.category definition

CREATE TABLE category (
	id SMALLSERIAL PRIMARY KEY,
	category_name TEXT NOT NULL UNIQUE CHECK (char_length(category_name) BETWEEN 1 AND 50),
	parent_id int2 NOT NULL REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- ozon.product_category definition

CREATE TABLE product_category (
	product_id int4 NOT NULL REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
	category_id int2 NOT NULL REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE,
	product_category_pk SMALLSERIAL NOT NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NOT NULL,
	CONSTRAINT product_category_pk PRIMARY KEY (product_id, category_id)
);

-- ozon.ordering definition

CREATE TYPE ordering_status AS ENUM (
    'created',
    'cancelled',
    'ready'
    );

CREATE TABLE ordering (
	id bigserial PRIMARY KEY,
	sum int4 DEFAULT 0 NOT NULL CHECK (sum >= 0),
	profile_id int4 NOT NULL REFERENCES user_profile(id) ON DELETE CASCADE ON UPDATE CASCADE ,
	order_status ordering_status NOT NULL,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NOT NULL
);

-- ozon.order_item definition

CREATE TABLE order_item (
	ordering_id int8 NOT NULL REFERENCES ordering(id) ON DELETE CASCADE ON UPDATE CASCADE,
	product_id int4 NOT NULL REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
	count int2 DEFAULT 1 NOT NULL CHECK (count > 0),
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NOT NULL,
	CONSTRAINT order_item_pk PRIMARY KEY (ordering_id, product_id)
);

-- ozon.cart_item definition

CREATE TABLE cart_item (
	product_id int4 NOT NULL REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
	profile_id int4 NOT NULL REFERENCES user_profile(id) ON DELETE CASCADE ON UPDATE CASCADE,
	count int2 DEFAULT 1 NOT NULL CHECK (count >= 0),
	CONSTRAINT cart_item_pk PRIMARY KEY (profile_id, product_id),
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP DEFAULT now() NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS cart_item;
DROP TABLE IF EXISTS order_item;
DROP TABLE IF EXISTS ordering;
DROP TYPE IF EXISTS ordering_status;
DROP TABLE IF EXISTS user_profile;
DROP TABLE IF EXISTS default_user;
DROP TABLE IF EXISTS product_category;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS category;

DROP SCHEMA IF EXISTS ozon CASCADE;
SET search_path TO public;

-- +goose StatementEnd