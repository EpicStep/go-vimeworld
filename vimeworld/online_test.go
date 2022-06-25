package vimeworld_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetOnline(t *testing.T) {
	t.Parallel()

	online, err := client.GetOnline(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, online.Total)
	require.NotEmpty(t, online.Separated)
}

func TestClient_GetOnlineStreams(t *testing.T) {
	t.Parallel()

	_, err := client.GetOnlineStreams(context.Background())
	require.NoError(t, err)
}

func TestClient_GetOnlineStaff(t *testing.T) {
	t.Parallel()

	_, err := client.GetOnlineStaff(context.Background())
	require.NoError(t, err)
}