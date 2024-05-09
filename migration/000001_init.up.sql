CREATE TABLE product (
    id INT UNSIGNED AUTO_INCREMENT,
    name VARCHAR(255),
    price DECIMAL(10, 2),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted TIMESTAMP,
    PRIMARY KEY (id)
);

ALTER TABLE product
ADD INDEX idx_product_name (name),
ADD INDEX idx_product_price (price),
ADD INDEX idx_product_created (created),
ADD INDEX idx_product_updated (updated),
ADD INDEX idx_product_deleted (deleted);