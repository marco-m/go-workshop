package main

import (
	"testing"
)

func TestRunSuccess(t *testing.T) {
	err := run()

	if err != nil {
		t.Fatalf("got: %s; want: <no error>", err)
	}
}
