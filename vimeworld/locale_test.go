package vimeworld_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetLocaleByName(t *testing.T) {
	t.Parallel()

	locales, err := client.GetLocaleByName(context.Background(), "ru")
	require.NoError(t, err)
	require.NotEmpty(t, locales)
}
