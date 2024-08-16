-- Create "password_reset" table
CREATE TABLE "password_reset" ("token" character(64) NOT NULL, "user_id" character varying(30) NOT NULL, "expires_at" timestamp NOT NULL, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("token"), CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "account" ("user_id") ON UPDATE CASCADE ON DELETE CASCADE);
