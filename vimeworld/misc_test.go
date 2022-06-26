package vimeworld_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetGames(t *testing.T) {
	t.Parallel()

	games, err := client.GetGames(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, games)
}

func TestClient_GetMaps(t *testing.T) {
	t.Parallel()

	maps, err := client.GetMaps(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, maps)
}

func TestClient_GetAchievements(t *testing.T) {
	t.Parallel()

	achievements, err := client.GetAchievements(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, achievements)
}

func TestClient_GetTokenInfo(t *testing.T) {
	t.Parallel()
	needDevToken(t)

	token, err := client.GetTokenInfo(context.Background(), devToken)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.Equal(t, devToken, token.Token)
}
