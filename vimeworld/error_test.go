package vimeworld_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/EpicStep/go-vimeworld/vimeworld"
)

func TestError_Error(t *testing.T) {
	t.Parallel()

	e := vimeworld.Error{
		ErrorCode: 1,
		ErrorMsg:  "Test",
		Comment:   "Test comment",
	}

	require.Equal(t, "go-vimeworld: Test", e.Error())
}
