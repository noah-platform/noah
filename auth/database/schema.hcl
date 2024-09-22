schema "public" {}

table "session" {
  schema = schema.public
  column "session_id" {
    type = char(64)
  }
  column "user_id" {
    type = varchar(30)
  }
  column "ip_address" {
    type = inet
  }
  column "user_agent" {
    type = text
  }
  column "created_at" {
    type = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  index "idx_user_id" {
    columns = [column.user_id]
  }
  primary_key {
    columns = [column.session_id]
  }
}
