-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-02-28T23:42:35.530Z

-- Install the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "account" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entry" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "account_id" uuid NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "from_account_id" uuid NOT NULL,
  "to_account_id" uuid NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "account" ("owner");

CREATE INDEX ON "entry" ("account_id");

CREATE INDEX ON "transfer" ("from_account_id");

CREATE INDEX ON "transfer" ("to_account_id");

CREATE INDEX ON "transfer" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entry"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfer"."amount" IS 'MUST be positive';

ALTER TABLE "entry" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");
