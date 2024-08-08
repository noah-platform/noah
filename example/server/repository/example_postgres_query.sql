-- name: GetExample :one
SELECT * FROM example WHERE example_id=$1 LIMIT 1;
