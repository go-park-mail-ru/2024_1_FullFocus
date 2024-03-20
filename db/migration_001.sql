-- ozon.user definition

CREATE TABLE IF NOT EXISTS user (
	id uuid NOT NULL,
	user_login text NOT NULL,
	password_hash text NOT NULL,
	created_at timetz DEFAULT now() NOT NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT user_check CHECK (char_length(user_login) BETWEEN 4 AND 32),
	CONSTRAINT user_pk PRIMARY KEY (id),
	CONSTRAINT user_unique UNIQUE (login)
);

-- ozon.profile definition

CREATE TABLE IF NOT EXISTS profile (
	id uuid NOT NULL,
	full_name text NOT NULL,
	email text NOT NULL,
	imgsrc text DEFAULT 'default-avatar.png' NOT NULL,
	phone_number text NOT NULL,
	user_id uuid NOT NULL,
	created_at timetz DEFAULT now() NOT NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT email_lenght CHECK (char_length(email) BETWEEN 4 AND 255),
	CONSTRAINT email_valid CHECK (email ~* '^[a-z0-9\.\-]+@[a-z0-9\.\-]+\.[a-z]+$'),
	CONSTRAINT name_length CHECK (char_length(full_name) BETWEEN 5 AND 255),
	CONSTRAINT phone_length CHECK (char_length(phone_number) BETWEEN 5 AND 15),
	CONSTRAINT phone_valid CHECK (phone_number ~ '\+?[0-9]+'),
	CONSTRAINT profile_pk PRIMARY KEY (id),
	CONSTRAINT profile_user_fk FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);

-- ozon.product definition

CREATE TABLE IF NOT EXISTS product (
	id uuid NOT NULL,
	product_name text NOT NULL,
	product_description text NULL,
	price numeric NOT NULL,
	imgsrc text DEFAULT 'default-product.png' NOT NULL,
	seller text NOT NULL,
	rating int4 DEFAULT 0 NOT NULL,
	category_id int4 NOT NULL,
	created_at timetz DEFAULT now() NOT NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT descriprion_length CHECK (char_length(procust_description) BETWEEN 1 AND 255),
	CONSTRAINT name_length CHECK (char_length(product_name) BETWEEN 1 AND 50),
	CONSTRAINT price_positive CHECK (price > 0),
	CONSTRAINT product_pk PRIMARY KEY (id)
);

-- ozon.category definition

CREATE TABLE IF NOT EXISTS category (
	id smallserial NOT NULL,
	category_name text NOT NULL,
	parent_id int2 NOT NULL,
	CONSTRAINT category_pk PRIMARY KEY (id),
	CONSTRAINT category_unique UNIQUE (category_name),
	CONSTRAINT name_length CHECK (char_length(category_name) BETWEEN 1 AND 50),
	CONSTRAINT category_category_fk FOREIGN KEY (parent_id) REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- ozon.product_category definition

CREATE TABLE IF NOT EXISTS product_category (
	product_id uuid NOT NULL,
	category_id int2 NOT NULL,
	product_category_pk int8 NOT NULL,
	created_at timetz DEFAULT now() NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT product_category_pk PRIMARY KEY (product_id, category_id),
	CONSTRAINT product_category_category_fk FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT product_category_product_fk FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- ozon.order_item definition

CREATE TABLE IF NOT EXISTS order_item (
	id uuid NOT NULL,
	product_id uuid NOT NULL,
	count int2 DEFAULT 1 NOT NULL,
	created_at timetz DEFAULT now() NOT NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT count_positive CHECK (count >= 0),
	CONSTRAINT order_item_pk PRIMARY KEY (id),
	CONSTRAINT order_item_product_fk FOREIGN KEY (product_id) REFERENCES product(id)
);

-- ozon.ordering definition

CREATE TABLE IF NOT EXISTS ordering (
	id uuid NOT NULL,
	sum int4 DEFAULT 0 NOT NULL,
	order_item uuid NOT NULL,
	profile_id uuid NOT NULL,
	created_at timetz DEFAULT now() NOT NULL,
	updated_at timetz DEFAULT now() NOT NULL,
	CONSTRAINT ordering_pk PRIMARY KEY (id),
	CONSTRAINT sum_positive CHECK (sum >= 0),
	CONSTRAINT order_item_fk FOREIGN KEY (order_item) REFERENCES order_item(id),
	CONSTRAINT ordering_profile_fk FOREIGN KEY (profile_id) REFERENCES profile(id)
);