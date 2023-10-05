package main

import (
	"testing"

	"github.com/go-quicktest/qt"
)

func TestRunSuccess(t *testing.T) {
	err := run()

	qt.Assert(t, qt.IsNil(err))
}
