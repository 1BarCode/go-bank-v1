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

ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "user" ("username");

-- CREATE UNIQUE INDEX ON "account" ("owner", "currency");
ALTER TABLE "account" ADD CONSTRAINT "unique_owner_currency" UNIQUE ("owner", "currency");