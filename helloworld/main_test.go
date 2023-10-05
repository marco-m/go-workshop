package main

import (
	"testing"

	"github.com/go-quicktest/qt"
)

// Note that we are still not testing the output to stdout.
func TestRunSuccess(t *testing.T) {
	err := run([]string{})

	qt.Assert(t, qt.IsNil(err))
}

func TestRunFailure(t *testing.T) {
	err := run([]string{"monkey"})

	qt.Assert(t, qt.ErrorMatches(err, `unexpected: \["monkey"\]`))
}
