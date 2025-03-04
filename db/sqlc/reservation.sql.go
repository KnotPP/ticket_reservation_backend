// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: reservation.sql

package db

import (
	"context"
	"database/sql"
)

const createReservation = `-- name: CreateReservation :one
INSERT INTO
    reservations (customer_id, event_id, tickets_reserved, status)
VALUES
    ($1, $2, $3, $4) RETURNING id, customer_id, event_id, tickets_reserved, status, created_at
`

type CreateReservationParams struct {
	CustomerID      int32          `json:"customer_id"`
	EventID         int32          `json:"event_id"`
	TicketsReserved int32          `json:"tickets_reserved"`
	Status          sql.NullString `json:"status"`
}

func (q *Queries) CreateReservation(ctx context.Context, arg CreateReservationParams) (Reservation, error) {
	row := q.db.QueryRowContext(ctx, createReservation,
		arg.CustomerID,
		arg.EventID,
		arg.TicketsReserved,
		arg.Status,
	)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.EventID,
		&i.TicketsReserved,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const deleteReservation = `-- name: DeleteReservation :exec
DELETE FROM reservations
WHERE id = $1
`

func (q *Queries) DeleteReservation(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteReservation, id)
	return err
}

const getReservation = `-- name: GetReservation :one
SELECT id, customer_id, event_id, tickets_reserved, status, created_at FROM reservations 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetReservation(ctx context.Context, id int32) (Reservation, error) {
	row := q.db.QueryRowContext(ctx, getReservation, id)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.EventID,
		&i.TicketsReserved,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listReservations = `-- name: ListReservations :many
SELECT id, customer_id, event_id, tickets_reserved, status, created_at FROM reservations
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListReservationsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListReservations(ctx context.Context, arg ListReservationsParams) ([]Reservation, error) {
	rows, err := q.db.QueryContext(ctx, listReservations, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Reservation{}
	for rows.Next() {
		var i Reservation
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.EventID,
			&i.TicketsReserved,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateReservation = `-- name: UpdateReservation :one
UPDATE reservations
SET status = $2
WHERE id = $1
RETURNING id, customer_id, event_id, tickets_reserved, status, created_at
`

type UpdateReservationParams struct {
	ID     int32          `json:"id"`
	Status sql.NullString `json:"status"`
}

func (q *Queries) UpdateReservation(ctx context.Context, arg UpdateReservationParams) (Reservation, error) {
	row := q.db.QueryRowContext(ctx, updateReservation, arg.ID, arg.Status)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.EventID,
		&i.TicketsReserved,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
