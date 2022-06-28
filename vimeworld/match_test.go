package vimeworld_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetMatchLatest(t *testing.T) {
	t.Parallel()

	m, err := client.GetMatchLatest(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, m)
}

func TestClient_GetMatchListBefore(t *testing.T) {
	t.Parallel()

	before, err := client.GetMatchListBefore(context.Background(), "461280233589309440", 1)
	require.NoError(t, err)
	require.Equal(t, "BP", before[0].Game)
}

func TestClient_GetMatchListAfter(t *testing.T) {
	t.Parallel()

	before, err := client.GetMatchListAfter(context.Background(), "461280233589309440", 1)
	require.NoError(t, err)
	require.Equal(t, "BWH", before[0].Game)
}
