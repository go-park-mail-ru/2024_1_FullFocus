-- +goose Up
-- +goose StatementBegin

-- ozon schema definition

CREATE SCHEMA ozon;
SET search_path TO ozon, public;

-- ozon.user definition

CREATE TABLE default_user (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_login TEXT NOT NULL UNIQUE CHECK (char_length(user_login) BETWEEN 4 AND 32),
    password_hash TEXT NOT NULL CHECK (char_length(password_hash) BETWEEN 8 AND 255),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

-- ozon.profile definition

CREATE TABLE user_profile (
    id INT NOT NULL UNIQUE REFERENCES default_user(id) ON DELETE CASCADE ON UPDATE CASCADE,
    full_name TEXT NOT NULL CHECK (char_length(full_name) BETWEEN 1 AND 255),
    email TEXT NOT NULL CHECK (char_length(email) BETWEEN 4 AND 255),
    phone_number TEXT NOT NULL CHECK ((char_length(phone_number) BETWEEN 5 AND 15) AND (phone_number ~ '\+?[0-9]+')),
    imgsrc TEXT DEFAULT '',
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

-- ozon.product definition

CREATE TABLE product (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_name TEXT NOT NULL CHECK (char_length(product_name) BETWEEN 1 AND 50),
    product_description TEXT DEFAULT NULL CHECK (char_length(product_description) BETWEEN 1 AND 255),
    price INT NOT NULL CHECK (price > 0),
    imgsrc TEXT DEFAULT '',
    seller TEXT NOT NULL CHECK (char_length(seller) > 0),
    rating SMALLINT DEFAULT 0 NOT NULL CHECK (rating >= 0),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

-- ozon.category definition

CREATE TABLE category (
    id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_name TEXT NOT NULL UNIQUE CHECK (char_length(category_name) BETWEEN 1 AND 50),
    parent_id SMALLINT DEFAULT NULL REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- ozon.product_category definition

CREATE TABLE product_category (
    product_id INT NOT NULL REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
    category_id SMALLINT NOT NULL REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    CONSTRAINT product_category_pk PRIMARY KEY (product_id, category_id)
);

-- ozon.order definition

CREATE TABLE "order" (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    sum INT DEFAULT 0 NOT NULL CHECK (sum >= 0),
    profile_id INT NOT NULL REFERENCES user_profile(id) ON DELETE CASCADE ON UPDATE CASCADE ,
    order_status TEXT NOT NULL CHECK (order_status IN ('created', 'cancelled', 'ready')),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

-- ozon.order_item definition

CREATE TABLE order_item (
    order_id BIGINT NOT NULL REFERENCES "order"(id) ON DELETE CASCADE ON UPDATE CASCADE,
    product_id INT NOT NULL REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
    count SMALLINT DEFAULT 1 NOT NULL CHECK (count > 0),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    CONSTRAINT order_item_pk PRIMARY KEY (order_id, product_id)
);

-- ozon.cart_item definition

CREATE TABLE cart_item (
    product_id INT NOT NULL REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
    profile_id INT NOT NULL REFERENCES user_profile(id) ON DELETE CASCADE ON UPDATE CASCADE,
    count SMALLINT DEFAULT 1 NOT NULL CHECK (count >= 0),
    CONSTRAINT cart_item_pk PRIMARY KEY (profile_id, product_id),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS cart_item;
DROP TABLE IF EXISTS order_item;
DROP TABLE IF EXISTS "order";
DROP TABLE IF EXISTS user_profile;
DROP TABLE IF EXISTS default_user;
DROP TABLE IF EXISTS product_category;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS category;

DROP SCHEMA IF EXISTS ozon CASCADE;
SET search_path TO public;

-- +goose StatementEnd