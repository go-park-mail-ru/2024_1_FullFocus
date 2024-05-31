CREATE SCHEMA ozon;

CREATE TABLE default_user (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_login TEXT NOT NULL UNIQUE CHECK (char_length(user_login) BETWEEN 4 AND 32),
    password_hash TEXT NOT NULL CHECK (char_length(password_hash) BETWEEN 8 AND 255),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE user_profile (
    id BIGINT NOT NULL UNIQUE REFERENCES default_user(id) ON DELETE CASCADE ON UPDATE CASCADE,
    full_name TEXT NOT NULL CHECK (char_length(full_name) BETWEEN 1 AND 255),
    email TEXT NOT NULL UNIQUE CHECK (char_length(email) BETWEEN 4 AND 255),
    phone_number TEXT NOT NULL UNIQUE CHECK (char_length(phone_number) BETWEEN 5 AND 15),
    imgsrc TEXT DEFAULT '',
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE product (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_name TEXT NOT NULL CHECK (char_length(product_name) BETWEEN 1 AND 50),
    product_description TEXT DEFAULT NULL CHECK (char_length(product_description) BETWEEN 1 AND 255),
    price INT NOT NULL CHECK (price > 0),
    imgsrc TEXT DEFAULT '',
    seller TEXT NOT NULL CHECK (char_length(seller) > 0),
    rating SMALLINT DEFAULT 0 NOT NULL CHECK (rating >= 0),
    on_sale BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE category (
    id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_name TEXT NOT NULL UNIQUE CHECK (char_length(category_name) BETWEEN 1 AND 50),
    parent_id SMALLINT DEFAULT NULL REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE product_category (
    product_id BIGINT NOT NULL REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
    category_id SMALLINT NOT NULL REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    CONSTRAINT product_category_pk PRIMARY KEY (product_id, category_id)
);

CREATE TABLE order_data (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    sum INT DEFAULT 0 NOT NULL CHECK (sum >= 0),
    profile_id BIGINT NOT NULL REFERENCES user_profile(id) ON DELETE CASCADE ON UPDATE CASCADE ,
    order_status TEXT NOT NULL CHECK (order_status IN ('created', 'cancelled', 'ready')),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE order_item (
    order_id BIGINT NOT NULL REFERENCES order_data(id) ON DELETE CASCADE ON UPDATE CASCADE,
    product_id BIGINT NOT NULL REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
    count SMALLINT DEFAULT 1 NOT NULL CHECK (count > 0),
    actual_price INT NOT NULL CHECK (actual_price > 0),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    CONSTRAINT order_item_pk PRIMARY KEY (order_id, product_id)
);

CREATE TABLE cart_item (
    product_id BIGINT NOT NULL REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE,
    profile_id BIGINT NOT NULL REFERENCES user_profile(id) ON DELETE CASCADE ON UPDATE CASCADE,
    count SMALLINT DEFAULT 1 NOT NULL CHECK (count >= 0),
    CONSTRAINT cart_item_pk PRIMARY KEY (profile_id, product_id),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS promocode (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    description TEXT NOT NULL UNIQUE CHECK (char_length(description) > 4),
    min_sum_give INTEGER NOT NULL DEFAULT 0,
    min_sum_activation INTEGER NOT NULL DEFAULT 0,
    benefit_type TEXT NOT NULL CHECK (benefit_type IN ('percentage', 'price discount', 'free delivery')),
    value INTEGER NOT NULL DEFAULT 0,
    ttl_hours SMALLINT NOT NULL CHECK (ttl_hours > 0)
);

CREATE TABLE IF NOT EXISTS promocode_item (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    promocode_type BIGINT NOT NULL REFERENCES promocode(id) ON DELETE CASCADE ON UPDATE CASCADE,
    profile_id BIGINT NOT NULL REFERENCES user_profile(id) ON DELETE CASCADE ON UPDATE CASCADE,
    code TEXT NOT NULL UNIQUE CHECK (char_length(code) BETWEEN 4 AND 8),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE promo_product (
    product_id INT PRIMARY KEY REFERENCES product(id),
    benefit_type TEXT NOT NULL CHECK (benefit_type IN ('percentSale', 'priceSale', 'finalPrice')),
    benefit_value INT NOT NULL CHECK(benefit_value > 0)
);

CREATE TABLE IF NOT EXISTS notification (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    profile_id BIGINT NOT NULL REFERENCES user_profile(id) ON DELETE CASCADE ON UPDATE CASCADE,
    type TEXT NOT NULL CHECK (type IN ('order_status_change')),
    read_status BOOLEAN DEFAULT FALSE,
    payload JSONB NOT NULL CHECK (length(payload::TEXT) > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_product_rating()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE product
    SET rating = (
        SELECT AVG(rating) FROM review WHERE product_id = NEW.product_id
    )
    WHERE id = NEW.product_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_product_rating_trigger
AFTER INSERT ON review
FOR EACH ROW
EXECUTE FUNCTION update_product_rating();
