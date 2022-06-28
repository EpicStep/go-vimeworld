package vimeworld_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetLocaleByName(t *testing.T) {
	t.Parallel()

	locales, err := client.GetLocaleByName(context.Background(), "ru")
	require.NoError(t, err)
	require.NotEmpty(t, locales)
}
