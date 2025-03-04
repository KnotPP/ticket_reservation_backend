-- name: CreateReservation :one
INSERT INTO
    reservations (customer_id, event_id, tickets_reserved, status)
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetReservation :one
SELECT * FROM reservations 
WHERE id = $1 LIMIT 1;

-- name: ListReservations :many
SELECT * FROM reservations
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateReservation :one
UPDATE reservations
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteReservation :exec
DELETE FROM reservations
WHERE id = $1;