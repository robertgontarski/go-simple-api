ALTER TABLE product
DROP INDEX idx_product_name (name),
DROP INDEX idx_product_price (price),
DROP INDEX idx_product_created (created),
DROP INDEX idx_product_updated (updated),
DROP INDEX idx_product_deleted (deleted);

DROP TABLE product;