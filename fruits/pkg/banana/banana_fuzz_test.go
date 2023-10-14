package banana_test

import (
	"testing"

	"github.com/go-quicktest/qt"

	"github.com/marco-m/go-workshop/fruits/pkg/banana"
)

// Go has built-in support for fuzzing. This shows how easy it is.
//
// To get started: https://go.dev/doc/tutorial/fuzz
// For more details: https://go.dev/security/fuzz/
//
// To run as plain test, just use "go test" as usual.
//
// Run as fuzz test until interrupted:
//
//	go test -fuzz=FuzzBananaCut ./pkg/banana
//
// Run as fuzz test with a time limit (example: in CI environment):
//
//	go test -fuzz=FuzzBananaCut -fuzztime=30s ./pkg/banana
func FuzzBananaCut(f *testing.F) {
	for _, seed := range []int{0, 1, 7} {
		f.Add(seed) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, cuts int) {
		// Design choice: we are not interested in negative values for the
		// number of cuts.
		if cuts < 0 {
			return
		}

		sut := banana.Banana{}

		got := sut.Cut(cuts)

		// As any fuzz test, we must find a property (an invariant) on which to
		// assert; we cannot assert on a specific expected output as we do for
		// normal tests.
		qt.Assert(t, qt.Equals(len(got), cuts+1), qt.Commentf("failing input: %v", cuts))
	})
}
