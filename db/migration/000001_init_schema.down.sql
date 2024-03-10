DROP TABLE IF EXISTS "entry";
DROP TABLE IF EXISTS "transfer";
DROP TABLE IF EXISTS "account";
DROP EXTENSION IF EXISTS "uuid-ossp";
-- drop entry and transfer tables first to avoid foreign key constraint violation