ALTER TABLE client
DROP INDEX idx_client_created (created),
DROP INDEX idx_client_updated (updated),
DROP INDEX idx_client_deleted (deleted);

DROP TABLE client;