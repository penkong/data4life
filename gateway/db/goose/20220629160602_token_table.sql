-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SEQUENCE seq START 1 INCREMENT 1 MAXVALUE 1000000;
CREATE TABLE token (
  id   BIGSERIAL PRIMARY KEY,
  name  VARCHAR(8) UNIQUE NOT NULL,
  occur integer NOT NULL DEFAULT 1
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE token CASCADE;
DROP SEQUENCE IF EXISTS seq CASCADE;
-- +goose StatementEnd
