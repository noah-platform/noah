-- Create index "idx_email" to table: "account"
CREATE UNIQUE INDEX "idx_email" ON "account" ("email");
-- Create index "idx_google_account_id" to table: "account"
CREATE UNIQUE INDEX "idx_google_account_id" ON "account" ("google_account_id");
