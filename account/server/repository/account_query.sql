-- name: GetAccount :one
SELECT * FROM account WHERE user_id=$1 LIMIT 1;
