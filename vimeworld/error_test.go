package vimeworld_test

import (
	"github.com/EpicStep/go-vimeworld/v1/vimeworld"
	"github.com/stretchr/testify/require"
	"testing"
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
