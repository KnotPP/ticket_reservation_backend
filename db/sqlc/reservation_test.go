package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/KnotPP/ticket_reservation_backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomReservation(t *testing.T) Reservation {
	customer := createRandomUser(t)
	event := createRandomEvent(t)
	arg := CreateReservationParams{
		CustomerID:      customer.ID,
		EventID:         event.ID,
		TicketsReserved: int32(util.RandomInt(1, int(event.TicketQuota))),
		Status:          sql.NullString{String: "reserved", Valid: true},
	}

	reservation, err := testQueries.CreateReservation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reservation)

	require.Equal(t, arg.CustomerID, reservation.CustomerID)
	require.Equal(t, arg.EventID, reservation.EventID)
	require.Equal(t, arg.TicketsReserved, reservation.TicketsReserved)
	require.Equal(t, arg.Status, reservation.Status)

	require.NotZero(t, reservation.ID)
	require.NotZero(t, reservation.CreatedAt)
	return reservation
}

func TestCreateReservation(t *testing.T) {
	createRandomReservation(t)
}

func TestGetReservation(t *testing.T) {
	reservation1 := createRandomReservation(t)
	reservation2, err := testQueries.GetReservation(context.Background(), reservation1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, reservation2)

	require.Equal(t, reservation1.ID, reservation2.ID)
	require.Equal(t, reservation1.CustomerID, reservation2.CustomerID)
	require.Equal(t, reservation1.EventID, reservation2.EventID)
	require.Equal(t, reservation1.TicketsReserved, reservation2.TicketsReserved)
	require.Equal(t, reservation1.Status, reservation2.Status)
	require.WithinDuration(t, reservation1.CreatedAt, reservation2.CreatedAt, time.Second)
}

func TestUpdateReservation(t *testing.T) {
	reservation1 := createRandomReservation(t)
	arg := UpdateReservationParams{
		ID:     reservation1.ID,
		Status: sql.NullString{String: "canceled", Valid: true},
	}
	reservation2, err := testQueries.UpdateReservation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reservation2)

	require.Equal(t, reservation1.ID, reservation2.ID)
	require.Equal(t, reservation1.CustomerID, reservation2.CustomerID)
	require.Equal(t, reservation1.EventID, reservation2.EventID)
	require.Equal(t, reservation1.TicketsReserved, reservation2.TicketsReserved)
	require.Equal(t, arg.Status, reservation2.Status)
	require.WithinDuration(t, reservation1.CreatedAt, reservation2.CreatedAt, time.Second)
}

func TestDeleteReservation(t *testing.T) {
	reservation1 := createRandomReservation(t)
	err := testQueries.DeleteReservation(context.Background(), reservation1.ID)
	require.NoError(t, err)

	reservation2, err2 := testQueries.GetReservation(context.Background(), reservation1.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, reservation2)
}

func TestListReservations(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomReservation(t)
	}

	arg := ListReservationsParams{
		Limit:  5,
		Offset: 5,
	}
	reservations, err := testQueries.ListReservations(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, reservations, 5)

	for _, user := range reservations {
		require.NotEmpty(t, user)
	}
}
