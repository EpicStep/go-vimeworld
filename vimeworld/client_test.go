package vimeworld_test

import (
	"os"
	"testing"

	"github.com/EpicStep/go-vimeworld/v1/vimeworld"
)

var client *vimeworld.Client
var devToken string

func TestMain(m *testing.M) {
	var err error

	client, err = vimeworld.NewClient(vimeworld.Options{})
	if err != nil {
		os.Exit(1)
	}

	if token := os.Getenv("VIMEWORLD_TOKEN"); token != "" {
		devToken = token
	}

	runTests := m.Run()
	os.Exit(runTests)
}

func needDevToken(t *testing.T) {
	if devToken == "" {
		t.Skip("need dev token")
	}
}
