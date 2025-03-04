-- name: CreateUser :one
INSERT INTO
    users (name, email, password_hash, role)
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetUser :one
SELECT id, name, email, role, created_at FROM users 
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET name = $2, role = $3
WHERE id = $1
RETURNING id, name, email, role, created_at;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;