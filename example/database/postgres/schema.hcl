schema "public" {}

table "example" {
  schema = schema.public
  column "example_id" {
    type = serial
  }
  column "title" {
    type = text
  }
  primary_key {
    columns = [column.example_id]
  }
}
