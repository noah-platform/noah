-- Create "session" table
CREATE TABLE "session" ("session_id" character(64) NOT NULL, "user_id" character varying(30) NOT NULL, "ip_address" inet NOT NULL, "user_agent" text NOT NULL, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("session_id"));
-- Create index "idx_user_id" to table: "session"
CREATE INDEX "idx_user_id" ON "session" ("user_id");
