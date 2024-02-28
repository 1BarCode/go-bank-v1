-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-02-28T10:50:12.048Z

CREATE TABLE "account" (
  "id" uuid PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()) ON UPDATE now()
);

CREATE TABLE "entry" (
  "id" uuid PRIMARY KEY,
  "account_id" uuid,
  "amount" bigint NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()) ON UPDATE now()
);

CREATE TABLE "transfer" (
  "id" uuid PRIMARY KEY,
  "from_account_id" uuid,
  "to_account_id" uuid,
  "amount" bigint NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()) ON UPDATE now()
);

CREATE INDEX ON "account" ("owner");

CREATE INDEX ON "entry" ("account_id");

CREATE INDEX ON "transfer" ("from_account_id");

CREATE INDEX ON "transfer" ("to_account_id");

CREATE INDEX ON "transfer" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "account"."updated_at" IS 'manually add timestamp on update';

COMMENT ON COLUMN "entry"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "entry"."updated_at" IS 'manually add timestamp on update';

COMMENT ON COLUMN "transfer"."amount" IS 'MUST be positive';

COMMENT ON COLUMN "transfer"."updated_at" IS 'manually add timestamp on update';

ALTER TABLE "entry" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");
