CREATE TABLE client (
    id INT UNSIGNED AUTO_INCREMENT,
    email VARCHAR(255),
    password VARCHAR(255),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted TIMESTAMP,
    PRIMARY KEY (id)
);

ALTER TABLE client
ADD INDEX idx_client_created (created),
ADD INDEX idx_client_updated (updated),
ADD INDEX idx_client_deleted (deleted);