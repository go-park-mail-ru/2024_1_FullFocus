-- +goose Up
-- +goose StatementBegin

SET search_path TO ozon, public;

BEGIN;
    ALTER TABLE order_item
    ADD COLUMN actual_price INT;
    
    UPDATE order_item 
    SET actual_price = (
	    SELECT product.price
	    FROM product
	    JOIN order_item o ON product.id = o.product_id
	    WHERE product.id = order_item.product_id 
    )
    WHERE product_id = (
	    SELECT product.id
	    FROM product
	    JOIN order_item o ON product.id = o.product_id
	    WHERE product.id = order_item.product_id 
    );

    ALTER TABLE order_item
    ALTER COLUMN actual_price SET NOT NULL;
COMMIT;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

SET search_path TO ozon, public;

ALTER TABLE order_item
DROP COLUMN actual_price;

-- +goose StatementEnd
