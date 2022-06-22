-- CREATE USER root;
CREATE DATABASE datalifepg;
--          -------------------------
-- GRANT ALL PRIVILEGES ON DATABASE devref-ts-express-rest-pg-auth TO root;

-- CREATE TABLE IF NOT EXISTS users (
--   "user_id" BIGSERIAL PRIMARY KEY,
--   "email" varchar(75) UNIQUE NOT NULL,
--   "hashed_pass" varchar NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now()),
--   "updated_at" timestamptz,
--   "deleted_at" timestamptz,
--   "deleted" boolean DEFAULT false,
--   "rv" integer NOT NULL DEFAULT 0,
--   CHECK(created_at < updated_at)
-- );
