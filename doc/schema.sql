-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-03-10T11:07:55.452Z

CREATE TABLE "user" (
  "username" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

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

CREATE UNIQUE INDEX ON "account" ("owner", "currency");

CREATE INDEX ON "entry" ("account_id");

CREATE INDEX ON "transfer" ("from_account_id");

CREATE INDEX ON "transfer" ("to_account_id");

CREATE INDEX ON "transfer" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entry"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfer"."amount" IS 'MUST be positive';

ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "user" ("username");

ALTER TABLE "entry" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");
