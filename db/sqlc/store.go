package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

// SQLStore provides all functions to execute db queries and transactions.
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store.
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type ReservationTxParams struct {
	CustomerID      int32 `json:"customer_id"`
	EventID         int32 `json:"event_id"`
	TicketsReserved int32 `json:"tickets_reserved"`
}

type ReservationTxResult struct {
	Reservation Reservation `json:"reservation"`
	Event       Event       `json:"event"`
}

func (store *Store) ReservationTx(ctx context.Context, arg ReservationTxParams) (ReservationTxResult, error) {
	var result ReservationTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Reservation, err = q.CreateReservation(ctx, CreateReservationParams{
			CustomerID:      arg.CustomerID,
			EventID:         arg.EventID,
			TicketsReserved: arg.TicketsReserved,
			Status:          sql.NullString{String: "reserved", Valid: true},
		})
		if err != nil {
			return err
		}

		result.Event, err = q.DeductTicketQuota(ctx, DeductTicketQuotaParams{
			EventID: arg.EventID,
			Amount:  arg.TicketsReserved,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
