schema "public" {}

table "account" {
  schema = schema.public
  column "user_id" {
    type = varchar(30)
  }
  column "email" {
    type = text
  }
  column "name" {
    type = text
  }
  column "password" {
    type = varchar(72)
    null = true
  }
  column "google_account_id" {
    type = varchar(255)
    null = true
  }
  column "is_verified" {
    type = boolean
    default = false
  }
  column "created_at" {
    type = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    type = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.user_id]
  }
  index "idx_email" {
    columns = [column.email]
    unique = true
  }
  index "idx_google_account_id" {
    columns = [column.google_account_id]
    unique = true
  }
}
