-- name: CreateUser :one
INSERT INTO
    users (name, email, password_hash, role)
VALUES
    ($1, $2, $3, $4) RETURNING *;