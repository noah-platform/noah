-- Create index "idx_verification_token" to table: "account"
CREATE UNIQUE INDEX "idx_verification_token" ON "account" ("verification_token");
