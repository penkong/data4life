CREATE TABLE token (
  id   BIGSERIAL PRIMARY KEY,
  name  VARCHAR(8) UNIQUE NOT NULL,
  occur integer NOT NULL DEFAULT 1
);
