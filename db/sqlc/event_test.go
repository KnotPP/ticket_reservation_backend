package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/KnotPP/ticket_reservation_backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomEvent(t *testing.T) Event {
	user := createRandomUser(t)
	arg := CreateEventParams{
		OrganizerID: user.ID,
		Name:        util.RandomName(),
		TicketQuota: util.RandomTicketQuota(),
		Price:       util.RandomPrice(),
	}

	event, err := testQueries.CreateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, arg.Name, event.Name)
	require.Equal(t, arg.OrganizerID, event.OrganizerID)
	require.Equal(t, arg.TicketQuota, event.TicketQuota)
	require.Equal(t, arg.Price, event.Price)

	require.NotZero(t, event.ID)
	require.NotZero(t, event.CreatedAt)
	return event
}

func TestCreateEvent(t *testing.T) {
	createRandomEvent(t)
}

func TestGetEvent(t *testing.T) {
	event1 := createRandomEvent(t)
	event2, err := testQueries.GetEvent(context.Background(), event1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, event2)

	require.Equal(t, event1.ID, event2.ID)
	require.Equal(t, event1.OrganizerID, event2.OrganizerID)
	require.Equal(t, event1.Name, event2.Name)
	require.Equal(t, event1.TicketQuota, event2.TicketQuota)
	require.Equal(t, event1.Price, event2.Price)
	require.WithinDuration(t, event1.CreatedAt, event2.CreatedAt, time.Second)
}

func TestUpdateEvent(t *testing.T) {
	event1 := createRandomEvent(t)
	arg := UpdateEventParams{
		ID:          event1.ID,
		TicketQuota: util.RandomTicketQuota(),
	}
	event2, err := testQueries.UpdateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event2)

	require.Equal(t, event1.ID, event2.ID)
	require.Equal(t, event1.OrganizerID, event2.OrganizerID)
	require.Equal(t, event1.Name, event2.Name)
	require.Equal(t, arg.TicketQuota, event2.TicketQuota)
	require.Equal(t, event1.Price, event2.Price)
	require.WithinDuration(t, event1.CreatedAt, event2.CreatedAt, time.Second)
}

func TestDeleteEvent(t *testing.T) {
	event1 := createRandomEvent(t)
	err := testQueries.DeleteEvent(context.Background(), event1.ID)
	require.NoError(t, err)

	event2, err2 := testQueries.GetEvent(context.Background(), event1.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, event2)
}

func TestListEvents(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEvent(t)
	}

	arg := ListEventsParams{
		Limit:  5,
		Offset: 5,
	}
	events, err := testQueries.ListEvents(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, events, 5)

	for _, user := range events {
		require.NotEmpty(t, user)
	}
}
