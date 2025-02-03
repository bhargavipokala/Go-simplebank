-- name: CreateEntry :one
INSERT INTO entries (account_id, amount) VALUES ($1, $2) RETURNING *;

-- name: GetEntryByID :one
SELECT * FROM entries WHERE id = $1;

-- name: GetEntryByAccountID :many
SELECT * FROM entries WHERE account_id = $1;
