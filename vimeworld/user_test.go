package vimeworld_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetUsersByNames(t *testing.T) {
	t.Parallel()

	users, err := client.GetUsersByNames(context.Background(), "xtrafrancyz", "EpicStep")
	require.NoError(t, err)
	require.Equal(t, 134568, users[0].ID)
	require.Equal(t, 1402876, users[1].ID)
}

func TestClient_GetUsersByIds(t *testing.T) {
	t.Parallel()

	users, err := client.GetUsersByIds(context.Background(), 134568, 1402876)
	require.NoError(t, err)
	require.Equal(t, "xtrafrancyz", users[0].Username)
	require.Equal(t, "EpicStep", users[1].Username)
}

func TestClient_GetUserFriends(t *testing.T) {
	t.Parallel()

	f, err := client.GetUserFriends(context.Background(), 1402876)
	require.NoError(t, err)
	require.NotEmpty(t, f)
}

func TestClient_GetUserSession(t *testing.T) {
	t.Parallel()

	s, err := client.GetUserSession(context.Background(), 1402876)
	require.NoError(t, err)
	require.Equal(t, "EpicStep", s.User.Username)
}

func TestClient_GetUserStats(t *testing.T) {
	t.Parallel()

	s, err := client.GetUserStats(context.Background(), 1402876)
	require.NoError(t, err)
	require.Equal(t, "EpicStep", s.User.Username)
}

func TestClient_GetUserAchievements(t *testing.T) {
	t.Parallel()

	s, err := client.GetUserAchievements(context.Background(), 1402876)
	require.NoError(t, err)
	require.Equal(t, "EpicStep", s.User.Username)
}

func TestClient_GetUserLeaderboards(t *testing.T) {
	t.Parallel()

	s, err := client.GetUserLeaderboards(context.Background(), 1402876)
	require.NoError(t, err)
	require.Equal(t, "EpicStep", s.User.Username)
}

func TestClient_GetUserMatchesBefore(t *testing.T) {
	t.Parallel()

	m, err := client.GetUserMatchesBefore(context.Background(), 1402876, "461220632323948544", 1)
	require.NoError(t, err)
	require.Equal(t, m.User.Username, "EpicStep")
}

func TestClient_GetUserMatchesAfter(t *testing.T) {
	t.Parallel()

	m, err := client.GetUserMatchesAfter(context.Background(), 1402876, "461220632323948544", 1)
	require.NoError(t, err)
	require.Equal(t, m.User.Username, "EpicStep")
}

func TestClient_GetUserMatchesOffset(t *testing.T) {
	t.Parallel()

	m, err := client.GetUserMatchesOffset(context.Background(), 1402876, 0, 10)
	require.NoError(t, err)
	require.Equal(t, "EpicStep", m.User.Username)
}

func TestClient_GetUsersSessions(t *testing.T) {
	t.Parallel()

	t.Run("GetFewUsers", func(t *testing.T) {
		t.Parallel()

		s, err := client.GetUsersSessions(context.Background(), 1402876, 134568)
		require.NoError(t, err)
		require.Len(t, s, 2)
	})

	t.Run("GetManyUsers", func(t *testing.T) {
		t.Parallel()

		need := 51

		users := make([]int, 0, need)
		for i := 1; i < need+1; i++ {
			users = append(users, i)
		}

		s, err := client.GetUsersSessions(context.Background(), users...)
		require.NoError(t, err)
		require.Len(t, s, need)
	})
}
