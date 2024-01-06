-- name: CreateItem :one
INSERT INTO items (id, created_at, updated_at, name, description, price, sold, seller_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
