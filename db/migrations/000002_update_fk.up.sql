ALTER TABLE "entries" DROP CONSTRAINT entries_account_id_fkey;
ALTER TABLE "transfers" DROP CONSTRAINT transfers_from_account_id_fkey;
ALTER TABLE "transfers" DROP CONSTRAINT transfers_to_account_id_fkey;

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;
