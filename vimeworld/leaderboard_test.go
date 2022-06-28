package vimeworld_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetLeaderboardList(t *testing.T) {
	t.Parallel()

	l, err := client.GetLeaderboardList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, l)
}

func TestClient_GetLeaderboard(t *testing.T) {
	t.Parallel()

	l, err := client.GetLeaderboard(context.Background(), "user", "level", 0, 0)
	require.NoError(t, err)
	require.NotEmpty(t, l)
}
