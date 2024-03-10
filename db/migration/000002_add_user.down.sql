ALTER TABLE IF EXISTS "account" DROP CONSTRAINT "unique_owner_currency";

ALTER TABLE IF EXISTS "account" DROP CONSTRAINT "account_owner_fkey";

DROP TABLE IF EXISTS "user";