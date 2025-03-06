-- name: CreateEvent :one
INSERT INTO
    events (organizer_id, name, ticket_quota, price)
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetEvent :one
SELECT * FROM events 
WHERE id = $1 LIMIT 1;

-- name: ListEvents :many
SELECT * FROM events
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEvent :one
UPDATE events
SET ticket_quota = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events 
WHERE id = $1;

-- name: DeductTicketQuota :one
UPDATE events
SET ticket_quota = ticket_quota - sqlc.arg(amount)
WHERE id = sqlc.arg(event_id)
RETURNING *;
