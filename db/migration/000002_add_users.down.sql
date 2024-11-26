-- Drop the unique index on accounts
DROP INDEX IF EXISTS accounts_owner_currency_idx;

-- Remove the foreign key constraint from accounts table
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS accounts_owner_fkey;

-- Drop the users table
DROP TABLE IF EXISTS "users";