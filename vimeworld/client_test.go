package vimeworld_test

import (
	"github.com/EpicStep/go-vimeworld/v1/vimeworld"
	"os"
	"testing"
)

var client *vimeworld.Client

func TestMain(m *testing.M) {
	var err error

	client, err = vimeworld.NewClient(vimeworld.Options{})
	if err != nil {
		os.Exit(1)
	}

	runTests := m.Run()
	os.Exit(runTests)
}
