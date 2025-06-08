package vimeworld_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GuildSearch(t *testing.T) {
	t.Parallel()

	guilds, err := client.GuildSearch(context.Background(), "f5")
	require.NoError(t, err)
	require.NotEmpty(t, guilds)
}

func TestClient_GetGuildByID(t *testing.T) {
	t.Parallel()

	guild, err := client.GetGuildByID(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, "Fantastic Five", guild.Name)
}

func TestClient_GetGuildByName(t *testing.T) {
	t.Parallel()

	guild, err := client.GetGuildByName(context.Background(), "VimeTop")
	require.NoError(t, err)
	require.Equal(t, 104, guild.ID)
}

func TestClient_GetGuildByTag(t *testing.T) {
	t.Parallel()

	guild, err := client.GetGuildByTag(context.Background(), "-F5-")
	require.NoError(t, err)
	require.Equal(t, "Fantastic Five", guild.Name)
}
