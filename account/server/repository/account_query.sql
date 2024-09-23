-- name: GetAccount :one
SELECT * FROM account WHERE user_id=$1 LIMIT 1;

-- name: GetAccountByEmail :one
SELECT * FROM account WHERE email=$1 LIMIT 1;

-- name: GetPasswordResetToken :one
SELECT * FROM password_reset WHERE token=$1 LIMIT 1;

-- name: CreateAccount :exec
INSERT INTO account (user_id, email, name, password, google_account_id, is_verified) VALUES ($1, $2, $3, $4, $5, $6);

-- name: CreatePasswordResetToken :exec
INSERT INTO password_reset (token, user_id, expires_at) VALUES ($1, $2, $3);

-- name: UpdateAccountPassword :exec
UPDATE account SET password=$2 WHERE user_id=$1;

-- name: DeletePasswordResetToken :exec
DELETE FROM password_reset WHERE token=$1;
