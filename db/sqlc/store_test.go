package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReservationTx(t *testing.T) {
	store := NewStore(testDB)

	user := createRandomUser(t)
	event := createRandomEvent(t)

	n := 5
	amount := int32(10)

	errs := make(chan error)
	results := make(chan ReservationTxResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.ReservationTx(ctx, ReservationTxParams{
				CustomerID:      user.ID,
				EventID:         event.ID,
				TicketsReserved: int32(amount),
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		reservation := result.Reservation
		require.NotEmpty(t, reservation)
		require.Equal(t, reservation.CustomerID, user.ID)
		require.Equal(t, reservation.EventID, event.ID)
		require.Equal(t, reservation.TicketsReserved, amount)

		event1 := result.Event
		require.NotEmpty(t, event1)
		fmt.Println("Before:", event1.TicketQuota, "After:", event.TicketQuota-amount*int32(i+1))
		require.Equal(t, event1.TicketQuota, event.TicketQuota-amount*int32(i+1))
	}

	eventAfter, err := store.GetEvent(context.Background(), event.ID)
	require.NoError(t, err)
	require.Equal(t, eventAfter.TicketQuota, event.TicketQuota-amount*int32(n))
}
