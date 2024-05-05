-- +goose Up
-- +goose StatementBegin

SET search_path TO ozon, public;

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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

SET search_path TO ozon, public;

DROP TRIGGER update_product_rating_trigger ON review;
DROP FUNCTION update_product_rating;

-- +goose StatementEnd
