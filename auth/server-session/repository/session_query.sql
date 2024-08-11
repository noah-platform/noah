-- name: GetSession :one
SELECT * FROM session WHERE session_id=$1;

-- name: CreateSession :one
INSERT INTO session (session_id, user_id, ip_address, user_agent) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: DeleteSession :one
DELETE FROM session WHERE session_id=$1 RETURNING *;
